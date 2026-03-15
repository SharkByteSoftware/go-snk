package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Post sends an HTTP POST request to the specified URL with context, headers, and timeout and parses the response.
func Post[T any, R any](ctx context.Context, url string, payload T, options ...Option) (*Response[R], error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return DoRequest[R](ctx, http.MethodPost, url, bytes.NewReader(body), options...)
}
