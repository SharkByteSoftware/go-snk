// Package httpx provides simplified HTTP client functionality.
package httpx

import (
	"context"
	"fmt"
	"net/http"
)

// Get sends an HTTP GET request to the specified URL with context, headers, and timeout and parses the response.
func Get[T any](ctx context.Context, url string, options ...Option) (*Response[T], error) {
	if ctx == nil {
		return nil, ErrContextCannotBeNil
	}

	config, err := configWithAppliedOptions(options)
	if err != nil {
		return nil, fmt.Errorf("failed to apply options: %w", err)
	}

	req, err := newRequestWithAppliedConfig(ctx, http.MethodGet, url, nil, config)
	if err != nil {
		return nil, err
	}

	client := clientWithAppliedConfig(config)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	return decodeResponse[T](resp, config)
}
