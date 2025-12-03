package httpx

import "net/http"

type Response[T any] struct {
	Status     string
	StatusCode int
	Header     http.Header
	Result     *T
}
