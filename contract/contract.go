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

var (
	ErrMethodNotFound    = fmt.Errorf("method not found")
	ErrEventTypeNotFound = fmt.Errorf("event type not found")
)

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
	Signer  client.Signer

	abiMethods      map[string]*abi.Method
	constantMethods map[string]ConstantMethod
	methods         map[string]Method

	eventSigMap map[string]*abi.Event
	events      map[string]*abi.Event
}

func New(client *client.Client, addr address.Address) *Contract {
	return &Contract{
		address:         addr,
		client:          client,
		abiMethods:      make(map[string]*abi.Method),
		constantMethods: make(map[string]ConstantMethod),
		methods:         make(map[string]Method),
		eventSigMap:     make(map[string]*abi.Event),
		events:          make(map[string]*abi.Event),
	}
}

func (c *Contract) Clone() *Contract {
	newContract := &Contract{
		address:         c.address,
		client:          c.client,
		Signer:          c.Signer,
		abiMethods:      c.abiMethods,
		constantMethods: make(map[string]ConstantMethod),
		methods:         make(map[string]Method),
		eventSigMap:     c.eventSigMap,
		events:          c.events,
	}

	for _, m := range c.abiMethods {
		if m.IsConstant {
			newContract.constantMethods[m.Name] = newContract.createConstantMethod(m)
		} else {
			newContract.methods[m.Name] = newContract.createMethod(m)
		}
	}

	return newContract
}

func (c *Contract) getSigner() client.Signer {
	if c.Signer != nil {
		return c.Signer
	}
	return c.client.Signer
}

func (c *Contract) LoadABI(abiJson []byte) error {
	iface, err := abi.Parse(abiJson)
	if err != nil {
		return err
	}
	for _, m := range iface.Methods {
		mm := m
		c.abiMethods[m.Name] = &mm
		if m.IsConstant {
			c.constantMethods[m.Name] = c.createConstantMethod(&mm)
		} else {
			c.methods[m.Name] = c.createMethod(&mm)
		}
	}
	for _, event := range iface.Events {
		ee := event
		if !event.IsAnonymous {
			c.eventSigMap[string(event.Sig)] = &ee
		}
		c.events[event.Name] = &ee
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
		OwnerAddress:    c.getSigner().Address(),
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
		return tt, tt.SignAndSend(ctx, c.getSigner())
	}
}

func (c *Contract) Call(ctx context.Context, methodName string, args ...any) ([]any, error) {
	m := c.constantMethods[methodName]
	if m == nil {
		return nil, ErrMethodNotFound
	}
	return m(ctx, args...)
}

func (c *Contract) Send(ctx context.Context, methodName string, args ...any) (*tx.Transaction, error) {
	m := c.methods[methodName]
	if m == nil {
		return nil, ErrMethodNotFound
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

func (c *Contract) getEventsByABIEvent(tx *tx.Transaction, ev *abi.Event) ([]Event, error) {
	var events []Event
	for _, log_ := range tx.Info.Log {
		if !bytes.Equal(log_.Address, c.address.ToEthAddress()) {
			continue
		}
		if !bytes.Equal(log_.Topics[0], ev.Sig) {
			continue
		}
		e, err := decodeEvent(ev, log_)
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
		return nil, ErrEventTypeNotFound
	}

	return c.getEventsByABIEvent(tx, ed)
}

func (c *Contract) GetResult(tx *tx.Transaction, methodName string) ([]any, error) {
	m := c.abiMethods[methodName]
	if m == nil {
		return nil, ErrMethodNotFound
	}
	if !tx.Confirmed || tx.Info == nil {
		return nil, fmt.Errorf("transaction not confirmed")
	}
	return m.OutputDecoder.Decode(tx.Info.ContractResult)
}
