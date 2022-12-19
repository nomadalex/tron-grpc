package contract

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strings"
)

type argumentEncoder []encoder
type argumentDecoder []decoder

func (e argumentEncoder) Encode(args []any) ([]byte, error) {
	ctx := &encodeContext{}
	for i, ee := range e {
		err := ee.Encode(ctx, args[i])
		if err != nil {
			return nil, err
		}
	}
	return ctx.Result(), nil
}

func (d argumentDecoder) Decode(result [][]byte) ([]any, error) {
	var args []any
	for i, dd := range d {
		ctx := &decodeContext{data: result[i]}
		v, err := dd.Decode(ctx)
		if err != nil {
			return nil, err
		}
		args = append(args, v)
	}
	return args, nil
}

func createArgumentEncoder(types []string) (argumentEncoder, error) {
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

func createArgumentDecoder(types []string) (argumentDecoder, error) {
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

type method struct {
	Name          string
	Sig           []byte
	inputEncoder  argumentEncoder
	outputDecoder argumentDecoder
}

type abiArguments struct {
	Name       string         `json:"name,omitempty"`
	Type       string         `json:"type,omitempty"`
	Components []abiArguments `json:"components,omitempty"`
	Indexed    bool           `json:"indexed,omitempty"`
}

type abiRecord struct {
	Type            string         `json:"type"`
	Name            string         `json:"name"`
	Inputs          []abiArguments `json:"inputs"`
	Outputs         []abiArguments `json:"outputs"`
	StateMutability string         `json:"stateMutability"`
	Anonymous       bool           `json:"anonymous,omitempty"`
}

func collectTypes(args []abiArguments) []string {
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

func parseAbi(jsonData []byte) ([]method, error) {
	var records []abiRecord
	err := json.Unmarshal(jsonData, &records)
	if err != nil {
		return nil, err
	}
	var methods []method
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

			methods = append(methods, method{
				Name:          r.Name,
				Sig:           calcFunctionSig(funcName),
				inputEncoder:  encoder,
				outputDecoder: decoder,
			})
		}
	}
	return methods, nil
}

func EncodeTypedData(types []string, data []any) ([]byte, error) {
	ctx := &encodeContext{}
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
	ctx := &decodeContext{data: data}
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
