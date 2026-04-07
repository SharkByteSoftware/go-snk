package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Encode encodes value as JSON into the provided [io.Writer].
//
// Returns an error if encoding fails.
func Encode[T any](writer io.Writer, value T, options ...EncodeOption) error {
	cfg := newEncodeOptions(options)

	enc := json.NewEncoder(writer)
	if cfg.escapeHTML {
		enc.SetEscapeHTML(true)
	} else {
		enc.SetEscapeHTML(false)
	}

	if cfg.indent != "" {
		enc.SetIndent("", cfg.indent)
	}

	err := enc.Encode(value)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}

// EncodeBytes encodes value as JSON and returns the result as a byte slice.
//
// Returns an error if encoding fails.
func EncodeBytes[T any](value T, options ...EncodeOption) ([]byte, error) {
	var buf bytes.Buffer

	err := Encode[T](&buf, value, options...)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// EncodeString encodes value as JSON and returns the result as a string.
//
// Returns an error if encoding fails.
func EncodeString[T any](value T, options ...EncodeOption) (string, error) {
	b, err := EncodeBytes[T](value, options...)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// EncodeToFile encodes value as JSON into the file at the given path.
//
// Returns an error if encoding fails.
func EncodeToFile[T any](path string, value T, options ...EncodeOption) error {
	file, err := os.Create(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer func() { _ = file.Close() }()

	return Encode(file, value, options...)
}
