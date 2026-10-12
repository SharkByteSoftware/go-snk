package httpx

import (
	"crypto/tls"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errStubRoundTripper = errors.New("stub round tripper")

type stubRoundTripper struct{}

func (stubRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errStubRoundTripper
}

func TestClientWithAppliedConfig_InsecureTransportIsShared(t *testing.T) {
	cfgA, cfgB := NewOptions(), NewOptions()
	cfgA.insecureSkipVerify = true
	cfgB.insecureSkipVerify = true

	clientA := clientWithAppliedConfig(cfgA)
	clientB := clientWithAppliedConfig(cfgB)

	require.NotNil(t, clientA.Transport)
	assert.Same(t, clientA.Transport, clientB.Transport)

	secure := clientWithAppliedConfig(NewOptions())
	assert.Nil(t, secure.Transport)
}

func TestNewInsecureTransport_ClonesTransportBase(t *testing.T) {
	base := &http.Transport{MaxIdleConns: 7}

	got := newInsecureTransport(base)

	require.NotNil(t, got.TLSClientConfig)
	assert.True(t, got.TLSClientConfig.InsecureSkipVerify) //nolint:gosec
	assert.Equal(t, uint16(tls.VersionTLS12), got.TLSClientConfig.MinVersion)
	assert.Contains(t, got.TLSClientConfig.NextProtos, "h2") // HTTP/2 negotiation is kept
	assert.Equal(t, 7, got.MaxIdleConns)                     // settings carried over from base
	assert.NotSame(t, base, got)

	// The insecure settings must not leak into base. Clone initializes base's HTTP/2
	// defaults, so its TLS config may legitimately be non-nil.
	if base.TLSClientConfig != nil {
		assert.False(t, base.TLSClientConfig.InsecureSkipVerify) //nolint:gosec
		assert.NotSame(t, base.TLSClientConfig, got.TLSClientConfig)
	}
}

func TestNewInsecureTransport_FallbackForOtherRoundTrippers(t *testing.T) {
	got := newInsecureTransport(stubRoundTripper{})

	require.NotNil(t, got.TLSClientConfig)
	assert.True(t, got.TLSClientConfig.InsecureSkipVerify) //nolint:gosec
	assert.Equal(t, uint16(tls.VersionTLS12), got.TLSClientConfig.MinVersion)
	assert.NotNil(t, got.Proxy)
}
