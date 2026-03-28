package httpx_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/require"
)

func TestDoRawRequest(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		ctx := context.Background()

		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.DoRawRequest(ctx, http.MethodGet, ts.URL, nil)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("happy path with client", func(t *testing.T) {
		ctx := context.Background()

		ts := setupTestServer(http.StatusOK, "")
		defer ts.Close()

		resp, err := httpx.DoRawRequest(ctx, http.MethodGet, ts.URL, nil,
			httpx.WithHTTPClient(http.DefaultClient))
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad option", func(t *testing.T) {
		ctx := context.Background()

		ts := setupTestServer(http.StatusOK, "")
		defer ts.Close()

		resp, err := httpx.DoRawRequest(ctx, http.MethodGet, ts.URL, nil, httpx.WithHTTPClient(nil))
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("bad url", func(t *testing.T) {
		resp, err := httpx.DoRawRequest(context.Background(), http.MethodGet, "snk://google.com <something>", nil)
		require.Error(t, err)
		require.Nil(t, resp)
	})
}
