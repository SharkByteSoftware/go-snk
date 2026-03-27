package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/SharkByteSoftware/go-snk/mapx"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

const defaultTimeout = time.Second * 15

// ConfigOptions contains the configuration options for HTTPX.
type ConfigOptions struct {
	httpClient     *http.Client
	headers        http.Header
	timeout        time.Duration
	params         url.Values
	includeRawBody bool
	strictDecoding bool
}

// NewHTTPXOptions creates a new ConfigOptions instance.
func NewHTTPXOptions() *ConfigOptions {
	return &ConfigOptions{
		headers:        make(http.Header),
		timeout:        defaultTimeout,
		params:         make(url.Values),
		includeRawBody: false,
		strictDecoding: false,
	}
}

// Option is a function that configures ConfigOptions.
type Option func(options *ConfigOptions) error

// WithHTTPClient sets the http client for the request.
// If the client is nil, an error is returned.
func WithHTTPClient(client *http.Client) Option {
	return func(options *ConfigOptions) error {
		if client == nil {
			return fmt.Errorf("%w: http client is nil", ErrOptions)
		}

		options.httpClient = client

		return nil
	}
}

// WithHeader adds a single header to the request.
func WithHeader(key string, value string) Option {
	return func(options *ConfigOptions) error {
		options.headers.Add(key, value)
		return nil
	}
}

// WithHeaders combines the provided headers with the existing headers.
func WithHeaders(headers http.Header) Option {
	return func(options *ConfigOptions) error {
		options.headers = mapx.Combine(options.headers, headers)
		return nil
	}
}

// WithTimeout sets the timeout for the request.
// If the timeout is <= 0, an error is returned.
func WithTimeout(timeout time.Duration) Option {
	return func(options *ConfigOptions) error {
		if timeout <= 0 {
			return fmt.Errorf("%w: invalid timeout, must be positive", ErrOptions)
		}

		options.timeout = timeout

		return nil
	}
}

// WithParam adds a single param to the request.
func WithParam(key string, value string) Option {
	return func(options *ConfigOptions) error {
		options.params[key] = []string{value}
		return nil
	}
}

// WithParams combines the provided params with the existing params.
func WithParams(params url.Values) Option {
	return func(options *ConfigOptions) error {
		options.params = mapx.Combine(options.params, params)
		return nil
	}
}

// AlwaysIncludeRawBody enables the inclusion of the raw request body in the ConfigOptions configuration.
func AlwaysIncludeRawBody() Option {
	return func(options *ConfigOptions) error {
		options.includeRawBody = true
		return nil
	}
}

// StrictDecoding enables strict decoding of the response body.
func StrictDecoding() Option {
	return func(options *ConfigOptions) error {
		options.strictDecoding = true
		return nil
	}
}

func configWithAppliedOptions(options []Option) (*ConfigOptions, error) {
	config := NewHTTPXOptions()

	err := errors.Join(slicex.Map(options, func(option Option) error { return option(config) })...)
	if err != nil {
		return nil, fmt.Errorf("failed applying option: %w", err)
	}

	return config, nil
}
