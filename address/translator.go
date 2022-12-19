package address

import (
	"github.com/fullstackwang/tron-grpc/abi"
)

type addressTranslator struct{}

func (d addressTranslator) FromEthAddress(addr []byte) (any, error) {
	return FromEthAddress(addr)
}

func (d addressTranslator) ToEthAddress(val any) ([]byte, error) {
	switch v := val.(type) {
	case string:
		b, err := FromHex(v)
		if err != nil {
			return nil, err
		}
		return b.ToEthAddress(), nil
	case []byte:
		b, err := FromBytes(v)
		if err != nil {
			return nil, err
		}
		return b.ToEthAddress(), nil
	case Address:
		return v.ToEthAddress(), nil
	default:
		return nil, abi.ErrValueTypeNotSupport
	}
}

func init() {
	abi.CustomAddressTranslator = &addressTranslator{}
}
