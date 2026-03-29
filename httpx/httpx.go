// Package httpx provides a simple HTTP client for Go.
package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Get sends an HTTP GET request to the specified URL and decodes the response body into T.
//
// Returns [ErrTransport] if the request fails to send or the response fails to read.
// Returns [ErrDecoding] as [*DecodingError] if the response body cannot be decoded into T.
// Returns [ErrResponse] as [*ResponseError] if the server responds with a non-2xx status code.
func Get[T any](ctx context.Context, url string, options ...Option) (*Response[T], error) {
	return DoRequest[T](ctx, http.MethodGet, url, nil, options...)
}

// GetRawResponse sends an HTTP GET request to the specified URL and returns the raw response.
//
// Returns [ErrTransport] if the request fails to send or the response fails to read.
func GetRawResponse(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodGet, url, nil, options...)
}

// Post sends an HTTP POST request to the specified URL with the given payload and decodes the response body into R.
//
// Returns [ErrEncoding] as [*EncodingError] if the payload cannot be marshaled to JSON.
// Returns [ErrTransport] if the request fails to send or the response fails to read.
// Returns [ErrDecoding] as [*DecodingError] if the response body cannot be decoded into R.
// Returns [ErrResponse] as [*ResponseError] if the server responds with a non-2xx status code.
func Post[R any, T any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, NewEncodingError(payload, err)
	}

	return DoRequest[R](ctx, http.MethodPost, url, bytes.NewReader(body), options...)
}

// PostRawResponse sends an HTTP POST request to the specified URL with the given payload and returns the raw response.
//
// Returns [ErrEncoding] as [*EncodingError] if the payload cannot be marshaled to JSON.
// Returns [ErrTransport] if the request fails to send or the response fails to read.
func PostRawResponse[T any](ctx context.Context, url string, payload T, options ...Option) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, NewEncodingError(payload, err)
	}

	return DoRawRequest(ctx, http.MethodPost, url, bytes.NewReader(body), options...)
}

// Put sends an HTTP PUT request to the specified URL with the given payload and decodes the response body into R.
//
// Returns [ErrEncoding] as [*EncodingError] if the payload cannot be marshaled to JSON.
// Returns [ErrTransport] if the request fails to send or the response fails to read.
// Returns [ErrDecoding] as [*DecodingError] if the response body cannot be decoded into R.
// Returns [ErrResponse] as [*ResponseError] if the server responds with a non-2xx status code.
func Put[R any, T any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, NewEncodingError(payload, err)
	}

	return DoRequest[R](ctx, http.MethodPut, url, bytes.NewReader(body), options...)
}

// PutRawResponse sends an HTTP PUT request to the specified URL with the given payload and returns the raw response.
//
// Returns [ErrEncoding] as [*EncodingError] if the payload cannot be marshaled to JSON.
// Returns [ErrTransport] if the request fails to send or the response fails to read.
func PutRawResponse[T any](ctx context.Context, url string, payload T, options ...Option) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, NewEncodingError(payload, err)
	}

	return DoRawRequest(ctx, http.MethodPut, url, bytes.NewReader(body), options...)
}

// Patch sends an HTTP PATCH request to the specified URL with the given payload and decodes the response body into R.
//
// Returns [ErrEncoding] as [*EncodingError] if the payload cannot be marshaled to JSON.
// Returns [ErrTransport] if the request fails to send or the response fails to read.
// Returns [ErrDecoding] as [*DecodingError] if the response body cannot be decoded into R.
// Returns [ErrResponse] as [*ResponseError] if the server responds with a non-2xx status code.
func Patch[R any, T any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, NewEncodingError(payload, err)
	}

	return DoRequest[R](ctx, http.MethodPatch, url, bytes.NewReader(body), options...)
}

// PatchRawResponse sends an HTTP PATCH request to the specified URL with the given payload and returns the raw response.
//
// Returns [ErrEncoding] as [*EncodingError] if the payload cannot be marshaled to JSON.
// Returns [ErrTransport] if the request fails to send or the response fails to read.
func PatchRawResponse[T any](ctx context.Context, url string, payload T, options ...Option) (*http.Response, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, NewEncodingError(payload, err)
	}

	return DoRawRequest(ctx, http.MethodPatch, url, bytes.NewReader(body), options...)
}

// Delete sends an HTTP DELETE request to the specified URL and decodes the response body into T.
//
// Returns [ErrTransport] if the request fails to send or the response fails to read.
// Returns [ErrDecoding] as [*DecodingError] if the response body cannot be decoded into T.
// Returns [ErrResponse] as [*ResponseError] if the server responds with a non-2xx status code.
func Delete[T any](ctx context.Context, url string, options ...Option) (*Response[T], error) {
	return DoRequest[T](ctx, http.MethodDelete, url, nil, options...)
}

// DeleteRawResponse sends an HTTP DELETE request to the specified URL and returns the raw response.
//
// Returns [ErrTransport] if the request fails to send or the response fails to read.
func DeleteRawResponse(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodDelete, url, nil, options...)
}

// Head sends an HTTP HEAD request to the specified URL and returns the raw response.
//
// Returns [ErrTransport] if the request fails to send or the response fails to read.
func Head(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodHead, url, nil, options...)
}

// Options sends an HTTP OPTIONS request to the specified URL and returns the raw response.
//
// Returns [ErrTransport] if the request fails to send or the response fails to read.
func Options(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodOptions, url, nil, options...)
}
