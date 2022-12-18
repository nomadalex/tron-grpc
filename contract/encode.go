package contract

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

type dynamicRef struct {
	headOffset    int
	dynamicOffset int
}

type encodeContext struct {
	buf        bytes.Buffer
	dynamicBuf bytes.Buffer
	dynRefs    []dynamicRef
}

func (ctx *encodeContext) AddDynamicRef() {
	ctx.dynRefs = append(ctx.dynRefs, dynamicRef{
		headOffset:    ctx.GetHeadOffset(),
		dynamicOffset: ctx.GetDynamicOffset(),
	})
	ctx.buf.Write(make([]byte, 32))
}

func (ctx *encodeContext) GetHeadOffset() int {
	return ctx.buf.Len()
}

func (ctx *encodeContext) GetDynamicOffset() int {
	return ctx.dynamicBuf.Len()
}

func (ctx *encodeContext) getBuf(dynamic bool) *bytes.Buffer {
	if dynamic {
		return &ctx.dynamicBuf
	} else {
		return &ctx.buf
	}
}

func (ctx *encodeContext) WriteBigInt(i *big.Int, negPad bool, dynamic bool) {
	head := make([]byte, 32)
	putBigInt(head, i, negPad)
	ctx.getBuf(dynamic).Write(head)
}

func (ctx *encodeContext) WriteBytes(b []byte, pad, dynamic bool) {
	buf := ctx.getBuf(dynamic)
	buf.Write(b)
	if pad {
		rem := len(b) % 32
		if rem != 0 {
			pad := make([]byte, 32-rem)
			buf.Write(pad)
		}
	}
}

func (ctx *encodeContext) updateDynamicHead(d dynamicRef, headLen int) {
	head := ctx.buf.Bytes()[d.headOffset : d.headOffset+32]
	putBigInt(head, big.NewInt(int64(headLen+d.dynamicOffset)), false)
}

func (ctx *encodeContext) Result() []byte {
	headLen := ctx.buf.Len()
	for _, d := range ctx.dynRefs {
		ctx.updateDynamicHead(d, headLen)
	}
	ctx.buf.Write(ctx.dynamicBuf.Bytes())
	return ctx.buf.Bytes()
}

func putBigInt(buf []byte, i *big.Int, negPad bool) {
	bytes := i.Bytes()

	if len(bytes) < len(buf) {
		idx := len(buf) - len(bytes)
		copy(buf[idx:], bytes)
		if negPad && i.Sign() < 0 {
			buf[31] = buf[31] - 1
			for i := 0; i < 32; i++ {
				buf[i] = ^buf[i]
			}
		}
	}
}

type encoder interface {
	IsDynamic() bool
	Encode(ctx *encodeContext, val any) error
}

type numEncoder struct {
	hasSign bool
}

func (e *numEncoder) IsDynamic() bool {
	return false
}

func (e *numEncoder) Encode(ctx *encodeContext, val any) error {
	var i *big.Int
	switch val.(type) {
	case int:
		i = big.NewInt(int64(val.(int)))
	case *big.Int:
		i = val.(*big.Int)
	default:
		return ErrValueTypeNotSupport
	}

	ctx.WriteBigInt(i, e.hasSign, false)
	return nil
}

type addressEncoder struct{}

func (e *addressEncoder) IsDynamic() bool {
	return false
}

func (e *addressEncoder) Encode(ctx *encodeContext, val any) error {
	var b []byte
	var err error
	switch val.(type) {
	case string:
		b, err = hex.DecodeString(val.(string))
		if err != nil {
			return err
		}
	case []byte:
		b = val.([]byte)
	default:
		return ErrValueTypeNotSupport
	}

	head := make([]byte, 32)
	if len(b) < 32 {
		copy(head[32-len(b):], b)
	}
	ctx.buf.Write(head)
	return nil
}

type boolEncoder struct{}

func (e *boolEncoder) IsDynamic() bool {
	return false
}

func (e *boolEncoder) Encode(ctx *encodeContext, val any) error {
	if v, ok := val.(bool); ok {
		head := make([]byte, 32)
		if v {
			head[31] = 1
		}
		ctx.buf.Write(head)
		return nil
	}
	return ErrValueTypeNotSupport
}

func encodeBytes(ctx *encodeContext, v []byte) {
	ctx.AddDynamicRef()
	ctx.WriteBigInt(big.NewInt(int64(len(v))), false, true)
	ctx.WriteBytes(v, true, true)
}

type bytesEncoder struct {
	size int
}

func (e *bytesEncoder) IsDynamic() bool {
	return e.size < 0
}

func (e *bytesEncoder) Encode(ctx *encodeContext, val any) error {
	if v, ok := val.([]byte); ok {
		if e.size < 0 {
			encodeBytes(ctx, v)
			return nil
		}
		if len(v) != e.size {
			return ErrBytesSizeNotMatch
		}
		ctx.WriteBytes(v, true, false)
		return nil
	}
	return ErrValueTypeNotSupport
}

type stringEncoder struct{}

func (e *stringEncoder) IsDynamic() bool {
	return true
}

func (e *stringEncoder) Encode(ctx *encodeContext, val any) error {
	if v, ok := val.(string); ok {
		encodeBytes(ctx, []byte(v))
		return nil
	}
	return ErrValueTypeNotSupport
}

type tupleEncoder struct {
	isDynamic   bool
	subEncoders []encoder
}

func (e *tupleEncoder) IsDynamic() bool {
	return e.isDynamic
}

func (e *tupleEncoder) Encode(ctx *encodeContext, val any) error {
	return encodeDynamic(ctx, val, e.IsDynamic(), e.subEncoders)
}

type arrayEncoder struct {
	size       int
	subEncoder encoder
}

func (e *arrayEncoder) IsDynamic() bool {
	return e.size < 0 || e.subEncoder.IsDynamic()
}

func (e *arrayEncoder) Encode(ctx *encodeContext, val any) error {
	return encodeDynamic(ctx, val, e.IsDynamic(), []encoder{e.subEncoder})
}

func encodeDynamic(ctx *encodeContext, val any, isDyn bool, encoders []encoder) error {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Slice {
		return ErrValueTypeNotSupport
	}
	cc := ctx
	if isDyn {
		ctx.AddDynamicRef()
		ctx.WriteBigInt(big.NewInt(int64(v.Len())), false, true)
		cc = &encodeContext{}
	}
	getEncoder := func(idx int) encoder {
		if len(encoders) == 1 {
			return encoders[0]
		}
		return encoders[idx]
	}
	for i := 0; i < v.Len(); i++ {
		err := getEncoder(i).Encode(cc, v.Index(i).Interface())
		if err != nil {
			return err
		}
	}
	if isDyn {
		ctx.WriteBytes(cc.Result(), false, true)
	}
	return nil
}

func parseSizeStr(str string) (int, error) {
	if str == "" {
		return -1, nil
	}
	size, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func splitTupleElem(t string) ([]string, error) {
	var elems []string
	opens := 0
	lastIdx := 0
	for i := lastIdx; i < len(t); i++ {
		v := t[i]
		switch v {
		case '(':
			opens++
		case ')':
			opens--
		case ',':
			if opens == 0 {
				elems = append(elems, t[lastIdx:i])
				lastIdx = i + 1
			}
		}
	}
	if opens > 0 {
		return nil, ErrTypeError
	}
	elems = append(elems, t[lastIdx:])
	return elems, nil
}

func parseComplexType(t string) (int, []string, error) {
	if t[len(t)-1:] == "]" {
		idx := strings.LastIndex(t, "[")
		size, err := parseSizeStr(t[idx+1 : len(t)-1])
		if err != nil {
			return 0, nil, err
		}
		subType := t[:idx]
		return size, []string{subType}, nil
	}
	if t[:1] == "(" {
		if t[len(t)-1:] != ")" {
			return 0, nil, ErrTypeError
		}
		elems, err := splitTupleElem(t[1 : len(t)-1])
		if err != nil {
			return 0, nil, err
		}
		return len(elems), elems, nil
	}
	return 0, nil, nil
}

func createBasicEncoder(type_ string) (encoder, error) {
	if type_ == "bool" {
		return &boolEncoder{}, nil
	}
	if type_ == "string" {
		return &stringEncoder{}, nil
	}
	if type_ == "address" {
		return &addressEncoder{}, nil
	}
	if strings.HasPrefix(type_, "uint") {
		return &numEncoder{hasSign: false}, nil
	}
	if strings.HasPrefix(type_, "int") {
		return &numEncoder{hasSign: true}, nil
	}
	if strings.HasPrefix(type_, "bytes") {
		size, err := parseSizeStr(type_[5:])
		if err != nil {
			return nil, err
		}
		return &bytesEncoder{size: size}, nil
	}
	return nil, ErrTypeNotSupport
}

func createEncoder(t string) (encoder, error) {
	size, types, err := parseComplexType(t)
	if err != nil {
		return nil, err
	}
	if size > 0 {
		if len(types) == 1 {
			subEncoder, err := createEncoder(types[0])
			if err != nil {
				return nil, err
			}
			return &arrayEncoder{
				size:       size,
				subEncoder: subEncoder,
			}, nil
		}

		isDyn := false
		var subEncoders []encoder
		for _, elem := range types {
			e, err := createEncoder(elem)
			if err != nil {
				return nil, err
			}
			isDyn = isDyn || e.IsDynamic()
			subEncoders = append(subEncoders, e)
		}
		return &tupleEncoder{
			isDynamic:   isDyn,
			subEncoders: subEncoders,
		}, nil
	}
	return createBasicEncoder(t)
}
