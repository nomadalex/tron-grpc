package abi

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strings"
)

type InputEncoder struct {
	encoders []encoder
}
type OutputDecoder struct {
	decoders []decoder
}

func (e *InputEncoder) Encode(args []any) ([]byte, error) {
	if len(e.encoders) == 0 {
		return nil, nil
	}
	ctx := newEncodeContext()
	for i, ee := range e.encoders {
		err := ee.Encode(ctx, args[i])
		if err != nil {
			return nil, err
		}
	}
	return ctx.Result(), nil
}

func (d OutputDecoder) Decode(result [][]byte) ([]any, error) {
	if len(d.decoders) == 0 {
		return nil, nil
	}
	var args []any
	for i, dd := range d.decoders {
		ctx := newDecodeContext(result[i])
		v, err := dd.Decode(ctx)
		if err != nil {
			return nil, err
		}
		args = append(args, v)
	}
	return args, nil
}

func createArgumentEncoder(types []string) (*InputEncoder, error) {
	var encoders []encoder
	for _, t := range types {
		e, err := createEncoder(t)
		if err != nil {
			return nil, err
		}
		encoders = append(encoders, e)
	}
	return &InputEncoder{encoders: encoders}, nil
}

func createArgumentDecoder(types []string) (*OutputDecoder, error) {
	var decoders []decoder
	for _, t := range types {
		e, err := createDecoder(t)
		if err != nil {
			return nil, err
		}
		decoders = append(decoders, e)
	}
	return &OutputDecoder{decoders: decoders}, nil
}

type EventDecoder struct {
	topicsDecoders []decoder
	dataDecoders   []decoder
}

func (d *EventDecoder) DecodeAddr(addr []byte) (any, error) {
	return decodeAddress(addr)
}

func (d *EventDecoder) DecodeTopic(idx int, topic []byte) (any, error) {
	ctx := newDecodeContext(topic)
	return d.topicsDecoders[idx].Decode(ctx)
}

func (d *EventDecoder) DecodeTopics(topics [][]byte) ([]any, error) {
	var vals []any
	for i, topic := range topics {
		ctx := newDecodeContext(topic)
		v, err := d.topicsDecoders[i].Decode(ctx)
		if err != nil {
			return nil, err
		}
		vals = append(vals, v)
	}
	return vals, nil
}

func (d *EventDecoder) DecodeData(data []byte) ([]any, error) {
	var vals []any
	ctx := newDecodeContext(data)
	for _, dd := range d.dataDecoders {
		v, err := dd.Decode(ctx)
		if err != nil {
			return nil, err
		}
		vals = append(vals, v)
	}
	return vals, nil
}

type Interface struct {
	Methods []Method
	Events  []Event
}

type Method struct {
	Name          string
	Sig           []byte
	InputEncoder  *InputEncoder
	OutputDecoder *OutputDecoder
	IsConstant    bool
}

type EventInput struct {
	Name    string
	Indexed bool
}

type Event struct {
	Name        string
	IsAnonymous bool
	Sig         []byte
	Inputs      []EventInput
	Decoder     *EventDecoder
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
		if arg.Type == "int" {
			ret = append(ret, "int256")
			continue
		}
		if arg.Type == "uint" {
			ret = append(ret, "uint256")
			continue
		}
		ret = append(ret, arg.Type)
	}
	return ret
}

func GetKeccak256Hash(data []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(data)
	hash := h.Sum(nil)
	return hash
}

func calcFunctionSig(funcDecl string) []byte {
	return GetKeccak256Hash([]byte(funcDecl))[:4]
}

func calcEventSig(decl string) []byte {
	return GetKeccak256Hash([]byte(decl))
}

func parseFunction(r *record) (Method, error) {
	inputTypes := collectTypes(r.Inputs)
	outputTypes := collectTypes(r.Outputs)
	funcName := fmt.Sprintf("%s(%s)", r.Name, strings.Join(inputTypes, ","))
	encoder, err := createArgumentEncoder(inputTypes)
	if err != nil {
		return Method{}, err
	}
	decoder, err := createArgumentDecoder(outputTypes)
	if err != nil {
		return Method{}, err
	}
	isConstant := r.StateMutability == "pure" || r.StateMutability == "view"

	return Method{
		Name:          r.Name,
		Sig:           calcFunctionSig(funcName),
		InputEncoder:  encoder,
		OutputDecoder: decoder,
		IsConstant:    isConstant,
	}, nil
}

func parseEvent(r *record) (Event, error) {
	var inputs []EventInput
	var topicDecoders []decoder
	var dataDecoders []decoder
	inputTypes := collectTypes(r.Inputs)

	var sig []byte
	if !r.Anonymous {
		eventDecl := fmt.Sprintf("%s(%s)", r.Name, strings.Join(inputTypes, ","))
		sig = calcEventSig(eventDecl)
	}

	for i, input := range r.Inputs {
		inputs = append(inputs, EventInput{
			Name:    input.Name,
			Indexed: input.Indexed,
		})
		d, err := createDecoder(inputTypes[i])
		if err != nil {
			return Event{}, err
		}
		if input.Indexed {
			topicDecoders = append(topicDecoders, d)
		} else {
			dataDecoders = append(dataDecoders, d)
		}
	}
	return Event{
		Name:        r.Name,
		Sig:         sig,
		IsAnonymous: r.Anonymous,
		Inputs:      inputs,
		Decoder: &EventDecoder{
			topicsDecoders: topicDecoders,
			dataDecoders:   dataDecoders,
		},
	}, nil
}

func Parse(jsonData []byte) (*Interface, error) {
	var records []record
	err := json.Unmarshal(jsonData, &records)
	if err != nil {
		return nil, err
	}
	var methods []Method
	var events []Event
	for _, r := range records {
		if r.Type == "function" {
			m, err := parseFunction(&r)
			if err != nil {
				return nil, err
			}
			methods = append(methods, m)
		}
		if r.Type == "event" {
			e, err := parseEvent(&r)
			if err != nil {
				return nil, err
			}
			events = append(events, e)
		}
	}
	return &Interface{
		Methods: methods,
		Events:  events,
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
