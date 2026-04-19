package httpx_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeResponse(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(goodResponse)),
	}

	resp, err := httpx.DecodeResponse[testResponse](response, httpx.NewOptions())
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
}

func TestDecodeResponse_DecodeFailure(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(badResponse)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{Scheme: "http", Host: "example.com", Path: "/api"},
		},
	}

	resp, err := httpx.DecodeResponse[testResponse](response, httpx.NewOptions())
	require.Error(t, err)
	require.ErrorIs(t, err, httpx.ErrDecoding)
	require.Nil(t, resp)
	require.ErrorContains(t, err, "decoding failed")

	var decodingError *httpx.DecodingError
	require.ErrorAs(t, err, &decodingError)
	assert.Equal(t, "application/json", decodingError.ContentType)
	assert.ErrorContains(t, decodingError.Err, "invalid character 'b' looking for beginning of value")
}

func TestDecodeResponse_500StatusCode(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Status:     http.StatusText(http.StatusInternalServerError),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(goodResponse)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{Scheme: "http", Host: "example.com", Path: "/api"},
		},
	}

	resp, err := httpx.DecodeResponse[testResponse](response, httpx.NewOptions())
	require.Error(t, err)
	require.ErrorIs(t, err, httpx.ErrResponse)
	require.Nil(t, resp)
	require.ErrorContains(t, err, "unexpected response")

	var respError *httpx.ResponseError
	require.ErrorAs(t, err, &respError)
	assert.Equal(t, http.StatusInternalServerError, respError.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), respError.Status)
}
