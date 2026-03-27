// Package httpx provides a simple HTTP client for Go.
package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrNon2xxStatusCode is returned when the server responds with a non-2xx HTTP status code.
	ErrNon2xxStatusCode = errors.New("non-2xx status code")

	// ErrConfig is returned when the options are invalid.
	ErrConfig = errors.New("configuration error")

	// ErrTimeout is returned when the request times out.
	ErrTimeout = errors.New("timeout")

	// ErrTransport is returned when the transport fails.
	ErrTransport = errors.New("transport failure")

	// ErrDecoding is returned when the decoding fails.
	ErrDecoding = errors.New("decoding failed")

	// ErrMarshaling is returned when the marshaling fails.
	ErrMarshaling = errors.New("marshaling failed")
)

// Get sends an HTTP GET request to the specified URL with context, headers, and timeout and parses the response.
func Get[T any](ctx context.Context, url string, options ...Option) (*Response[T], error) {
	return DoRequest[T](ctx, http.MethodGet, url, nil, options...)
}

// GetRawResponse sends an HTTP GET request to the specified URL with context, headers, and timeout.
func GetRawResponse(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodGet, url, nil, options...)
}

// Post sends an HTTP POST request to the specified URL with context, headers, and timeout and parses the response.
func Post[R any, T any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshaling, err)
	}

	return DoRequest[R](ctx, http.MethodPost, url, bytes.NewReader(body), options...)
}

// PostRawResponse sends an HTTP POST request to the specified URL with context, headers, and timeout.
func PostRawResponse[T any](ctx context.Context, url string, payload T, options ...Option) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshaling, err)
	}

	return DoRawRequest(ctx, http.MethodPost, url, bytes.NewReader(body), options...)
}

// Put sends an HTTP PUT request to the specified URL with context, headers, and timeout and parses the response.
func Put[R any, T any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshaling, err)
	}

	return DoRequest[R](ctx, http.MethodPut, url, bytes.NewReader(body), options...)
}

// PutRawResponse sends an HTTP PUT request to the specified URL with context, headers, and timeout.
func PutRawResponse[T any](ctx context.Context, url string, payload T, options ...Option) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshaling, err)
	}

	return DoRawRequest(ctx, http.MethodPut, url, bytes.NewReader(body), options...)
}

// Patch sends an HTTP PATCH request to the specified URL with context, headers, and timeout and parses the response.
func Patch[R any, T any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshaling, err)
	}

	return DoRequest[R](ctx, http.MethodPatch, url, bytes.NewReader(body), options...)
}

// PatchRawResponse sends an HTTP PATCH request to the specified URL with context, headers, and timeout.
func PatchRawResponse[T any](ctx context.Context, url string, payload T, options ...Option) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMarshaling, err)
	}

	return DoRawRequest(ctx, http.MethodPatch, url, bytes.NewReader(body), options...)
}

// Delete sends an HTTP DELETE request to the specified URL with context, headers, and timeout and parses the response.
func Delete[T any](ctx context.Context, url string, options ...Option) (*Response[T], error) {
	return DoRequest[T](ctx, http.MethodDelete, url, nil, options...)
}

// DeleteRawResponse sends an HTTP DELETE request to the specified URL with context, headers, and timeout.
func DeleteRawResponse(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodDelete, url, nil, options...)
}

// Head sends an HTTP HEAD request to the specified URL with context, headers, and timeout.
func Head(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodHead, url, nil, options...)
}

// Options sends an HTTP OPTIONS request to the specified URL with context, headers, and timeout.
func Options(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodOptions, url, nil, options...)
}
