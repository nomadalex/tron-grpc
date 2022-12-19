package tx

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/fullstackwang/tron-grpc/api"
	"github.com/fullstackwang/tron-grpc/core"
	"github.com/golang/protobuf/proto"
	"time"
)

type Signer interface {
	SignTransactionHash(txHash []byte) ([]byte, error)
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
	sig, err := signer.SignTransactionHash(tx.Txid)
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

func NewWithDecoder(client api.WalletClient, tx *core.Transaction, resultDecoder ResultDecoder) *Transaction {
	return &Transaction{
		client:        client,
		Transaction:   tx,
		resultDecoder: resultDecoder,
	}
}
