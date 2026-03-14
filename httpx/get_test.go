package httpx_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

type MyStruct struct {
	Name string
	Age  int
}

func TestGet(t *testing.T) {
	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getResponse))
	}))
	defer ts.Close()

	resp, err := httpx.Get[MyStruct](context.Background(), ts.URL, httpx.WithTimeout(time.Second*1))
	require.NoError(t, err)
	require.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp.Header)
	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)

	resp, err = httpx.Get[MyStruct](nil, ts.URL, httpx.WithTimeout(time.Second*1))
	require.Error(t, err)
	require.Nil(t, resp)

	resp, err = httpx.Get[MyStruct](context.Background(), "ts.URL", httpx.WithTimeout(0))
	require.Error(t, err)
	require.Nil(t, resp)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
}

func TestGet_BadPayload(t *testing.T) {
	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(badResponse))
	}))
	defer ts.Close()

	response, err := httpx.Get[MyStruct](context.Background(), ts.URL, httpx.WithTimeout(time.Second*1))
	require.Error(t, err)
	require.Nil(t, response)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGet_BadStatusCode(t *testing.T) {
	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(getResponse))
	}))
	defer ts.Close()

	response, err := httpx.Get[MyStruct](context.Background(), ts.URL, httpx.WithTimeout(time.Second*1))
	require.Error(t, err)
	require.Nil(t, response)
	assert.Contains(t, err.Error(), "Name")
}
