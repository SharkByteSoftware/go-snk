package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Post sends an HTTP POST request to the specified URL with context, headers, and timeout and parses the response.
func Post[T any, R any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	if ctx == nil {
		return nil, ErrContextCannotBeNil
	}

	config, err := configWithAppliedOptions(options)
	if err != nil {
		return nil, fmt.Errorf("failed to apply options: %w", err)
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithAppliedConfig(ctx, http.MethodPost, url, bytes.NewBuffer(body), config)
	if err != nil {
		return nil, err
	}

	client := clientWithAppliedConfig(config)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	return decodeResponse[R](resp)
}
