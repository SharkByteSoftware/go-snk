package httpx

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/SharkByteSoftware/go-snk/mapx"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

const defaultTimeout = time.Second * 15

type httpxOptions struct {
	httpClient *http.Client
	headers    http.Header
	timeout    time.Duration
	params     url.Values
}

func newHTTPConfig() *httpxOptions {
	return &httpxOptions{
		headers: make(http.Header),
		timeout: defaultTimeout,
		params:  make(url.Values),
	}
}

// Option is a function that configures the httpxOptions.
type Option func(options *httpxOptions) error

// WithHTTPClient sets the http client for the request.
func WithHTTPClient(client *http.Client) Option {
	return func(options *httpxOptions) error {
		if client == nil {
			return ErrHTTPClientIsNil
		}

		options.httpClient = client

		return nil
	}
}

// WithHeader adds a single header to the request.
func WithHeader(key string, value string) Option {
	return func(options *httpxOptions) error {
		options.headers.Add(key, value)
		return nil
	}
}

// WithHeaders combines the provided headers with the existing headers.
func WithHeaders(headers http.Header) Option {
	return func(options *httpxOptions) error {
		options.headers = mapx.Combine(options.headers, headers)
		return nil
	}
}

// WithTimeout sets the timeout for the request.
func WithTimeout(timeout time.Duration) Option {
	return func(options *httpxOptions) error {
		if timeout <= 0 {
			return ErrInvalidTimeout
		}

		options.timeout = timeout

		return nil
	}
}

// WithParam adds a single param to the request.
func WithParam(key string, value string) Option {
	return func(options *httpxOptions) error {
		options.params[key] = []string{value}
		return nil
	}
}

// WithParams combines the provided params with the existing params.
func WithParams(params url.Values) Option {
	return func(options *httpxOptions) error {
		options.params = mapx.Combine(options.params, params)
		return nil
	}
}

func configWithAppliedOptions(options []Option) (*httpxOptions, error) {
	config := newHTTPConfig()

	errs := slicex.Map(options, func(option Option) error {
		return option(config)
	})

	return config, errors.Join(errs...)
}
