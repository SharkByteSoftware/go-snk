package httpx_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const getResponse = `
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

func TestGet(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-go-snk", "go-snk")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getResponse))
	}))
	defer ts.Close()

	resp, err := httpx.Get[testResponse](ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Contains(t, resp.Header, "X-Go-Snk")
	assert.Equal(t, "go-snk", resp.Header.Get("x-go-snk"))
	assert.Contains(t, resp.Header, "Content-Type")
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
	assert.Empty(t, resp.RawBody)
}

func TestGet_EmptyContext(t *testing.T) {
	resp, err := httpx.Get[testResponse](nil, "http://localhost")

	require.Error(t, err)
	require.Nil(t, resp)

	assert.ErrorIs(t, err, httpx.ErrContextCannotBeNil)
}

func TestGet_FailConfigWithAppliedOptions(t *testing.T) {
	ctx := context.Background()

	resp, err := httpx.Get[testResponse](ctx, "http://localhost", httpx.WithHTTPClient(nil))

	require.Error(t, err)
	require.Nil(t, resp)

	assert.ErrorIs(t, err, httpx.ErrHTTPClientCanNotBeNil)
}

func TestGetEmptyResponse(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer ts.Close()

	resp, err := httpx.Get[testResponse](ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Equal(t, "", resp.Result.Name)
	assert.Equal(t, 0, resp.Result.Age)
	assert.Empty(t, resp.RawBody)
}

func TestGet_InvalidURL(t *testing.T) {
	ctx := context.Background()

	resp, err := httpx.Get[testResponse](ctx, "http://invalid url")
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "invalid url")

	resp, err = httpx.Get[testResponse](ctx, "file://localhost")
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "localhost")

	resp, err = httpx.Get[testResponse](ctx, "")
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "unsupported protocol scheme")
}

func TestGet_NoContent(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	resp, err := httpx.Get[testResponse](ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Empty(t, resp.RawBody)
}
func TestGet_BadResponseBody(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(badResponse))
	}))
	defer ts.Close()

	response, err := httpx.Get[testResponse](ctx, ts.URL)
	require.Error(t, err)
	require.NotNil(t, response)
	assert.Contains(t, err.Error(), "failed to decode response body")

	assert.Equal(t, "200 OK", response.Status)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Nil(t, response.Result)
	assert.Equal(t, []byte(badResponse), response.RawBody)
}

func TestGet_BadRequest(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	response, err := httpx.Get[testResponse](ctx, ts.URL)
	require.Error(t, err)
	require.NotNil(t, response)

	assert.Equal(t, "400 Bad Request", response.Status)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.NotEmpty(t, response.Header)
	assert.Nil(t, response.Result)
	assert.Empty(t, response.RawBody)
}

func TestGet_BadRequestRawBodyOnError(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badResponse))
	}))
	defer ts.Close()

	response, err := httpx.Get[testResponse](ctx, ts.URL)
	require.Error(t, err)
	require.NotNil(t, response)

	assert.Equal(t, "400 Bad Request", response.Status)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.NotEmpty(t, response.Header)
	assert.Nil(t, response.Result)
	assert.Equal(t, []byte(badResponse), response.RawBody)
}
