package httpx

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

var (
	// ErrOptions is returned when the options are invalid.
	ErrOptions = errors.New("invalid options")

	// ErrTransport is returned when the HTTP request fails to send or receive.
	ErrTransport = errors.New("transport failure")

	// ErrEncoding is returned when the request payload fails to encode.
	ErrEncoding = errors.New("encoding failed")

	// ErrDecoding is returned when the response body fails to decode.
	ErrDecoding = errors.New("decoding failed")

	// ErrResponse is returned when the server responds with a non-2xx HTTP status code.
	ErrResponse = errors.New("unexpected response")
)

// OptionsError is returned when one or more options fail to apply.
type OptionsError struct {
	Option  string
	Message string
	Err     error
}

// NewOptionsError returns a new OptionsError.
func NewOptionsError(option, msg string, err error) *OptionsError {
	return &OptionsError{
		Option:  option,
		Message: msg,
		Err:     err,
	}
}

func (e *OptionsError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("%s: %s: %s", ErrOptions, e.Option, e.Message)
	}

	return fmt.Sprintf("%s: %s: %s: %v", ErrOptions, e.Option, e.Message, e.Err)
}

func (e *OptionsError) Unwrap() error { return errors.Join(ErrOptions, e.Err) }

// NewTransportError returns a new TransportError.
func NewTransportError(err error) error {
	return fmt.Errorf("%w: %w", ErrTransport, err)
}

// EncodingError is returned when the request payload fails to encode.
type EncodingError struct {
	PayloadType string
	Err         error
}

// NewEncodingError returns a new EncodingError.
func NewEncodingError[T any](payload T, err error) *EncodingError {
	return &EncodingError{
		PayloadType: reflect.TypeOf(payload).String(),
		Err:         err,
	}
}

func (e *EncodingError) Error() string {
	return fmt.Sprintf("%s: %s: %v", ErrEncoding, e.PayloadType, e.Err)
}

func (e *EncodingError) Unwrap() error { return errors.Join(ErrEncoding, e.Err) }

// DecodingError is returned when the response body fails to decode.
type DecodingError struct {
	ContentType string
	Err         error
}

// NewDecodingError returns a new DecodingError.
func NewDecodingError(resp *http.Response, err error) *DecodingError {
	return &DecodingError{
		ContentType: resp.Header.Get("Content-Type"),
		Err:         err,
	}
}

func (e *DecodingError) Error() string {
	return fmt.Sprintf("%s: %s: %v", ErrDecoding, e.ContentType, e.Err)
}

func (e *DecodingError) Unwrap() error { return errors.Join(ErrDecoding, e.Err) }

// ResponseError is returned when the server responds with a non-2xx status code.
type ResponseError struct {
	Method     string
	URL        string
	StatusCode int
	Status     string
	Body       io.Closer
}

// NewResponseError returns a new ResponseError.
func NewResponseError(resp *http.Response) *ResponseError {
	return &ResponseError{
		Method:     resp.Request.Method,
		URL:        resp.Request.URL.String(),
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       resp.Body,
	}
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%s: %s %s: %d %s", ErrResponse, e.Method, e.URL, e.StatusCode, e.Status)
}

func (e *ResponseError) Unwrap() error { return ErrResponse }
