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
	config := newHTTPConfig()
	require.NotNil(t, config)
}

func TestWithHttpClient(t *testing.T) {
	config := newHTTPConfig()

	err := WithHTTPClient(http.DefaultClient)(config)
	require.NoError(t, err)

	err = WithHTTPClient(nil)(config)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrHTTPClientIsNil)
}

func TestWithHeader(t *testing.T) {
	config := newHTTPConfig()

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
	config := newHTTPConfig()

	err := WithHeaders(http.Header{})(config)
	require.NoError(t, err)
	assert.Equal(t, http.Header{}, config.headers)

	err = WithHeaders(http.Header{"Content-Type": []string{"application/json"}})(config)
	require.NoError(t, err)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, config.headers)
}

func TestWithParam(t *testing.T) {
	config := newHTTPConfig()

	err := WithParam("key", "value")(config)

	require.NoError(t, err)
	assert.Equal(t, url.Values{"key": []string{"value"}}, config.params)
}

func TestWithParams(t *testing.T) {
	config := newHTTPConfig()

	err := WithParams(url.Values{
		"key":  []string{"value"},
		"key2": []string{"value2"},
	})(config)

	require.NoError(t, err)
	assert.Equal(t, url.Values{"key": []string{"value"}, "key2": []string{"value2"}}, config.params)
}

func TestWithTimeout(t *testing.T) {
	config := newHTTPConfig()

	err := WithTimeout(100)(config)
	require.NoError(t, err)
	assert.Equal(t, time.Duration(100), config.timeout)

	err = WithTimeout(0)(config)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidTimeout)

	err = WithTimeout(-1)(config)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidTimeout)
}

func Test_configWithAppliedOptions(t *testing.T) {
}
