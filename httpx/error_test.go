package httpx_test

import (
	"net/http"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestPayload struct{}

func Test_OptionsError(t *testing.T) {
	err := httpx.NewOptionsError("MyOption", "my error message", httpx.ErrEncoding)
	require.Error(t, err)
	require.ErrorIs(t, err, httpx.ErrOptions)
	require.ErrorIs(t, err, httpx.ErrEncoding)

	assert.Equal(t, "MyOption", err.Option)
	assert.Equal(t, "my error message", err.Message)
	assert.Equal(t, httpx.ErrEncoding, err.Err)

	assert.Equal(t, "invalid options: MyOption: my error message: encoding failed", err.Error())
}

func Test_NewTransportError(t *testing.T) {
	err := httpx.NewTransportError(httpx.ErrOptions)
	require.Error(t, err)
	require.ErrorIs(t, err, httpx.ErrTransport)
	require.ErrorIs(t, err, httpx.ErrOptions)

	assert.Equal(t, "transport failure: invalid options", err.Error())
}

func Test_EncodingError(t *testing.T) {
	err := httpx.NewEncodingError(&TestPayload{}, httpx.ErrOptions)
	require.Error(t, err)
	require.ErrorIs(t, err, httpx.ErrEncoding)
	require.ErrorIs(t, err, httpx.ErrOptions)

	assert.Equal(t, "*httpx_test.TestPayload", err.PayloadType)
	assert.Equal(t, err.Err, httpx.ErrOptions)

	assert.Equal(t, "encoding failed: *httpx_test.TestPayload: invalid options", err.Error())
}

func Test_DecodingError(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       nil,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
	err := httpx.NewDecodingError(resp, httpx.ErrOptions)
	require.Error(t, err)
	require.ErrorIs(t, err, httpx.ErrDecoding)
	require.ErrorIs(t, err, httpx.ErrOptions)

	assert.Equal(t, "application/json", err.ContentType)
	assert.Equal(t, err.Err, httpx.ErrOptions)

	assert.Equal(t, "decoding failed: application/json : invalid options", err.Error())
}
