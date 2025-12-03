package httpx_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/assert"
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

	resp, err := httpx.Get[MyStruct](context.Background(), ts.URL, http.Header{}, time.Second*1)
	assert.NoError(t, err)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp.Header)
	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)

	resp, err = httpx.Get[MyStruct](nil, ts.URL, http.Header{}, time.Second*1)
	assert.Error(t, err)
	assert.Nil(t, resp)

	resp, err = httpx.Get[MyStruct](context.Background(), "ts.URL", http.Header{}, 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
}

func TestGet_BadPayload(t *testing.T) {
	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(badResponse))
	}))
	defer ts.Close()

	response, err := httpx.Get[MyStruct](context.Background(), ts.URL, http.Header{}, time.Second*1)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGet_BadStatusCode(t *testing.T) {
	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(getResponse))
	}))
	defer ts.Close()

	response, err := httpx.Get[MyStruct](context.Background(), ts.URL, http.Header{}, time.Second*1)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "Name")
}
