package httpx

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_clientWithAppliedConfig(t *testing.T) {
	config := newHTTPConfig()

	client := clientWithAppliedConfig(config)
	require.NotNil(t, client)
	assert.Equal(t, config.timeout, client.Timeout)

	config.timeout = time.Second * 5
	client = clientWithAppliedConfig(config)
	require.NotNil(t, client)
	assert.Equal(t, config.timeout, client.Timeout)
	assert.NotEqual(t, http.DefaultClient, client)

	config.httpClient = http.DefaultClient
	client = clientWithAppliedConfig(config)
	require.NotNil(t, client)
	assert.Equal(t, http.DefaultClient, client)
}

func Test_newRequestWithAppliedConfig(t *testing.T) {
	config := newHTTPConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := newRequestWithAppliedConfig(ctx, http.MethodGet, "https://google.com", nil, config)
	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Equal(t, ctx, req.Context())

	config.params = url.Values{"q": []string{"test"}}
	req, err = newRequestWithAppliedConfig(ctx, http.MethodGet, "https://google.com", nil, config)
	require.NoError(t, err)
	require.NotNil(t, req)
	assert.Equal(t, "q=test", req.URL.RawQuery)
	assert.Equal(t, "https://google.com?q=test", req.URL.String())

	req, err = newRequestWithAppliedConfig(nil, http.MethodGet, "https://google.com", nil, config)
	require.Error(t, err)
	require.Nil(t, req)
}

func Test_is2xx(t *testing.T) {
	assert.True(t, is2xx(http.StatusOK))
	assert.True(t, is2xx(http.StatusCreated))
	assert.True(t, is2xx(http.StatusAccepted))
	assert.False(t, is2xx(http.StatusBadRequest))
	assert.False(t, is2xx(http.StatusInternalServerError))

}
