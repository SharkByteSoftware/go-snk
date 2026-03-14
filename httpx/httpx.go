package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrClientCannotBeNil  = errors.New("http client cannot be nil") //nolint: revive
	ErrContextCannotBeNil = errors.New("context cannot be nil")
	ErrNon2xxStatusCode   = errors.New("non-2xx status code")
	ErrInvalidTimeout     = errors.New("invalid timeout, must be positive")
)

func clientWithAppliedConfig(config *httpxOptions) *http.Client {
	if config.httpClient != nil {
		return config.httpClient
	}

	return &http.Client{
		Timeout: config.timeout,
	}
}

func newRequestWithAppliedConfig(
	ctx context.Context,
	method string,
	baseURL string,
	body io.Reader,
	config *httpxOptions,
) (*http.Request, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	base.RawQuery = config.params.Encode()

	req, err := http.NewRequestWithContext(ctx, method, base.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header = config.headers

	return req, nil
}

func is2xx(code int) bool {
	return code >= 200 && code <= 299
}

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
