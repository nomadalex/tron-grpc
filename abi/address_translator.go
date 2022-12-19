package abi

import "encoding/hex"

type AddressTranslator interface {
	FromEthAddress(addr []byte) (any, error)
	ToEthAddress(addr any) ([]byte, error)
}

var CustomAddressTranslator AddressTranslator

type defaultAddressTranslator struct{}

func (d defaultAddressTranslator) FromEthAddress(addr []byte) (any, error) {
	return addr, nil
}

func (d defaultAddressTranslator) ToEthAddress(val any) ([]byte, error) {
	switch v := val.(type) {
	case string:
		b, err := hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
		return b, nil
	case []byte:
		return v, nil
	default:
		return nil, ErrValueTypeNotSupport
	}
}

func decodeAddress(addr []byte) (any, error) {
	if CustomAddressTranslator == nil {
		return defaultAddressTranslator{}.FromEthAddress(addr)
	}
	return CustomAddressTranslator.FromEthAddress(addr)
}

func encodeAddress(addr any) ([]byte, error) {
	if CustomAddressTranslator == nil {
		return defaultAddressTranslator{}.ToEthAddress(addr)
	}
	return CustomAddressTranslator.ToEthAddress(addr)
}
