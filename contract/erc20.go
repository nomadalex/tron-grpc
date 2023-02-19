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
	"math/big"
)

const erc20Abi = "[ { \"inputs\": [ { \"internalType\": \"string\", \"name\": \"name_\", \"type\": \"string\" }, { \"internalType\": \"string\", \"name\": \"symbol_\", \"type\": \"string\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"constructor\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"internalType\": \"address\", \"name\": \"owner\", \"type\": \"address\" }, { \"indexed\": true, \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Approval\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"internalType\": \"address\", \"name\": \"from\", \"type\": \"address\" }, { \"indexed\": true, \"internalType\": \"address\", \"name\": \"to\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Transfer\", \"type\": \"event\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"owner\", \"type\": \"address\" }, { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" } ], \"name\": \"allowance\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"approve\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\" } ], \"name\": \"balanceOf\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"decimals\", \"outputs\": [ { \"internalType\": \"uint8\", \"name\": \"\", \"type\": \"uint8\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"subtractedValue\", \"type\": \"uint256\" } ], \"name\": \"decreaseAllowance\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"addedValue\", \"type\": \"uint256\" } ], \"name\": \"increaseAllowance\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"name\", \"outputs\": [ { \"internalType\": \"string\", \"name\": \"\", \"type\": \"string\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"symbol\", \"outputs\": [ { \"internalType\": \"string\", \"name\": \"\", \"type\": \"string\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"totalSupply\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"recipient\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"transfer\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\" }, { \"internalType\": \"address\", \"name\": \"recipient\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"transferFrom\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" } ]"

type Erc20Event struct {
	Name string
	TransferEvent
	ApprovalEvent
}

func (e Erc20Event) String() string {
	switch e.Name {
	case "Transfer":
		return fmt.Sprint(e.Name, e.TransferEvent)
	case "Approval":
		return fmt.Sprint(e.Name, e.ApprovalEvent)
	default:
		return ""
	}
}

type TransferEvent struct {
	Address  address.Address
	From, To address.Address
	Value    *big.Int
}

type ApprovalEvent struct {
	Address        address.Address
	Owner, Spender address.Address
	Value          *big.Int
}

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

func (c *Erc20) Clone() *Erc20 {
	cc := c.contract.Clone()
	return &Erc20{
		contract:      cc,
		name:          cc.GetConstantMethod("name"),
		symbol:        cc.GetConstantMethod("symbol"),
		decimals:      cc.GetConstantMethod("decimals"),
		totalSupply:   cc.GetConstantMethod("totalSupply"),
		balanceOf:     cc.GetConstantMethod("balanceOf"),
		allowance:     cc.GetConstantMethod("allowance"),
		transfer:      cc.GetMethod("transfer"),
		approve:       cc.GetMethod("approve"),
		transferFrom:  cc.GetMethod("transferFrom"),
		transferEvent: cc.events["Transfer"],
		approvalEvent: cc.events["Approval"],
	}
}

func (c *Erc20) Signer() client.Signer {
	return c.contract.Signer
}

func (c *Erc20) SetSigner(signer client.Signer) {
	c.contract.Signer = signer
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

func (c *Erc20) Transfer(ctx context.Context, to address.Address, amount *big.Int, option *SendOption) (*tx.Transaction, error) {
	tx, err := c.transfer(ctx, to, amount, option)
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

func (c *Erc20) Approve(ctx context.Context, spender address.Address, amount *big.Int, option *SendOption) (*tx.Transaction, error) {
	tx, err := c.approve(ctx, spender, amount, option)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Erc20) TransferFrom(ctx context.Context, from, to address.Address, amount *big.Int, option *SendOption) (*tx.Transaction, error) {
	tx, err := c.transferFrom(ctx, from, to, amount, option)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Erc20) GetEvents(tx *tx.Transaction) ([]Erc20Event, error) {
	myAddr := c.contract.address.ToEthAddress()

	var events []Erc20Event
	parsers := []erc20EventParser{
		{
			sig: c.transferEvent.Sig,
			parser: func(log *core.TransactionInfo_Log) error {
				e, err := parseTransferEvent(log, c.contract.address, c.transferEvent.Decoder)
				if err != nil {
					return err
				}
				events = append(events, Erc20Event{Name: c.transferEvent.Name, TransferEvent: e})
				return nil
			},
		},
		{
			sig: c.approvalEvent.Sig,
			parser: func(log *core.TransactionInfo_Log) error {
				e, err := parseApproveEvent(log, c.contract.address, c.approvalEvent.Decoder)
				if err != nil {
					return err
				}
				events = append(events, Erc20Event{Name: c.approvalEvent.Name, ApprovalEvent: e})
				return nil
			},
		},
	}

	err := forEachLog(tx, myAddr, parsers)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (c *Erc20) GetTransferEvents(tx *tx.Transaction) ([]TransferEvent, error) {
	return parserEventWithABIEvent(tx, c.contract, c.transferEvent, parseTransferEvent)
}

func (c *Erc20) GetApprovalEvents(tx *tx.Transaction) ([]ApprovalEvent, error) {
	return parserEventWithABIEvent(tx, c.contract, c.approvalEvent, parseApproveEvent)
}

func parserEventWithABIEvent[E any](tx *tx.Transaction, c *Contract, ev *abi.Event, cb func(log *core.TransactionInfo_Log, addr address.Address, decoder *abi.EventDecoder) (E, error)) ([]E, error) {
	myAddr := c.address.ToEthAddress()

	var events []E
	parsers := []erc20EventParser{
		{
			sig: ev.Sig,
			parser: func(log *core.TransactionInfo_Log) error {
				e, err := cb(log, c.address, ev.Decoder)
				if err != nil {
					return err
				}
				events = append(events, e)
				return nil
			},
		},
	}

	err := forEachLog(tx, myAddr, parsers)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func parseTransferEvent(log_ *core.TransactionInfo_Log, addr address.Address, dec *abi.EventDecoder) (TransferEvent, error) {
	var err error
	var from, to any
	var data []any
	from, err = dec.DecodeTopic(0, log_.Topics[1])
	to, err = dec.DecodeTopic(1, log_.Topics[2])
	data, err = dec.DecodeData(log_.Data)
	if err != nil {
		return TransferEvent{}, err
	}

	return TransferEvent{
		Address: addr,
		From:    from.(address.Address),
		To:      to.(address.Address),
		Value:   data[0].(*big.Int),
	}, nil
}

func parseApproveEvent(log_ *core.TransactionInfo_Log, addr address.Address, dec *abi.EventDecoder) (ApprovalEvent, error) {
	var err error
	var owner, spender any
	var data []any
	owner, err = dec.DecodeTopic(0, log_.Topics[1])
	spender, err = dec.DecodeTopic(1, log_.Topics[2])
	data, err = dec.DecodeData(log_.Data)
	if err != nil {
		return ApprovalEvent{}, err
	}

	return ApprovalEvent{
		Address: addr,
		Owner:   owner.(address.Address),
		Spender: spender.(address.Address),
		Value:   data[0].(*big.Int),
	}, nil
}

type erc20EventParser struct {
	sig    []byte
	parser func(log *core.TransactionInfo_Log) error
}

func (p *erc20EventParser) CheckSig(sig []byte) bool {
	return bytes.Equal(p.sig, sig)
}

func (p *erc20EventParser) Parse(log *core.TransactionInfo_Log) error {
	return p.parser(log)
}

func forEachLog(tx *tx.Transaction, myAddr []byte, parsers []erc20EventParser) error {
	if len(parsers) == 1 {
		p := parsers[0]
		for _, log_ := range tx.Info.Log {
			if !bytes.Equal(log_.Address, myAddr) {
				continue
			}
			if p.CheckSig(log_.Topics[0]) {
				err := p.Parse(log_)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	for _, log_ := range tx.Info.Log {
		if !bytes.Equal(log_.Address, myAddr) {
			continue
		}
		for _, p := range parsers {
			if p.CheckSig(log_.Topics[0]) {
				err := p.Parse(log_)
				if err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}
