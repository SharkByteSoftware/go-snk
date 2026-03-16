package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response represents the structure for an HTTP response, supporting generics for handling various result types.
type Response[T any] struct {
	StatusCode int
	Status     string
	Header     http.Header
	Result     *T
	RawBody    []byte
	Request    *http.Request
}

// DecodeResponse decodes an HTTP response into a Response struct, handling various status codes and decoding the response body.
func DecodeResponse[T any](resp *http.Response, config *ConfigOptions) (*Response[T], error) {
	response := Response[T]{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Request:    resp.Request,
	}

	if resp.StatusCode == http.StatusNoContent {
		return &response, nil
	}

	if !is2xx(resp.StatusCode) {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return &response, fmt.Errorf("failed to read response body: %w", err)
		}

		response.RawBody = body

		return &response, ErrNon2xxStatusCode
	}

	var rawBody bytes.Buffer

	resp.Body = io.NopCloser(io.TeeReader(resp.Body, &rawBody))

	var result T

	decoder := json.NewDecoder(resp.Body)

	if config.strictDecoding {
		decoder.DisallowUnknownFields()
	}

	err := decoder.Decode(&result)
	if err != nil {
		response.RawBody = rawBody.Bytes()
		return &response, fmt.Errorf("failed to decode response body: %w", err)
	}

	response.Result = &result

	if config.includeRawBody {
		response.RawBody = rawBody.Bytes()
	}

	return &response, nil
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

func is2xx(code int) bool {
	return code >= 200 && code <= 299
}
