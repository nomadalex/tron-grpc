package client

import (
	"github.com/fullstackwang/tron-grpc/address"
)

type Signer interface {
	Address() address.Address
	PubkeyString() string
	SignTransaction(txHash []byte) ([]byte, error)
	SignMessage(msg string) ([]byte, error)
}
