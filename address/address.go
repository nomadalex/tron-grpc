package address

import (
	"encoding/hex"
	"math/big"
)

const (
	// Length is the expected length of the address
	Length = 21
	// LengthBase58 is the expected length of the address in base58format
	LengthBase58 = 34
	// LengthEthAddress is the expected length of the eth address
	LengthEthAddress = 20
	// TronBytePrefix is the hex prefix to address
	TronBytePrefix = byte(0x41)
)

// Address represents the 21 byte address of a tron account.
type Address []byte

// Bytes get bytes from address
func (a Address) Bytes() []byte {
	return a[:]
}

func (a Address) ToEthAddress() []byte {
	return a[1:]
}

// Hex get bytes from address in string
func (a Address) Hex() string {
	return hex.EncodeToString(a[:])
}

// String implements fmt.Stringer.
func (a Address) String() string {
	if a[0] == 0 {
		return new(big.Int).SetBytes(a.Bytes()).String()
	}
	return EncodeCheck(a)
}

func FromBytes(b []byte) (Address, error) {
	return b, nil
}

func FromHex(s string) (Address, error) {
	addr, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func FromHexUnsafe(s string) Address {
	a, _ := FromHex(s)
	return a
}

func FromBase58(s string) (Address, error) {
	addr, err := DecodeCheck(s)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func FromBase58Unsafe(s string) Address {
	a, _ := FromBase58(s)
	return a
}

func FromEthAddress(addr []byte) (Address, error) {
	a := make([]byte, 0, Length)
	a = append(a, TronBytePrefix)
	a = append(a, addr...)
	return a, nil
}
