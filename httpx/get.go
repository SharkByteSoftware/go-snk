// Package httpx provides simplified HTTP client functionality.
package httpx

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SharkByteSoftware/go-snk/slicex"
)

// Get sends an HTTP GET request to the specified URL with context, headers, and timeout and parses the response.
func Get[T any](ctx context.Context, url string, options ...Option) (*Response[T], error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	config := httpConfig{}
	slicex.Apply(options, func(option Option) {
		option(&config)
	})

	req.Header = config.headers

	client := &http.Client{Timeout: config.timeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response[T]{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}

	if response.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(" %d: %s", resp.StatusCode, string(respBody))
	}

	var result T

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &Response[T]{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Result:     &result,
	}, nil
}
