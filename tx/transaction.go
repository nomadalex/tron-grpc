package tx

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/fullstackwang/tron-grpc/api"
	"github.com/fullstackwang/tron-grpc/core"
	"github.com/golang/protobuf/proto"
	"time"
)

type Signer interface {
	SignTransaction(tx *core.Transaction) ([]byte, error)
}

type ResultDecoder func([][]byte) ([]any, error)

type Transaction struct {
	*core.Transaction

	client        api.WalletClient
	Confirmed     bool
	Txid          []byte
	Info          *core.TransactionInfo
	resultDecoder ResultDecoder
}

func (tx *Transaction) Send(ctx context.Context, signer Signer) error {
	err := tx.updateHash()
	if err != nil {
		return err
	}
	sig, err := signer.SignTransaction(tx.Transaction)
	if err != nil {
		return err
	}
	tx.Signature = append(tx.Signature, sig)
	ret, err := tx.client.BroadcastTransaction(ctx, tx.Transaction)
	if err != nil {
		return err
	}
	if ret.Code > 0 {
		return fmt.Errorf(string(ret.Message))
	}
	return nil
}

func (tx *Transaction) updateHash() error {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	tx.Txid = h256h.Sum(nil)
	return nil
}

func (tx *Transaction) CheckConfirmation() error {
	in := api.BytesMessage{Value: tx.Txid}
	info, err := tx.client.GetTransactionInfoById(context.Background(), &in)
	if err != nil {
		return err
	}
	if info != nil && info.Id != nil {
		tx.Info = info
		tx.Confirmed = true
	}
	return nil
}

func (tx *Transaction) WaitConfirmation() error {
	in := api.BytesMessage{Value: tx.Txid}
	timeout := time.Now().Add(10 * time.Second)
	for {
		info, err := tx.client.GetTransactionInfoById(context.Background(), &in)
		if err != nil {
			return err
		}
		if info != nil && info.Id != nil {
			tx.Info = info
			tx.Confirmed = true
			return nil
		}
		if time.Now().After(timeout) {
			return fmt.Errorf("timeout")
		}
		time.Sleep(2 * time.Second)
	}
}

func (tx *Transaction) GetResult() ([]any, error) {
	if !tx.Confirmed {
		return nil, fmt.Errorf("tx not confirmed")
	}
	if tx.Info.Result > 0 {
		return nil, fmt.Errorf(string(tx.Info.ResMessage))
	}
	if tx.resultDecoder == nil {
		return nil, fmt.Errorf("no result decoder")
	}
	return tx.resultDecoder(tx.Info.ContractResult)
}

func New(client api.WalletClient, tx *core.Transaction) *Transaction {
	return &Transaction{
		client:      client,
		Transaction: tx,
	}
}

func NewWithInfo(client api.WalletClient, tx *core.Transaction, info *core.TransactionInfo) *Transaction {
	return &Transaction{
		Transaction: tx,
		client:      client,
		Confirmed:   info != nil,
		Txid:        info.Id,
		Info:        info,
	}
}

func NewWithDecoder(client api.WalletClient, tx *core.Transaction, resultDecoder ResultDecoder) *Transaction {
	return &Transaction{
		client:        client,
		Transaction:   tx,
		resultDecoder: resultDecoder,
	}
}

func GetFromID(ctx context.Context, client api.WalletClient, ID string, checkConfirm bool) (*Transaction, error) {
	hash, err := hex.DecodeString(ID)
	if err != nil {
		return nil, err
	}

	in := &api.BytesMessage{Value: hash}
	t, err := client.GetTransactionById(ctx, in)
	if err != nil {
		return nil, err
	}

	tx := New(client, t)
	tx.Txid = hash
	if checkConfirm {
		err = tx.CheckConfirmation()
		if err != nil {
			return nil, err
		}
	}
	return tx, nil
}
