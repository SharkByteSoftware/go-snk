package httpx

import "net/http"

// Response represents the structure for an HTTP response, supporting generics for handling various result types.
type Response[T any] struct {
	StatusCode int
	Status     string
	Header     http.Header
	Result     *T
	RawBody    []byte
}
