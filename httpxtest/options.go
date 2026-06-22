package httpxtest

import (
	"net/http"
	"time"
)

// Option is a function that configures how the test server should respond.
type Option func(w http.ResponseWriter, r *http.Request)

// WithHeader sets a header on the response.
func WithHeader(key, value string) Option {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set(key, value)
	}
}

// WithDelay adds a delay to the response.
func WithDelay(delay time.Duration) Option {
	return func(_ http.ResponseWriter, _ *http.Request) {
		time.Sleep(delay)
	}
}
