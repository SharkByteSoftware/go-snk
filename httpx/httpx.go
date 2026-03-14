package httpx

import (
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = time.Second * 15

type httpConfig struct {
	headers http.Header
	timeout time.Duration
	params  url.Values
}

func newHttpConfig() *httpConfig {
	return &httpConfig{
		headers: make(http.Header),
		timeout: defaultTimeout,
		params:  make(url.Values),
	}
}

type Option func(options *httpConfig)

func WithHeader(key string, value string) Option {
	return func(options *httpConfig) {
		options.headers.Add(key, value)
	}
}

func WithHeaders(headers http.Header) Option {
	return func(options *httpConfig) {
		options.headers = headers
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(options *httpConfig) {
		options.timeout = timeout
	}
}

func WithParam(key string, value string) Option {
	return func(options *httpConfig) {
		options.params[key] = []string{value}
	}
}

func WithParams(params url.Values) Option {
	return func(options *httpConfig) {
		options.params = params
	}
}
