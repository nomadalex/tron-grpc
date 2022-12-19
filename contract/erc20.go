package contract

import (
	"context"
	"github.com/fullstackwang/tron-grpc/abi"
	"github.com/fullstackwang/tron-grpc/address"
	"github.com/fullstackwang/tron-grpc/client"
	"github.com/fullstackwang/tron-grpc/tx"
	"math/big"
)

const erc20Abi = "[ { \"inputs\": [ { \"internalType\": \"string\", \"name\": \"name_\", \"type\": \"string\" }, { \"internalType\": \"string\", \"name\": \"symbol_\", \"type\": \"string\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"constructor\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"internalType\": \"address\", \"name\": \"owner\", \"type\": \"address\" }, { \"indexed\": true, \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Approval\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"internalType\": \"address\", \"name\": \"from\", \"type\": \"address\" }, { \"indexed\": true, \"internalType\": \"address\", \"name\": \"to\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Transfer\", \"type\": \"event\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"owner\", \"type\": \"address\" }, { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" } ], \"name\": \"allowance\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"approve\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\" } ], \"name\": \"balanceOf\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"decimals\", \"outputs\": [ { \"internalType\": \"uint8\", \"name\": \"\", \"type\": \"uint8\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"subtractedValue\", \"type\": \"uint256\" } ], \"name\": \"decreaseAllowance\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"addedValue\", \"type\": \"uint256\" } ], \"name\": \"increaseAllowance\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"name\", \"outputs\": [ { \"internalType\": \"string\", \"name\": \"\", \"type\": \"string\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"symbol\", \"outputs\": [ { \"internalType\": \"string\", \"name\": \"\", \"type\": \"string\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"totalSupply\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"recipient\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"transfer\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\" }, { \"internalType\": \"address\", \"name\": \"recipient\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"transferFrom\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" } ]"

type Erc20 struct {
	contract *Contract

	name,
	symbol,
	decimals,
	totalSupply,
	balanceOf,
	allowance ConstantMethod

	transfer,
	approve,
	transferFrom Method

	transferEvent,
	approvalEvent *abi.Event
}

func NewErc20(client *client.Client, addr address.Address) *Erc20 {
	c := New(client, addr)
	_ = c.LoadABI([]byte(erc20Abi))
	return &Erc20{
		contract:      c,
		name:          c.GetConstantMethod("name"),
		symbol:        c.GetConstantMethod("symbol"),
		decimals:      c.GetConstantMethod("decimals"),
		totalSupply:   c.GetConstantMethod("totalSupply"),
		balanceOf:     c.GetConstantMethod("balanceOf"),
		allowance:     c.GetConstantMethod("allowance"),
		transfer:      c.GetMethod("transfer"),
		approve:       c.GetMethod("approve"),
		transferFrom:  c.GetMethod("transferFrom"),
		transferEvent: c.events["Transfer"],
		approvalEvent: c.events["Approval"],
	}
}

func (c *Erc20) Name(ctx context.Context) (string, error) {
	ret, err := c.name(ctx)
	if err != nil {
		return "", err
	}
	return ret[0].(string), nil
}

func (c *Erc20) Symbol(ctx context.Context) (string, error) {
	ret, err := c.symbol(ctx)
	if err != nil {
		return "", err
	}
	return ret[0].(string), nil
}

func (c *Erc20) Decimals(ctx context.Context) (int, error) {
	ret, err := c.decimals(ctx)
	if err != nil {
		return 0, err
	}
	i := ret[0].(*big.Int)
	return int(i.Uint64()), nil
}

func (c *Erc20) TotalSupply(ctx context.Context) (*big.Int, error) {
	ret, err := c.totalSupply(ctx)
	if err != nil {
		return nil, err
	}
	i := ret[0].(*big.Int)
	return i, nil
}

func (c *Erc20) BalanceOf(ctx context.Context, addr address.Address) (*big.Int, error) {
	ret, err := c.balanceOf(ctx, addr)
	if err != nil {
		return nil, err
	}
	i := ret[0].(*big.Int)
	return i, nil
}

func (c *Erc20) Transfer(ctx context.Context, to address.Address, amount *big.Int) (*tx.Transaction, error) {
	tx, err := c.transfer(ctx, to, amount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Erc20) Allowance(ctx context.Context, owner, spender address.Address) (*big.Int, error) {
	ret, err := c.allowance(ctx, owner, spender)
	if err != nil {
		return nil, err
	}
	i := ret[0].(*big.Int)
	return i, nil
}

func (c *Erc20) Approve(ctx context.Context, spender address.Address, amount *big.Int) (*tx.Transaction, error) {
	tx, err := c.approve(ctx, spender, amount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Erc20) TransferFrom(ctx context.Context, from, to address.Address, amount *big.Int) (*tx.Transaction, error) {
	tx, err := c.transferFrom(ctx, from, to, amount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Erc20) GetEvents(tx *tx.Transaction) ([]Event, error) {
	return c.contract.GetEvents(tx)
}

func (c *Erc20) GetTransferEvents(tx *tx.Transaction) ([]Event, error) {
	return c.contract.getEventsByABIEvent(tx, c.transferEvent)
}

func (c *Erc20) GetApprovalEvents(tx *tx.Transaction) ([]Event, error) {
	return c.contract.getEventsByABIEvent(tx, c.approvalEvent)
}
