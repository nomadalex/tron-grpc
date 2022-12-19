package address

import (
	"encoding/hex"
	"math/big"
)

const (
	// AddressLength is the expected length of the address
	AddressLength = 21
	// AddressLengthBase58 is the expected length of the address in base58format
	AddressLengthBase58 = 34
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

func BytesToAddress(b []byte) Address {
	return b
}

// HexToAddress returns Address with byte values of s.
// If s is larger than len(h), s will be cropped from the left.
func HexToAddress(s string) Address {
	addr, err := hex.DecodeString(s)
	if err != nil {
		return nil
	}
	return addr
}

// Base58ToAddress returns Address with byte values of s.
func Base58ToAddress(s string) (Address, error) {
	addr, err := DecodeCheck(s)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func FromEthAddress(addr []byte) Address {
	a := make([]byte, 0, AddressLength)
	a = append(a, TronBytePrefix)
	a = append(a, addr...)
	return a
}
