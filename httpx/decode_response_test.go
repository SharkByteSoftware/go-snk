package httpx

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const goodResponse = `
{
	"Name": "Test",
	"Age": 18
}
`

const badResponse = "bad response"

type testResponse struct {
	Name string
	Age  int
}

type errReader struct{}

func (e errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func (e errReader) Close() error {
	return nil
}

func Test_decodeResponse(t *testing.T) {
	config := newHTTPConfig()

	response := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(goodResponse)),
	}

	resp, err := decodeResponse[testResponse](response, config)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
	assert.Empty(t, resp.RawBody)

	config.rawBodyOnError = true
	response.Body = io.NopCloser(strings.NewReader(goodResponse))

	resp, err = decodeResponse[testResponse](response, config)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
	assert.Empty(t, resp.RawBody)
}

func Test_decodeResponseDecodeFailure(t *testing.T) {
	config := newHTTPConfig()

	response := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(badResponse)),
	}

	resp, err := decodeResponse[testResponse](response, config)
	require.Error(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusOK), resp.Status)
	assert.Nil(t, resp.Result)
	assert.Empty(t, resp.RawBody)

	config.rawBodyOnError = true
	response.Body = io.NopCloser(strings.NewReader(badResponse))

	resp, err = decodeResponse[testResponse](response, config)
	require.Error(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusOK), resp.Status)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)
}

func Test_decodeResponse5500StatusCode(t *testing.T) {
	config := newHTTPConfig()

	response := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Status:     http.StatusText(http.StatusInternalServerError),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(goodResponse)),
	}

	resp, err := decodeResponse[testResponse](response, config)
	require.Error(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), resp.Status)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(goodResponse), resp.RawBody)

	response.Body = errReader{}
	resp, err = decodeResponse[testResponse](response, config)
	require.Error(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), resp.Status)
	assert.Nil(t, resp.Result)
	assert.Empty(t, resp.RawBody)
}
