package client

import (
	"github.com/fullstackwang/tron-grpc/address"
	"github.com/fullstackwang/tron-grpc/core"
)

type Signer interface {
	Address() address.Address
	PublicKey() []byte
	SignTransaction(tx *core.Transaction) error
	SignMessage(msg string) ([]byte, error)
}
