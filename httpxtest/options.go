package httpxtest

import (
	"net/http"
	"time"

	"github.com/SharkByteSoftware/go-snk/slicex"
)

// Option is a function that configures how the test server should respond.
type Option func(w http.ResponseWriter, r *http.Request)

// WithHeader sets a header on the response.
func WithHeader(key, value string) Option {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set(key, value)
	}
}

// WithHeaders sets multiple headers on the response.
func WithHeaders(headers http.Header) Option {
	return func(w http.ResponseWriter, _ *http.Request) {
		for key, values := range headers {
			slicex.Apply(values, func(value string) { w.Header().Add(key, value) })
		}
	}
}

// WithContentType sets the Content-Type header on the response.
func WithContentType(contentType string) Option {
	return WithHeader("Content-Type", contentType)
}

// WithJSONContentType sets the Content-Type header to "application/json".
func WithJSONContentType() Option {
	return WithHeader("Content-Type", "application/json")
}

// WithCookie sets a cookie on the response.
func WithCookie(name, value string) Option {
	return func(w http.ResponseWriter, _ *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: name, Value: value}) //nolint:exhaustruct
	}
}

// WithDelay adds a delay to the response.
func WithDelay(delay time.Duration) Option {
	return func(_ http.ResponseWriter, _ *http.Request) {
		time.Sleep(delay)
	}
}
