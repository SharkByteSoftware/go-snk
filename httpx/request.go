package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// DoRawRequest sends an HTTP request with the given method and body, without any shared request logic.
func DoRawRequest(ctx context.Context, method string, url string, body io.Reader, options ...Option) (*http.Response, error) {
	config, err := applyOptions(options)
	if err != nil {
		return nil, fmt.Errorf("DoRawRequest: %w", err)
	}

	return doRawRequest(ctx, method, url, body, config)
}

// DoRequest sends an HTTP request with the given method and body, applying shared request logic.
func DoRequest[T any](ctx context.Context, method string, url string, body io.Reader, options ...Option) (*Response[T], error) {
	config, err := applyOptions(options)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %w", err)
	}

	resp, err := doRawRequest(ctx, method, url, body, config)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	return DecodeResponse[T](resp, config)
}

func doRawRequest(ctx context.Context, method string, url string, body io.Reader, config *ConfigOptions) (*http.Response, error) {
	if ctx == nil {
		return nil, NewTransportError(fmt.Errorf("%w: nil context", ErrOptions))
	}

	req, err := newRequestWithAppliedConfig(ctx, method, url, body, config)
	if err != nil {
		return nil, NewTransportError(err)
	}

	client := clientWithAppliedConfig(config)

	resp, err := client.Do(req)
	if err != nil {
		return nil, NewTransportError(err)
	}

	return resp, nil
}

func clientWithAppliedConfig(config *ConfigOptions) *http.Client {
	if config.httpClient != nil {
		return config.httpClient
	}

	//nolint:exhaustruct
	return &http.Client{
		Timeout: config.timeout,
	}
}

func newRequestWithAppliedConfig(
	ctx context.Context,
	method string,
	baseURL string,
	body io.Reader,
	config *ConfigOptions,
) (*http.Request, error) {
	base, err := config.parseURLFunc(baseURL)
	if err != nil {
		return nil, err
	}

	base.RawQuery = config.params.Encode()

	req, err := http.NewRequestWithContext(ctx, method, base.String(), body)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req.Header = config.headers

	return req, nil
}
