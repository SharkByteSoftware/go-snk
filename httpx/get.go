// Package httpx provides simplified HTTP client functionality.
package httpx

import (
	"context"
	"net/http"
)

// Get sends an HTTP GET request to the specified URL with context, headers, and timeout and parses the response.
func Get[T any](ctx context.Context, url string, options ...Option) (*Response[T], error) {
	return DoRequest[T](ctx, http.MethodGet, url, nil, options...)
}

// GetRawResponse sends an HTTP GET request to the specified URL with context, headers, and timeout.
func GetRawResponse(ctx context.Context, url string, options ...Option) (*http.Response, error) {
	return DoRawRequest(ctx, http.MethodGet, url, nil, options...)
}
