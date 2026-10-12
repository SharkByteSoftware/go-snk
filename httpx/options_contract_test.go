package httpx_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptions_ParamContract(t *testing.T) {
	queries := make(chan url.Values, 1)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries <- r.URL.Query()

		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	resp, err := httpx.DoRawRequest(context.Background(), http.MethodGet, srv.URL+"?a=fromurl&b=keep", nil,
		httpx.WithParam("a", "1"),
		httpx.WithParam("a", "2"),
	)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	got := <-queries
	assert.Equal(t, []string{"2"}, got["a"]) // WithParam replaces, and overrides the URL's value
	assert.Equal(t, []string{"keep"}, got["b"])
}

func TestOptions_HeaderContract(t *testing.T) {
	headers := make(chan http.Header, 1)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers <- r.Header.Clone()

		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	resp, err := httpx.DoRawRequest(context.Background(), http.MethodGet, srv.URL, nil,
		httpx.WithHeader("X-Multi", "a"),
		httpx.WithHeader("X-Multi", "b"),
		httpx.WithHeaders(http.Header{"X-Replaced": {"new"}}),
	)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	got := <-headers
	assert.Equal(t, []string{"a", "b"}, got.Values("X-Multi")) // WithHeader appends
	assert.Equal(t, []string{"new"}, got.Values("X-Replaced"))
}
