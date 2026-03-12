package httpx

import "net/http"

// Response represents the structure for an HTTP response, supporting generics for handling various result types.
type Response[T any] struct {
	Status     string
	StatusCode int
	Header     http.Header
	Result     *T
}
