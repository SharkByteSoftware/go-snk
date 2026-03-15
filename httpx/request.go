package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DoRawRequest sends an HTTP request with the given method and body, without any shared request logic.
func DoRawRequest(
	ctx context.Context,
	method string,
	url string,
	body io.Reader,
	options ...Option,
) (*http.Response, error) {
	if ctx == nil {
		return nil, ErrContextIsNil
	}

	config, err := configWithAppliedOptions(options)
	if err != nil {
		return nil, fmt.Errorf("failed to apply options: %w", err)
	}

	req, err := newRequestWithAppliedConfig(ctx, method, url, body, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s request: %w", method, err)
	}

	client := clientWithAppliedConfig(config)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send %s request: %w", method, err)
	}

	return resp, nil
}

// DoRequest sends an HTTP request with the given method and body, applying shared request logic.
func DoRequest[T any](ctx context.Context, method string, url string, body io.Reader, options ...Option) (*Response[T], error) {
	resp, err := DoRawRequest(ctx, method, url, body, options...)
	if err != nil {
		return nil, err
	}

	//nolint: errcheck
	defer resp.Body.Close()

	return decodeResponse[T](resp)
}

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
