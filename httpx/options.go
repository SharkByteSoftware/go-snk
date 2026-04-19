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
	strictDecoding bool
	parseURLFunc   func(url string) (*url.URL, error)
}

// NewHTTPXOptions creates a new ConfigOptions instance.
func NewHTTPXOptions() *ConfigOptions {
	return &ConfigOptions{
		httpClient:     nil,
		headers:        make(http.Header),
		timeout:        defaultTimeout,
		params:         make(url.Values),
		strictDecoding: false,
		parseURLFunc:   url.Parse,
	}
}

// Option is a function that configures ConfigOptions.
type Option func(options *ConfigOptions) error

// WithHTTPClient sets the http client for the request.
// If the client is nil, an error is returned.
func WithHTTPClient(client *http.Client) Option {
	return func(options *ConfigOptions) error {
		if client == nil {
			return NewOptionsError("WithHTTPClient", "http client is nil", nil)
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
			return NewOptionsError("WithTimeout", "invalid timeout, must be positive", nil)
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

// StrictDecoding enables strict decoding of the response body.
func StrictDecoding() Option {
	return func(options *ConfigOptions) error {
		options.strictDecoding = true
		return nil
	}
}

// WithParseURLFunc sets the function to parse the URL.
func WithParseURLFunc(fn func(url string) (*url.URL, error)) Option {
	return func(options *ConfigOptions) error {
		options.parseURLFunc = fn
		return nil
	}
}

// WithBearerToken sets the Authorization header to "Bearer <token>".
// It is a named convenience wrapper around WithHeader to reduce the chance of
// misformatting the Authorization header.
func WithBearerToken(token string) Option {
	return func(options *ConfigOptions) error {
		if token == "" {
			return NewOptionsError("WithBearerToken", "token must not be empty", nil)
		}

		options.headers.Set("Authorization", "Bearer "+token)

		return nil
	}
}

// WithBasicAuth sets the Authorization header using HTTP Basic authentication.
// username and password are encoded per RFC 7617.
func WithBasicAuth(username string, password string) Option {
	return func(options *ConfigOptions) error {
		if username == "" {
			return NewOptionsError("WithBasicAuth", "username must not be empty", nil)
		}

		//nolint:exhaustruct
		req := &http.Request{Header: make(http.Header)}
		req.SetBasicAuth(username, password)
		options.headers.Set("Authorization", req.Header.Get("Authorization"))

		return nil
	}
}

// WithUserAgent sets the User-Agent request header.
func WithUserAgent(ua string) Option {
	return func(options *ConfigOptions) error {
		if ua == "" {
			return NewOptionsError("WithUserAgent", "user agent must not be empty", nil)
		}

		options.headers.Set("User-Agent", ua)

		return nil
	}
}

func applyOptions(options []Option) (*ConfigOptions, error) {
	cfg := NewHTTPXOptions()

	err := errors.Join(slicex.Map(options, func(option Option) error { return option(cfg) })...)
	if err != nil {
		return nil, fmt.Errorf("apply options: %w", err)
	}

	return cfg, nil
}
