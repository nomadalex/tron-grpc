package wallet

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dustinxie/ecc"
	"github.com/fullstackwang/tron-grpc/address"
	"github.com/fullstackwang/tron-grpc/core"
	"github.com/golang/protobuf/proto"
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

func (w *Wallet) PublicKey() []byte {
	pubBytes := make([]byte, 0, 65)
	pubBytes = append(pubBytes, 0x04) // uncompressed
	pubBytes = append(pubBytes, w.privKey.X.Bytes()...)
	pubBytes = append(pubBytes, w.privKey.Y.Bytes()...)
	return pubBytes
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

func (w *Wallet) SignTransaction(tx *core.Transaction) error {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	sig, err := ecc.SignEthereum(hash, w.privKey)
	if err != nil {
		return err
	}
	tx.Signature = append(tx.Signature, sig)
	return nil
}

func (w *Wallet) SignMessage(msg string) ([]byte, error) {
	msg = fmt.Sprintf("\x19Tron Signed Message:\n%d%s", len(msg), msg)
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(msg))
	hash := h.Sum(nil)
	return ecc.SignEthereum(hash, w.privKey)
}
