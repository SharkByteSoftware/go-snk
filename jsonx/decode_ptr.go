package jsonx

import (
	"encoding/json"
	"fmt"
	"io"
)

func DecodePtr[T any](r io.Reader, options ...DecodeOption) (*T, error) {
	cfg := newDecodeOptions(options)

	dec := json.NewDecoder(r)
	if cfg.strictDecoding {
		dec.DisallowUnknownFields()
	}

	if cfg.useNumber {
		dec.UseNumber()
	}

	var result T

	err := dec.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return &result, nil
}
