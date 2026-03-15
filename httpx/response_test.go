package httpx_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type errReader struct{}

func (e errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func (e errReader) Close() error {
	return nil
}

func Test_decodeResponse(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(goodResponse)),
	}

	resp, err := httpx.DecodeResponse[testResponse](response)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
	assert.Empty(t, resp.RawBody)
}

func Test_decodeResponseDecodeFailure(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(badResponse)),
	}

	resp, err := httpx.DecodeResponse[testResponse](response)
	require.Error(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusOK), resp.Status)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)
}

func Test_decodeResponse5500StatusCode(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Status:     http.StatusText(http.StatusInternalServerError),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(goodResponse)),
	}

	resp, err := httpx.DecodeResponse[testResponse](response)
	require.Error(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), resp.Status)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(goodResponse), resp.RawBody)

	response.Body = errReader{}
	resp, err = httpx.DecodeResponse[testResponse](response)
	require.Error(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), resp.Status)
	assert.Nil(t, resp.Result)
	assert.Empty(t, resp.RawBody)
}

func TestDecodeRawBody(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		resp := &httpx.Response[testResponse]{
			StatusCode: http.StatusOK,
			RawBody:    []byte(goodResponse),
		}

		result, err := httpx.DecodeRawBody[testResponse](resp)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Test", result.Name)
		assert.Equal(t, 18, result.Age)
	})

	t.Run("nil raw body", func(t *testing.T) {
		resp := &httpx.Response[testResponse]{
			StatusCode: http.StatusOK,
			RawBody:    nil,
		}

		result, err := httpx.DecodeRawBody[testResponse](resp)
		require.Error(t, err)
		require.Nil(t, result)
		assert.ErrorIs(t, err, httpx.ErrRawBodyIsNil)
	})

	t.Run("decode failure", func(t *testing.T) {
		resp := &httpx.Response[testResponse]{
			StatusCode: http.StatusOK,
			RawBody:    []byte(badResponse),
		}

		result, err := httpx.DecodeRawBody[testResponse](resp)
		require.Error(t, err)
		require.Nil(t, result)
		assert.ErrorContains(t, err, "failed to decode raw body:")
	})

}
