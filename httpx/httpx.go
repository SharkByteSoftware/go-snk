package httpx

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrHTTPClientCanNotBeNil = errors.New("http client cannot be nil") //nolint: revive
	ErrContextCannotBeNil    = errors.New("context cannot be nil")
	ErrNon2xxStatusCode      = errors.New("non-2xx status code")
	ErrInvalidTimeout        = errors.New("invalid timeout, must be positive")
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
