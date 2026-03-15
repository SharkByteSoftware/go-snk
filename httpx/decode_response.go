package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func decodeResponse[T any](resp *http.Response, config *httpxOptions) (*Response[T], error) {
	response := Response[T]{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
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
	if config.rawBodyOnError {
		resp.Body = io.NopCloser(io.TeeReader(resp.Body, &rawBody))
	}

	var result T

	err := json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		response.RawBody = rawBody.Bytes()
		return &response, fmt.Errorf("failed to decode response body: %w", err)
	}

	response.Result = &result

	return &response, nil
}
