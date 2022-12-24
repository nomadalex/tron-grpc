package abi

import "fmt"

type GenericEncoder struct {
	encoders []encoder
}

type GenericDecoder struct {
	decoders []decoder
}

func CreateGenericEncoder(types []string) (enc *GenericEncoder, err error) {
	encoders := make([]encoder, len(types))
	for i, t := range types {
		encoders[i], err = createEncoder(t)
		if err != nil {
			return
		}
	}
	enc = &GenericEncoder{encoders: encoders}
	return
}

func CreateGenericDecoder(types []string) (dec *GenericDecoder, err error) {
	decoders := make([]decoder, len(types))
	for i, t := range types {
		decoders[i], err = createDecoder(t)
		if err != nil {
			return
		}
	}
	dec = &GenericDecoder{decoders: decoders}
	return
}

func (e *GenericEncoder) EncodeAll(data []any) (ret []byte, err error) {
	ctx := newEncodeContext()
	for i, ee := range e.encoders {
		err = ee.Encode(ctx, data[i])
		if err != nil {
			return
		}
	}
	ret = ctx.Result()
	return
}

func (e *GenericEncoder) EncodeSingle(data any) ([]byte, error) {
	if len(e.encoders) != 1 {
		return nil, fmt.Errorf("EncodeSingle only allow one encoder exist")
	}
	ctx := newEncodeContext()
	err := e.encoders[0].Encode(ctx, data)
	if err != nil {
		return nil, err
	}
	return ctx.Result(), nil
}

func (e *GenericDecoder) DecodeAll(data []byte) (ret []any, err error) {
	ctx := newDecodeContext(data)
	ret = make([]any, len(e.decoders))
	for i, d := range e.decoders {
		ret[i], err = d.Decode(ctx)
		if err != nil {
			return
		}
	}
	return
}

func (e *GenericDecoder) DecodeSingle(data []byte) (any, error) {
	if len(e.decoders) != 1 {
		return nil, fmt.Errorf("DecodeSingle only allow one decoder exist")
	}
	ctx := newDecodeContext(data)
	return e.decoders[0].Decode(ctx)
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
