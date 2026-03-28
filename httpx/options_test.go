package httpx

import (
	"net/http"
	"net/url"
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
	assert.ErrorIs(t, err, ErrConfig)
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
	require.ErrorIs(t, err, ErrConfig)

	err = WithTimeout(-1)(config)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrConfig)
}

func TestAlwaysIncludeRawBody(t *testing.T) {
	config := NewHTTPXOptions()

	assert.False(t, config.includeRawBody)

	err := AlwaysIncludeRawBody()(config)
	require.NoError(t, err)
	assert.True(t, config.includeRawBody)
}

func Test_configWithAppliedOptions(t *testing.T) {
	options := []Option{
		WithTimeout(100),
		WithHeader("Content-Type", "application/json"),
		WithHeaders(http.Header{"Auth": []string{"secret"}}),
		WithParam("key", "value"),
		WithParams(url.Values{"key": []string{"value"}}),
		WithHTTPClient(http.DefaultClient),
	}

	config, err := configWithAppliedOptions(options)
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

	config, err = configWithAppliedOptions(options)
	require.Error(t, err)
	assert.Nil(t, config)
}
