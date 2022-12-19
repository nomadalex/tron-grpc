package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/dustinxie/ecc"
	"github.com/fullstackwang/tron-grpc/address"
	"golang.org/x/crypto/sha3"
	"math/big"
)

type Wallet struct {
	privKey *ecdsa.PrivateKey
	address address.Address
}

func (w *Wallet) Address() address.Address {
	return w.address
}

func FromPrivateKey(privateKey string) (*Wallet, error) {
	data, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	if len(data) != 32 {
		return nil, fmt.Errorf("size error")
	}

	p := privKeyFromBytes(data)
	addr := genAddressFromPrivKey(p)
	return &Wallet{
		privKey: p,
		address: addr,
	}, nil
}

func genAddressFromPrivKey(p *ecdsa.PrivateKey) address.Address {
	pubBytes := make([]byte, 0, 64)
	pubBytes = append(pubBytes, p.X.Bytes()...)
	pubBytes = append(pubBytes, p.Y.Bytes()...)

	s := sha3.NewLegacyKeccak256()
	s.Write(pubBytes)
	hash := s.Sum(nil)
	addr, _ := address.FromEthAddress(hash[12:])
	return addr
}

func privKeyFromBytes(data []byte) *ecdsa.PrivateKey {
	p256k1 := ecc.P256k1()
	priv := new(ecdsa.PrivateKey)
	priv.D = new(big.Int)
	priv.D.SetBytes(data)
	priv.Curve = p256k1
	priv.X, priv.Y = p256k1.ScalarBaseMult(data)
	return priv
}

func (w *Wallet) SignTransactionHash(txHash []byte) ([]byte, error) {
	return ecc.SignEthereum(txHash, w.privKey)
}

func (w *Wallet) SignMessage(msg string) ([]byte, error) {
	msg = fmt.Sprintf("\x19Tron Signed Message:\n%d%s", len(msg), msg)
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(msg))
	hash := h.Sum(nil)
	return ecc.SignEthereum(hash, w.privKey)
}
