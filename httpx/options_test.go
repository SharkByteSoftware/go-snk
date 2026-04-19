package httpx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newHttpConfig(t *testing.T) {
	config := NewHTTPXOptions()
	require.NotNil(t, config)
}

func TestWithHttpClient(t *testing.T) {
	config := NewHTTPXOptions()

	err := WithHTTPClient(http.DefaultClient)(config)
	require.NoError(t, err)

	err = WithHTTPClient(nil)(config)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrOptions)
	require.ErrorContains(t, err, "invalid options: WithHTTPClient: http client is nil: <nil>")
}

func TestWithHeader(t *testing.T) {
	config := NewHTTPXOptions()

	err := WithHeader("Content-Type", "application/json")(config)
	require.NoError(t, err)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, config.headers)

	err = WithHeader("auth", "secret")(config)
	require.NoError(t, err)
	assert.Equal(t,
		http.Header{"Auth": []string{"secret"}, "Content-Type": []string{"application/json"}},
		config.headers)
}

func TestWithHeaders(t *testing.T) {
	config := NewHTTPXOptions()

	err := WithHeaders(http.Header{})(config)
	require.NoError(t, err)
	assert.Equal(t, http.Header{}, config.headers)

	err = WithHeaders(http.Header{"Content-Type": []string{"application/json"}})(config)
	require.NoError(t, err)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, config.headers)
}

func TestWithParam(t *testing.T) {
	config := NewHTTPXOptions()

	err := WithParam("key", "value")(config)

	require.NoError(t, err)
	assert.Equal(t, url.Values{"key": []string{"value"}}, config.params)
}

func TestWithParams(t *testing.T) {
	config := NewHTTPXOptions()

	err := WithParams(url.Values{
		"key":  []string{"value"},
		"key2": []string{"value2"},
	})(config)

	require.NoError(t, err)
	assert.Equal(t, url.Values{"key": []string{"value"}, "key2": []string{"value2"}}, config.params)
}

func TestWithTimeout(t *testing.T) {
	config := NewHTTPXOptions()

	err := WithTimeout(100)(config)
	require.NoError(t, err)
	assert.Equal(t, time.Duration(100), config.timeout)

	err = WithTimeout(0)(config)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrOptions)
	require.ErrorContains(t, err, "invalid options: WithTimeout: invalid timeout, must be positive: <nil>")

	err = WithTimeout(-1)(config)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrOptions)
	require.ErrorContains(t, err, "invalid options: WithTimeout: invalid timeout, must be positive: <nil>")
}

func Test_applyOptions(t *testing.T) {
	options := []Option{
		WithTimeout(100),
		WithHeader("Content-Type", "application/json"),
		WithHeaders(http.Header{"Auth": []string{"secret"}}),
		WithParam("key", "value"),
		WithParams(url.Values{"key": []string{"value"}}),
		WithHTTPClient(http.DefaultClient),
	}

	config, err := applyOptions(options)
	require.NoError(t, err)
	require.NotNil(t, config)

	assert.Equal(t, time.Duration(100), config.timeout)
	assert.Equal(t, http.Header{"Auth": []string{"secret"}, "Content-Type": []string{"application/json"}}, config.headers)
	assert.Equal(t, url.Values{"key": []string{"value"}}, config.params)

	options = []Option{
		WithTimeout(-1),
		WithHeader("Content-Type", "application/json"),
		WithHeaders(http.Header{"Auth": []string{"secret"}}),
		WithParam("key", "value"),
		WithParams(url.Values{"key": []string{"value"}}),
		WithHTTPClient(nil),
	}

	config, err = applyOptions(options)
	require.Error(t, err)
	assert.Nil(t, config)
}

func TestWithBearerToken(t *testing.T) {
	t.Run("sets Authorization header", func(t *testing.T) {
		var got string

		ts := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			got = r.Header.Get("Authorization")
		}))
		defer ts.Close()

		_, _ = Get[any](context.Background(), ts.URL, WithBearerToken("my-secret"))

		assert.Equal(t, "Bearer my-secret", got)
	})

	t.Run("empty token returns error", func(t *testing.T) {
		_, err := Get[any](context.Background(), "http://example.com", WithBearerToken(""))
		require.Error(t, err)
		require.ErrorIs(t, err, ErrOptions)
	})
}

func TestWithBasicAuth(t *testing.T) {
	t.Run("sets Authorization header", func(t *testing.T) {
		var got string

		ts := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			got = r.Header.Get("Authorization")
		}))
		defer ts.Close()

		_, _ = Get[any](context.Background(), ts.URL, WithBasicAuth("alice", "s3cret"))

		assert.True(t, strings.HasPrefix(got, "Basic "), "expected Basic auth prefix, got: %s", got)
	})

	t.Run("empty username returns error", func(t *testing.T) {
		_, err := Get[any](context.Background(), "http://example.com", WithBasicAuth("", "pass"))

		require.Error(t, err)
		require.ErrorIs(t, err, ErrOptions)
	})
}

func TestWithUserAgent(t *testing.T) {
	t.Run("sets User-Agent header", func(t *testing.T) {
		var got string

		ts := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			got = r.Header.Get("User-Agent")
		}))
		defer ts.Close()

		_, _ = Get[any](context.Background(), ts.URL, WithUserAgent("my-agent/1.0"))

		assert.Equal(t, "my-agent/1.0", got)
	})

	t.Run("empty user agent returns error", func(t *testing.T) {
		_, err := Get[any](context.Background(), "http://example.com", WithUserAgent(""))

		require.Error(t, err)
		require.ErrorIs(t, err, ErrOptions)
	})
}
