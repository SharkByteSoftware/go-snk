// Package jsonx provides helpers for decoding JSON from common sources.
package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Decode decodes JSON from an [io.Reader] into T.
// The caller is responsible for closing the reader if applicable.
//
// Returns an error if decoding fails.
func Decode[T any](r io.Reader, options ...DecodeOption) (*T, error) {
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

// DecodeBytes decodes JSON from a byte slice into T.
//
// Returns an error if decoding fails.
func DecodeBytes[T any](b []byte, options ...DecodeOption) (*T, error) {
	return Decode[T](bytes.NewReader(b), options...)
}

// DecodeString decodes JSON from a string into T.
//
// Returns an error if decoding fails.
func DecodeString[T any](s string, options ...DecodeOption) (*T, error) {
	return Decode[T](strings.NewReader(s), options...)
}

// DecodeFile decodes JSON from a file path into T.
//
// Returns an error if decoding fails.
func DecodeFile[T any](name string, options ...DecodeOption) (*T, error) {
	f, err := os.Open(filepath.Clean(name))
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	return Decode[T](f, options...)
}
