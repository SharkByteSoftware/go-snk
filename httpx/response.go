package httpx

import (
	"net/http"

	"github.com/SharkByteSoftware/go-snk/jsonx"
)

// Response represents the structure for an HTTP response, supporting generics for handling various result types.
type Response[T any] struct {
	StatusCode int
	Status     string
	Header     http.Header
	Result     *T
}

// DecodeResponse decodes an HTTP response into a Response struct, handling various status codes and decoding the response body.
func DecodeResponse[T any](resp *http.Response, config *ConfigOptions) (*Response[T], error) {
	response := Response[T]{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Result:     nil,
	}

	if resp.StatusCode == http.StatusNoContent {
		return &response, nil
	}

	if !is2xx(resp.StatusCode) {
		return nil, NewResponseError(resp)
	}

	var opts []jsonx.DecodeOption
	if config.strictDecoding {
		opts = append(opts, jsonx.WithStrictDecoding())
	}

	result, err := jsonx.Decode[T](resp.Body, opts...)
	if err != nil {
		return nil, NewDecodingError(resp, err)
	}

	response.Result = result

	return &response, nil
}

func is2xx(code int) bool {
	return code >= 200 && code <= 299
}
