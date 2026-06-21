package httpxtest_test

import (
	"context"
	"net/http"
	"testing"
	"time"

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
	sb := httpxtest.NewServerBuilder(t)
	ts := sb.Build()

	require.NotNil(t, ts)
	require.NotEmpty(t, ts.URL)

	result, err := httpx.Get[string](context.Background(), ts.URL)
	require.Error(t, err)
	require.Nil(t, result)
}

func TestServerBuilder_WithDefaultHandler(t *testing.T) {
	wasCalled := false

	ts := httpxtest.NewServerBuilder(t).
		OnFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(myStructReturn))
			wasCalled = true
		}).
		Build()

	require.NotNil(t, ts)
	require.NotEmpty(t, ts.URL)

	result, err := httpx.Get[myStruct](context.Background(), ts.URL)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.True(t, wasCalled)
	assert.Equal(t, "test", result.Result.Name)
}

func TestServerBuilder_HowToUseIt(t *testing.T) {
	ts := httpxtest.NewServerBuilder(t, httpxtest.WithHeader("X-SVR-LVL", "server-level")).
		On(http.StatusOK, myStruct{Name: "defaultHorton"}).
		OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"},
			httpxtest.WithDelay(5*time.Second),
			httpxtest.WithHeader("X-SVR-LVL", "route-level"),
			httpxtest.WithHeader("X-Test", "test")).
		OnRoute(http.MethodPost, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"}).
		OnRoute(http.MethodGet, "/v1/name", http.StatusOK, myStructReturn).
		OnRoute(http.MethodPost, "/v1/namne", http.StatusOK, myStruct{Name: "Horton"}).
		Build()

	require.NotNil(t, ts)
	require.NotEmpty(t, ts.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := httpx.Get[myStruct](ctx, ts.URL+"/v1/Horton")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "Horton", result.Result.Name)

	// req, _ := recorder.GetRequest("/v1/horton")
	// // assert.True(t, recorder.IsRouteCalled("/v1/horton"))
	// assert.Equal(t, http.MethodGet, req.Method)
	// assert.Equal(t, "/v1/horton", req.URL.Path)
}
