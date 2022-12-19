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

type Event struct {
	Name string
}

type Contract struct {
	constantMethods map[string]ConstantMethod
	methods         map[string]Method
	address         address.Address
	client          *client.Client
}

func New(client *client.Client, addr address.Address) *Contract {
	return &Contract{
		address: addr,
		client:  client,
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

func (c *Contract) GetEvents(tx *tx.Transaction) ([]Event, error) {
	return nil, nil
}
