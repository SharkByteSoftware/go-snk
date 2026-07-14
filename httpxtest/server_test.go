package httpxtest_test

import (
	"context"
	"encoding/json"
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

func TestServer_NewServerBuilder(t *testing.T) {
	t.Run("NewServerBuilder (no TLS)", func(t *testing.T) {
		sb := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "no TLS"}).
			Build()

		require.NotNil(t, sb)
		require.NotEmpty(t, sb.URL)

		result, err := httpx.Get[myStruct](context.Background(), sb.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "no TLS", result.Result.Name)
	})

	t.Run("NewServerBuilder (TLS)", func(t *testing.T) {
		sb := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "with TLS"}).
			BuildTLS()

		require.NotNil(t, sb)
		require.NotEmpty(t, sb.URL)

		result, err := httpx.Get[myStruct](context.Background(), sb.URL, httpx.WithInsecureSkipVerify())
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "with TLS", result.Result.Name)
	})
}

func TestServer_DefaultHandler(t *testing.T) {
	sb := httpxtest.NewServerBuilder(t)
	ts := sb.Build()

	result, err := httpx.GetRawResponse(context.Background(), ts.URL)

	defer func() { _ = result.Body.Close() }()

	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
}

func TestServer_OnFunc(t *testing.T) {
	wasCalled := false

	ts := httpxtest.NewServerBuilder(t).
		OnFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(myStructReturn))
			wasCalled = true
		}).
		Build()

	result, err := httpx.Get[myStruct](context.Background(), ts.URL)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, wasCalled)
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, "test", result.Result.Name)
	require.Contains(t, result.Header, "Content-Type")
	assert.Contains(t, result.Header.Get("Content-Type"), "text/plain")
}

func TestServer_On(t *testing.T) {
	t.Run("On with nil return", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusCreated, nil).
			Build()

		result, err := httpx.Post[myStruct](context.Background(), ts.URL, myStruct{Name: "Horton"})
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Empty(t, result.Result)
		assert.Equal(t, http.StatusNoContent, result.StatusCode)
		require.NotContains(t, result.Header, "Content-Type")
	})

	t.Run("On with []byte return", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusCreated, []byte(myStructReturn)).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "test", result.Result.Name)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "text/plain")
	})

	t.Run("On with string return", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusCreated, myStructReturn).
			Build()

		result, err := httpx.Post[myStruct](context.Background(), ts.URL, myStruct{Name: "Horton"})
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "text/plain")
	})

	t.Run("On with json.RawMessage return", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusCreated, json.RawMessage(myStructReturn)).
			Build()

		result, err := httpx.Put[myStruct](context.Background(), ts.URL, myStruct{Name: "Horton"})
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusCreated, result.StatusCode)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "text/plain")
	})

	t.Run("On with struct return", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "defaultHorton"}).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "text/plain")
	})

	t.Run("On with encoding error", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, complex(1, 2)).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("On with options", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "defaultHorton"}, httpxtest.WithContentType("application/go-snk")).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "application/go-snk")
	})

	t.Run("On already defined", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = httpxtest.NewServerBuilder(t).
				On(http.StatusOK, myStruct{Name: "defaultHorton"}, httpxtest.WithContentType("application/go-snk")).
				On(http.StatusOK, myStruct{Name: "defaultHorton"}, httpxtest.WithContentType("application/go-snk")).
				Build()
		})
	})
}

func TestServer_OnRoute(t *testing.T) {
	t.Run("OnRoute GET", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"}).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "Horton", result.Result.Name)
	})

	t.Run("OnRoute POST (other method)", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/v1/horton", http.StatusOK, myStruct{Name: "hears a Who"}).
			Build()

		result, err := httpx.Post[myStruct](context.Background(), ts.URL+"/v1/horton", myStruct{Name: "Horton"})
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "hears a Who", result.Result.Name)
	})

	t.Run("OnRoute undefined route -> default handler", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"}).
			Build()

		result, err := httpx.GetRawResponse(context.Background(), ts.URL)

		defer func() { _ = result.Body.Close() }()

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})

	t.Run("OnRoute empty method", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t)

		assert.Panics(t, func() {
			ts.OnRoute("", "/v1/horton", http.StatusOK, myStruct{Name: "Horton"})
		})
	})

	t.Run("OnRoute duplicate path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"})

		assert.Panics(t, func() {
			ts.OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"})
		})
	})

	t.Run("OnRoute multiple routes", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"}).
			OnRoute(http.MethodGet, "/v1/other", http.StatusOK, myStruct{Name: "other"}).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "Horton", result.Result.Name)

		result, err = httpx.Get[myStruct](context.Background(), ts.URL+"/v1/other")
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "other", result.Result.Name)
	})
}

func TestServer_OnSequence(t *testing.T) {
	t.Run("On 1 sequence with ExhaustCycle", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnSequence(httpxtest.ExhaustCycle,
				httpxtest.Response(http.StatusOK, myStruct{Name: "defaultHorton"}),
			).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)

		result, err = httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)
	})

	t.Run("On 2 sequence with ExhaustRepeatLast", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnSequence(httpxtest.ExhaustRepeatLast,
				httpxtest.Response(http.StatusOK, myStruct{Name: "defaultHorton"}, httpxtest.WithJSONContentType()),
				httpxtest.Response(http.StatusOK, myStruct{Name: "nextHorton"}),
			).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)

		result, err = httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "nextHorton", result.Result.Name)
	})

	t.Run("On 1 sequence with ExhaustServerError", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnSequence(httpxtest.ExhaustServerError,
				httpxtest.Response(http.StatusOK, myStruct{Name: "defaultHorton"}),
			).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)

		result, err = httpx.Get[myStruct](context.Background(), ts.URL)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("On sequence already defined", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = httpxtest.NewServerBuilder(t).
				OnSequence(httpxtest.ExhaustServerError,
					httpxtest.Response(http.StatusOK, myStruct{Name: "defaultHorton"}),
				).
				OnSequence(httpxtest.ExhaustServerError,
					httpxtest.Response(http.StatusOK, myStruct{Name: "defaultHorton"}),
				).
				Build()
		})
	})
}

func TestServer_OnRouteSequence(t *testing.T) {
	t.Run("On route with ExhaustCycle", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRouteSequence(http.MethodGet, "/v1/horton", httpxtest.ExhaustCycle,
				httpxtest.Response(http.StatusOK, myStruct{Name: "defaultHorton"}),
			).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)

		result, err = httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, "defaultHorton", result.Result.Name)
	})
}

func TestServer_SrvLevelOptions(t *testing.T) {
	t.Run("With ContentType", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t, httpxtest.WithContentType("application/go-snk")).
			On(http.StatusOK, myStruct{Name: "defaultHorton"}).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "application/go-snk")
	})

	t.Run("With multiple options", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t, httpxtest.WithContentType("application/go-snk"),
			httpxtest.WithHeader("X-Srv-Lvl", "server-level")).
			On(http.StatusOK, myStruct{Name: "defaultHorton"}).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "application/go-snk")
		assert.Contains(t, result.Header.Get("X-Srv-Lvl"), "server-level")
	})
}

func TestServer_RouteLevelOptions(t *testing.T) {
	t.Run("With ContentType", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"},
				httpxtest.WithContentType("application/go-snk")).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "application/go-snk")
	})

	t.Run("With multiple options", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"},
				httpxtest.WithContentType("application/go-snk"),
				httpxtest.WithHeader("X-Rt-Lvl", "route-level"),
				httpxtest.WithHeader("X-Test", "test")).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "application/go-snk")
		assert.Contains(t, result.Header.Get("X-Rt-Lvl"), "route-level")
		assert.Contains(t, result.Header.Get("X-Test"), "test")
	})

	t.Run("Route level supersedes server level", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t, httpxtest.WithHeader("X-SVR-LVL", "server-level")).
			OnRoute(http.MethodGet, "/v1/other", http.StatusOK, myStruct{Name: "Horton"},
				httpxtest.WithHeader("X-Rt-Lvl", "route-level")).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/other")
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Contains(t, result.Header.Get("X-Rt-Lvl"), "route-level")
	})
}

func TestServer_Options(t *testing.T) {
	t.Run("With Header", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "defaultHorton"},
				httpxtest.WithHeader("X-Test", "test")).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Contains(t, result.Header.Get("X-Test"), "test")
	})

	t.Run("With Headers", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("X-Test", "test")
		headers.Add("X-Test2", "test2")

		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "defaultHorton"},
				httpxtest.WithHeaders(headers)).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Contains(t, result.Header.Get("X-Test"), "test")
		assert.Contains(t, result.Header.Get("X-Test2"), "test2")
	})

	t.Run("With ContentType", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "defaultHorton"},
				httpxtest.WithContentType("application/go-snk")).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Contains(t, result.Header, "Content-Type")
		assert.Contains(t, result.Header.Get("Content-Type"), "application/go-snk")
	})

	t.Run("With Cookie", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "defaultHorton"},
				httpxtest.WithCookie("test", "test")).
			Build()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Contains(t, result.Header.Get("Set-Cookie"), "test=test")
	})

	t.Run("With Delay", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, myStruct{Name: "defaultHorton"},
				httpxtest.WithDelay(500*time.Millisecond)).
			Build()

		now := time.Now()

		result, err := httpx.Get[myStruct](context.Background(), ts.URL)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.GreaterOrEqual(t, time.Since(now), 500*time.Millisecond)
	})
}

func TestServerBuilder_HowToUseIt(t *testing.T) {
	ts := httpxtest.NewServerBuilder(t, httpxtest.WithHeader("X-SVR-LVL", "server-level")).
		On(http.StatusOK, myStruct{Name: "defaultHorton"}).
		OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"},
			httpxtest.WithDelay(10*time.Millisecond),
			httpxtest.WithHeader("X-SVR-LVL", "route-level"),
			httpxtest.WithHeader("X-Test", "test")).
		OnRoute(http.MethodPost, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"}).
		OnRoute(http.MethodGet, "/v1/name", http.StatusOK, myStructReturn).
		OnRoute(http.MethodPost, "/v1/namne", http.StatusOK, myStruct{Name: "Horton"}).
		Build()

	require.NotNil(t, ts)
	require.NotEmpty(t, ts.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	result, err := httpx.Get[myStruct](ctx, ts.URL+"/v1/horton")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "Horton", result.Result.Name)
}
