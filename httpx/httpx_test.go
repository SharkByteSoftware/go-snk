package httpx

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHeader(t *testing.T) {
	config := newHttpConfig()

	WithHeader("Content-Type", "application/json")(config)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, config.headers)

	WithHeader("auth", "secret")(config)
	assert.Equal(t,
		http.Header{"Auth": []string{"secret"}, "Content-Type": []string{"application/json"}},
		config.headers)
}

func TestWithHeaders(t *testing.T) {
	config := newHttpConfig()

	WithHeaders(http.Header{})(config)
	assert.Equal(t, http.Header{}, config.headers)

	WithHeaders(http.Header{"Content-Type": []string{"application/json"}})(config)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, config.headers)
}

func TestWithTimeout(t *testing.T) {
}

func TestWithParam(t *testing.T) {
}

func TestWithParams(t *testing.T) {
}
