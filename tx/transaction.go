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
	SignTransaction(tx *Transaction) error
}

type Transaction struct {
	*core.Transaction

	client    api.WalletClient
	signer    Signer
	Confirmed bool
	Txid      []byte
	Return    *api.Return
	Info      *core.TransactionInfo
}

func (tx *Transaction) Send(ctx context.Context) error {
	err := tx.updateHash()
	if err != nil {
		return err
	}
	err = tx.signer.SignTransaction(tx)
	if err != nil {
		return err
	}
	tx.Return, err = tx.client.BroadcastTransaction(ctx, tx.Transaction)
	return err
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
		if info != nil {
			tx.Info = info
			return nil
		}
		if time.Now().After(timeout) {
			return fmt.Errorf("timeout")
		}
		time.Sleep(2 * time.Second)
	}
}

func New(client api.WalletClient, signer Signer, tx *core.Transaction) *Transaction {
	return &Transaction{
		client:      client,
		signer:      signer,
		Transaction: tx,
	}
}