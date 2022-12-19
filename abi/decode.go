package abi

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
)

type decodeContext struct {
	data   []byte
	offset int
}

func (ctx *decodeContext) GoNext(offset int) {
	ctx.offset = ctx.offset + offset
}

func (ctx *decodeContext) ReadBool() (bool, error) {
	ret := ctx.data[ctx.offset+31] > 0
	ctx.GoNext(32)
	return ret, nil
}

func (ctx *decodeContext) ReadBigInt(hasSign bool) (*big.Int, error) {
	neg := false
	if hasSign {
		if ctx.data[ctx.offset] > 128 {
			neg = true
			data := ctx.data[ctx.offset : ctx.offset+32]
			for i, _ := range data {
				data[i] = ^data[i]
			}
			data[31] = data[31] + 1
		}
	}
	i := new(big.Int).SetBytes(ctx.data[ctx.offset : ctx.offset+32])
	if neg {
		i = i.Neg(i)
	}
	ctx.GoNext(32)
	return i, nil
}

func (ctx *decodeContext) ReadAddress() ([]byte, error) {
	addr := make([]byte, 20)
	copy(addr, ctx.data[ctx.offset+12:ctx.offset+32])
	return addr, nil
}

func (ctx *decodeContext) ReadLen() int {
	l := int(binary.BigEndian.Uint32(ctx.data[ctx.offset+28 : ctx.offset+32]))
	ctx.GoNext(32)
	return l
}

func (ctx *decodeContext) ReadDynamicIndex() int {
	l := int(binary.BigEndian.Uint32(ctx.data[ctx.offset+28 : ctx.offset+32]))
	ctx.GoNext(32)
	return l
}

func (ctx *decodeContext) GetDynamicBuf(index int) []byte {
	return ctx.data[index:]
}

func (ctx *decodeContext) RemainingBytes() []byte {
	return ctx.data[ctx.offset:]
}

type decoder interface {
	IsDynamic() bool
	Decode(ctx *decodeContext) (any, error)
}

type numDecoder struct {
	hasSign bool
}

func (d *numDecoder) IsDynamic() bool {
	return false
}

func (d *numDecoder) Decode(ctx *decodeContext) (any, error) {
	return ctx.ReadBigInt(d.hasSign)
}

type boolDecoder struct{}

func (d *boolDecoder) IsDynamic() bool {
	return false
}

func (d *boolDecoder) Decode(ctx *decodeContext) (any, error) {
	return ctx.ReadBool()
}

type addressDecoder struct{}

func (d *addressDecoder) IsDynamic() bool {
	return false
}

func (d *addressDecoder) Decode(ctx *decodeContext) (any, error) {
	addr, err := ctx.ReadAddress()
	if err != nil {
		return nil, err
	}
	return decodeAddress(addr)
}

func getDecodeContext(ctx *decodeContext, isDyn bool, size int) (*decodeContext, int) {
	cc := ctx
	if isDyn {
		index := ctx.ReadDynamicIndex()
		cc = &decodeContext{data: ctx.GetDynamicBuf(index)}
	}
	if size < 0 {
		size = cc.ReadLen()
	}
	return ctx, size
}

func decodeBytes(ctx *decodeContext, size int) ([]byte, error) {
	var cc *decodeContext
	isDyn := size < 0
	cc, size = getDecodeContext(ctx, isDyn, size)
	buf := make([]byte, size)
	copy(buf, cc.RemainingBytes())
	if !isDyn {
		cc.GoNext(size + 32 - (size % 32))
	}
	return buf, nil
}

type bytesDecoder struct {
	size int
}

func (d *bytesDecoder) IsDynamic() bool {
	return d.size < 0
}

func (d *bytesDecoder) Decode(ctx *decodeContext) (any, error) {
	return decodeBytes(ctx, d.size)
}

type stringDecoder struct{}

func (d *stringDecoder) IsDynamic() bool {
	return true
}

func (d *stringDecoder) Decode(ctx *decodeContext) (any, error) {
	b, err := decodeBytes(ctx, -1)
	return string(b), err
}

type tupleDecoder struct {
	isDynamic   bool
	subDecoders []decoder
}

func (d *tupleDecoder) IsDynamic() bool {
	return d.isDynamic
}

func (d *tupleDecoder) Decode(ctx *decodeContext) (any, error) {
	cc, _ := getDecodeContext(ctx, d.IsDynamic(), len(d.subDecoders))
	var arr []any
	for _, dd := range d.subDecoders {
		v, err := dd.Decode(cc)
		if err != nil {
			return nil, err
		}
		arr = append(arr, v)
	}
	return arr, nil
}

type arrayDecoder struct {
	size       int
	subDecoder decoder
}

func (d *arrayDecoder) IsDynamic() bool {
	return d.size < 0 || d.subDecoder.IsDynamic()
}

func (d *arrayDecoder) Decode(ctx *decodeContext) (any, error) {
	cc, size := getDecodeContext(ctx, d.IsDynamic(), d.size)
	var arr []any
	for i := 0; i < size; i++ {
		v, err := d.subDecoder.Decode(cc)
		if err != nil {
			return nil, err
		}
		arr = append(arr, v)
	}
	return arr, nil
}

func createBasicDecoder(type_ string) (decoder, error) {
	if type_ == "bool" {
		return &boolDecoder{}, nil
	}
	if type_ == "string" {
		return &stringDecoder{}, nil
	}
	if type_ == "address" {
		return &addressDecoder{}, nil
	}
	if strings.HasPrefix(type_, "uint") {
		return &numDecoder{hasSign: false}, nil
	}
	if strings.HasPrefix(type_, "int") {
		return &numDecoder{hasSign: true}, nil
	}
	if strings.HasPrefix(type_, "bytes") {
		size, err := parseSizeStr(type_[5:])
		if err != nil {
			return nil, err
		}
		return &bytesDecoder{size: size}, nil
	}
	return nil, fmt.Errorf("type not support")
}

func createDecoder(t string) (decoder, error) {
	size, types, err := parseComplexType(t)
	if err != nil {
		return nil, err
	}
	if size > 0 {
		if len(types) == 1 {
			subDecoder, err := createDecoder(types[0])
			if err != nil {
				return nil, err
			}
			return &arrayDecoder{
				size:       size,
				subDecoder: subDecoder,
			}, nil
		}

		isDyn := false
		var subDecoders []decoder
		for _, elem := range types {
			e, err := createDecoder(elem)
			if err != nil {
				return nil, err
			}
			isDyn = isDyn || e.IsDynamic()
			subDecoders = append(subDecoders, e)
		}
		return &tupleDecoder{
			isDynamic:   isDyn,
			subDecoders: subDecoders,
		}, nil
	}
	return createBasicDecoder(t)
}
