package httpxtest

import (
	"net/http"
	"time"
)

type Option func(w http.ResponseWriter, r *http.Request)


func WithHeader(key, value string) Option {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set(key, value)
	}
}

func WithDelay(delay time.Duration) Option {
	return func(_ http.ResponseWriter, _ *http.Request) {
		time.Sleep(delay)
	}
}
