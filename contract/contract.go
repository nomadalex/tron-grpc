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

type Event struct {
	Name string
}

type Contract struct {
	methods []abi.Method
	address address.Address
	client  *client.Client
}

func New(client *client.Client, addr address.Address) *Contract {
	return &Contract{
		address: addr,
		client:  client,
	}
}

func (c *Contract) LoadABI(abiJson []byte) error {
	var err error
	c.methods, err = abi.ParseMethods(abiJson)
	return err
}

func (c *Contract) getMethod(name string) *abi.Method {
	for _, m := range c.methods {
		if m.Name == name {
			return &m
		}
	}
	return nil
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

func (c *Contract) Call(ctx context.Context, methodName string, args ...any) ([]any, error) {
	m := c.getMethod(methodName)
	if m == nil {
		return nil, fmt.Errorf("method not found")
	}
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

func getSendOption(args []any) *SendOption {
	if len(args) == 0 {
		return nil
	}
	if option, ok := args[len(args)-1].(*SendOption); ok {
		return option
	}
	return nil
}

func (c *Contract) Send(ctx context.Context, methodName string, args ...any) (*tx.Transaction, error) {
	m := c.getMethod(methodName)
	if m == nil {
		return nil, fmt.Errorf("method not found")
	}

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

func (c *Contract) GetEvents(tx *tx.Transaction) ([]Event, error) {
	return nil, nil
}
