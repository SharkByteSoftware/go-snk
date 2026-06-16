package httpxtest_test

import (
	"context"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/SharkByteSoftware/go-snk/httpxtest"
	"github.com/stretchr/testify/require"
)

func TestServerBuilder_DefaultHandler(t *testing.T) {
	sb := httpxtest.NewServerBuilder()
	ts := sb.Build()

	require.NotNil(t, ts)
	require.NotNil(t, ts.URL)

	result, err := httpx.Get[string](context.Background(), ts.URL)
	require.Error(t, err)
	require.Nil(t, result)
}
