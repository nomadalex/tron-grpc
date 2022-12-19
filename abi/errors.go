package abi

import "fmt"

var (
	ErrValueTypeNotSupport = fmt.Errorf("value type not support")
	ErrBytesSizeNotMatch   = fmt.Errorf("bytes size not match")
	ErrTypeError           = fmt.Errorf("type error")
	ErrTypeNotSupport      = fmt.Errorf("type not support")
)
