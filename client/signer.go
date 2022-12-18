package client

import (
	"github.com/fullstackwang/tron-grpc/address"
	"github.com/fullstackwang/tron-grpc/tx"
)

type Signer interface {
	Address() address.Address
	PubkeyString() string
	SignTransaction(tx *tx.Transaction) error
	SignMessage(msg string) ([]byte, error)
}
