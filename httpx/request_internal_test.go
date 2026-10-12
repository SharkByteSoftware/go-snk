package httpx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
