package httpxtest_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/SharkByteSoftware/go-snk/httpxtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type myStruct struct {
	Name string `json:"name"`
}

const myStructReturn = `{"name":"test"}`

func TestServerBuilder_DefaultHandler(t *testing.T) {
	sb := httpxtest.NewServerBuilder()
	ts := sb.Build(t)

	require.NotNil(t, ts)
	require.NotEmpty(t, ts.URL)

	result, err := httpx.Get[string](context.Background(), ts.URL)
	require.Error(t, err)
	require.Nil(t, result)
}

func TestServerBuilder_WithDefaultHandler(t *testing.T) {
	wasCalled := false

	ts := httpxtest.NewServerBuilder().
		WithDefaultHandler(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(myStructReturn))
			wasCalled = true
		}).
		Build(t)

	require.NotNil(t, ts)
	require.NotEmpty(t, ts.URL)

	result, err := httpx.Get[myStruct](context.Background(), ts.URL)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.True(t, wasCalled)
	assert.Equal(t, "test", result.Result.Name)
}

func TestServerBuilder_HowToUseIt(t *testing.T) {
	ts := httpxtest.NewServerBuilder().
		WithDefaultHandler(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(myStructReturn))
		}).
		OnGet("/v1/horton", http.StatusOK, myStruct{Name: "Horton"}).
		OnGetStr("/v1/name", http.StatusOK, myStructReturn).
		Build(t)

	require.NotNil(t, ts)
	require.NotEmpty(t, ts.URL)
}
