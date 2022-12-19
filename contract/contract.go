package contract

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fullstackwang/tron-grpc/abi"
	"github.com/fullstackwang/tron-grpc/address"
	"github.com/fullstackwang/tron-grpc/client"
	"github.com/fullstackwang/tron-grpc/core"
	"github.com/fullstackwang/tron-grpc/tx"
)

const defaultFeeLimit = 10000000

type SendOption struct {
	FeeLimit int64
}

type ConstantMethod func(ctx context.Context, args ...any) ([]any, error)
type Method func(ctx context.Context, args ...any) (*tx.Transaction, error)

type EventInput struct {
	Name  string
	Value any
}

type Event struct {
	Name        string
	Sig         []byte
	IsAnonymous bool
	Address     address.Address
	Inputs      []EventInput
}

type Contract struct {
	address address.Address
	client  *client.Client

	constantMethods map[string]ConstantMethod
	methods         map[string]Method

	eventSigMap map[string]*abi.Event
	events      map[string]*abi.Event
}

func New(client *client.Client, addr address.Address) *Contract {
	return &Contract{
		address:         addr,
		client:          client,
		constantMethods: make(map[string]ConstantMethod),
		methods:         make(map[string]Method),
		eventSigMap:     make(map[string]*abi.Event),
		events:          make(map[string]*abi.Event),
	}
}

func (c *Contract) LoadABI(abiJson []byte) error {
	iface, err := abi.Parse(abiJson)
	if err != nil {
		return err
	}
	for _, m := range iface.Methods {
		if m.IsConstant {
			c.constantMethods[m.Name] = c.createConstantMethod(&m)
		} else {
			c.methods[m.Name] = c.createMethod(&m)
		}
	}
	for _, event := range iface.Events {
		if !event.IsAnonymous {
			c.eventSigMap[string(event.Sig)] = &event
		}
		c.events[event.Name] = &event
	}
	return nil
}

func (c *Contract) GetConstantMethod(methodName string) ConstantMethod {
	return c.constantMethods[methodName]
}

func (c *Contract) GetMethod(methodName string) Method {
	return c.methods[methodName]
}

func (c *Contract) getTriggerSmartContract(m *abi.Method, args []any) (*core.TriggerSmartContract, error) {
	inputData, err := m.InputEncoder.Encode(args)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	buf.Write(m.Sig)
	buf.Write(inputData)
	return &core.TriggerSmartContract{
		OwnerAddress:    c.client.Signer.Address(),
		ContractAddress: c.address,
		Data:            buf.Bytes(),
	}, nil
}

func (c *Contract) createConstantMethod(m *abi.Method) ConstantMethod {
	return func(ctx context.Context, args ...any) ([]any, error) {
		in, err := c.getTriggerSmartContract(m, args)
		if err != nil {
			return nil, err
		}
		t, err := c.client.TriggerConstantContract(ctx, in)
		if err != nil {
			return nil, err
		}
		if t.Result.Code > 0 {
			return nil, fmt.Errorf(string(t.Result.Message))
		}
		return m.OutputDecoder.Decode(t.ConstantResult)
	}
}

func getSendOption(args []any) *SendOption {
	if len(args) == 0 {
		return nil
	}
	if option, ok := args[len(args)-1].(*SendOption); ok {
		return option
	}
	return nil
}

func (c *Contract) createMethod(m *abi.Method) Method {
	return func(ctx context.Context, args ...any) (*tx.Transaction, error) {
		option := getSendOption(args)

		feeLimit := int64(defaultFeeLimit)
		if option != nil {
			feeLimit = option.FeeLimit
		}

		in, err := c.getTriggerSmartContract(m, args)
		if err != nil {
			return nil, err
		}

		t, err := c.client.TriggerContract(ctx, in)
		if err != nil {
			return nil, err
		}
		if t.Result.Code > 0 {
			return nil, fmt.Errorf(string(t.Result.Message))
		}

		t.Transaction.RawData.FeeLimit = feeLimit
		tt := tx.NewWithDecoder(c.client, t.Transaction, m.OutputDecoder.Decode)
		return tt, tt.Send(ctx, c.client.Signer)
	}
}

func (c *Contract) Call(ctx context.Context, methodName string, args ...any) ([]any, error) {
	m := c.constantMethods[methodName]
	if m == nil {
		return nil, fmt.Errorf("method not found")
	}
	return m(ctx, args...)
}

func (c *Contract) Send(ctx context.Context, methodName string, args ...any) (*tx.Transaction, error) {
	m := c.methods[methodName]
	if m == nil {
		return nil, fmt.Errorf("method not found")
	}
	return m(ctx, args...)
}

func decodeEvent(ed *abi.Event, log *core.TransactionInfo_Log) (Event, error) {
	dataValues, err := ed.Decoder.DecodeData(log.Data)
	if err != nil {
		return Event{}, err
	}

	addr, err := ed.Decoder.DecodeAddr(log.Address)
	if err != nil {
		return Event{}, err
	}

	topicOffset := 0
	dataOffset := 0
	var inputs []EventInput

	for _, input := range ed.Inputs {
		if input.Indexed {
			v, err := ed.Decoder.DecodeTopic(topicOffset, log.Topics[topicOffset+1])
			if err != nil {
				return Event{}, err
			}
			inputs = append(inputs, EventInput{
				Name:  input.Name,
				Value: v,
			})
			topicOffset = topicOffset + 1
		} else {
			inputs = append(inputs, EventInput{
				Name:  input.Name,
				Value: dataValues[dataOffset],
			})
			dataOffset = dataOffset + 1
		}
	}
	return Event{
		Name:        ed.Name,
		Sig:         ed.Sig,
		IsAnonymous: ed.IsAnonymous,
		Address:     addr.(address.Address),
		Inputs:      inputs,
	}, nil
}

func (c *Contract) GetEvents(tx *tx.Transaction) ([]Event, error) {
	var events []Event
	for _, log_ := range tx.Info.Log {
		if !bytes.Equal(log_.Address, c.address.ToEthAddress()) {
			continue
		}
		ed := c.eventSigMap[string(log_.Topics[0])]
		if ed == nil {
			continue
		}
		e, err := decodeEvent(ed, log_)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (c *Contract) GetEventsByName(tx *tx.Transaction, eventName string) ([]Event, error) {
	ed := c.events[eventName]
	if ed == nil {
		return nil, fmt.Errorf("event type not found")
	}

	var events []Event
	for _, log_ := range tx.Info.Log {
		if !bytes.Equal(log_.Address, c.address.ToEthAddress()) {
			continue
		}
		if !bytes.Equal(log_.Topics[0], ed.Sig) {
			continue
		}
		e, err := decodeEvent(ed, log_)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
