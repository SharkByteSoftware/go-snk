package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response represents the structure for an HTTP response, supporting generics for handling various result types.
type Response[T any] struct {
	StatusCode int
	Status     string
	Header     http.Header
	Result     *T
	RawBody    []byte
}

// DecodeRawBody decodes the raw body of an HTTP response into the specified type.
func DecodeRawBody[T any, R any](resp *Response[R]) (*T, error) {
	if resp.RawBody == nil {
		return nil, ErrRawBodyIsNil
	}

	var result T

	err := json.Unmarshal(resp.RawBody, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode raw body: %w", err)
	}

	return &result, nil
}
