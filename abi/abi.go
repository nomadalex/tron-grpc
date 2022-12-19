package abi

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strings"
)

type ArgumentEncoder []encoder
type ArgumentDecoder []decoder

func (e ArgumentEncoder) Encode(args []any) ([]byte, error) {
	ctx := newEncodeContext()
	for i, ee := range e {
		err := ee.Encode(ctx, args[i])
		if err != nil {
			return nil, err
		}
	}
	return ctx.Result(), nil
}

func (d ArgumentDecoder) Decode(result [][]byte) ([]any, error) {
	var args []any
	for i, dd := range d {
		ctx := newDecodeContext(result[i])
		v, err := dd.Decode(ctx)
		if err != nil {
			return nil, err
		}
		args = append(args, v)
	}
	return args, nil
}

func createArgumentEncoder(types []string) (ArgumentEncoder, error) {
	var encoders []encoder
	for _, t := range types {
		e, err := createEncoder(t)
		if err != nil {
			return nil, err
		}
		encoders = append(encoders, e)
	}
	return encoders, nil
}

func createArgumentDecoder(types []string) (ArgumentDecoder, error) {
	var decoders []decoder
	for _, t := range types {
		e, err := createDecoder(t)
		if err != nil {
			return nil, err
		}
		decoders = append(decoders, e)
	}
	return decoders, nil
}

type Interface struct {
	Methods []Method
}

type Method struct {
	Name          string
	Sig           []byte
	InputEncoder  ArgumentEncoder
	OutputDecoder ArgumentDecoder
	IsConstant    bool
}

type arguments struct {
	Name       string      `json:"name,omitempty"`
	Type       string      `json:"type,omitempty"`
	Components []arguments `json:"components,omitempty"`
	Indexed    bool        `json:"indexed,omitempty"`
}

type record struct {
	Type            string      `json:"type"`
	Name            string      `json:"name"`
	Inputs          []arguments `json:"inputs"`
	Outputs         []arguments `json:"outputs"`
	StateMutability string      `json:"stateMutability"`
	Anonymous       bool        `json:"anonymous,omitempty"`
}

func collectTypes(args []arguments) []string {
	var ret []string
	for _, arg := range args {
		if strings.HasPrefix(arg.Type, "tuple") {
			types := collectTypes(arg.Components)
			ret = append(ret, fmt.Sprintf("(%s)%s", strings.Join(types, ","), arg.Type[5:]))
			continue
		}
		ret = append(ret, arg.Type)
	}
	return ret
}

func calcFunctionSig(funcDecl string) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(funcDecl))
	hash := h.Sum(nil)
	return hash[:4]
}

func Parse(jsonData []byte) (*Interface, error) {
	var records []record
	err := json.Unmarshal(jsonData, &records)
	if err != nil {
		return nil, err
	}
	var methods []Method
	for _, r := range records {
		if r.Type == "function" {
			inputTypes := collectTypes(r.Inputs)
			outputTypes := collectTypes(r.Outputs)
			funcName := fmt.Sprintf("%s(%s)", r.Name, strings.Join(inputTypes, ","))
			encoder, err := createArgumentEncoder(inputTypes)
			if err != nil {
				return nil, err
			}
			decoder, err := createArgumentDecoder(outputTypes)
			if err != nil {
				return nil, err
			}
			isConstant := r.StateMutability == "pure" || r.StateMutability == "view"

			methods = append(methods, Method{
				Name:          r.Name,
				Sig:           calcFunctionSig(funcName),
				InputEncoder:  encoder,
				OutputDecoder: decoder,
				IsConstant:    isConstant,
			})
		}
	}
	return &Interface{
		Methods: methods,
	}, nil
}

func EncodeTypedData(types []string, data []any) ([]byte, error) {
	ctx := newEncodeContext()
	for i, t := range types {
		e, err := createEncoder(t)
		if err != nil {
			return nil, err
		}
		err = e.Encode(ctx, data[i])
		if err != nil {
			return nil, err
		}
	}
	return ctx.Result(), nil
}

func DecodeTypedData(types []string, data []byte) ([]any, error) {
	var val []any
	ctx := newDecodeContext(data)
	for _, t := range types {
		d, err := createDecoder(t)
		if err != nil {
			return nil, err
		}
		v, err := d.Decode(ctx)
		if err != nil {
			return nil, err
		}
		val = append(val, v)
	}
	return val, nil
}
