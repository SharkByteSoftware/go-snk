//nolint:bodyclose
package httpx_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/SharkByteSoftware/go-snk/httpxtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	goodResponse        = `{"Name": "Test","Age": 18}`
	badResponse         = "bad response"
	errResponse         = `{ "Message": "custom error message", "Code": 400}`
	internalServerError = "internal server error: something went wrong"
	badURL              = "snk://localhost:1234"
)

var badParseURLFunc = func(_ string) (*url.URL, error) {
	url := &url.URL{
		Scheme: "://invalid",
		Host:   "example.com",
	}

	return url, nil
}

type testResponse struct {
	Name string
	Age  int
}

type testPayload struct {
	Name string
	Age  int
}

func TestGet(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Get[testResponse](ctx, badURL)
		assertTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Get[testResponse](nil, ts.URL)
		assertNilContext(t, err, resp)
	})

	t.Run("fail creating request", func(t *testing.T) {
		resp, err := httpx.Get[testResponse](ctx, badURL, httpx.WithParseURLFunc(badParseURLFunc))

		assertNewRequestError(t, err, resp)
	})
}

func TestGetRawResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.GetRawResponse(ctx, badURL)
		assertRawTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodGet, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.GetRawResponse(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestPost(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.Post[testResponse](ctx, "http://example.com", complex(1, 2))
		assertInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Post[testResponse](ctx, badURL, payload)
		assertTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Post[testResponse](nil, ts.URL, payload)
		assertNilContext(t, err, resp)
	})
}

func TestPostRawResponse(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		require.NoError(t, err)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.PostRawResponse(ctx, "http://example.com", complex(1, 2))
		assertRawInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.PostRawResponse(ctx, badURL, payload)
		assertRawTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPost, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.PostRawResponse(nil, ts.URL, payload)
		assertRawNilContext(t, err, resp)
	})
}

func TestPut(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.Put[testResponse](ctx, "http://example.com", complex(1, 2))
		assertInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Put[testResponse](ctx, badURL, payload)
		assertTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Put[testResponse](nil, ts.URL, payload)
		assertNilContext(t, err, resp)
	})
}
func TestPutRawResponse(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		require.NoError(t, err)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.PutRawResponse(ctx, "http://example.com", complex(1, 2))
		assertRawInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.PutRawResponse(ctx, badURL, payload)
		assertRawTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPut, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.PutRawResponse(nil, ts.URL, payload)
		assertRawNilContext(t, err, resp)
	})
}

func TestPatch(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.Patch[testResponse](ctx, "http://example.com", complex(1, 2))
		assertInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Patch[testResponse](ctx, badURL, payload)
		assertTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Patch[testResponse](nil, ts.URL, payload)
		assertNilContext(t, err, resp)
	})
}

func TestPatchRawResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, complex(1, 2))
		assertRawInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.PatchRawResponse(ctx, badURL, testPayload{Name: "Test", Age: 18})
		assertRawTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodPatch, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.PatchRawResponse(nil, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawNilContext(t, err, resp)
	})
}

func TestDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Delete[testResponse](ctx, badURL)
		assertTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Delete[testResponse](nil, ts.URL)
		assertNilContext(t, err, resp)
	})
}

func TestDeleteRawResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.DeleteRawResponse(ctx, badURL)
		assertRawTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodDelete, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.DeleteRawResponse(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestHead(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodHead, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodHead, "/", http.StatusNoContent, nil).
			Build()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodHead, "/", http.StatusOK, badResponse).
			Build()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodHead, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Head(ctx, badURL)
		assertRawTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodHead, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Head(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestOptions(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodOptions, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Options(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("options invalid response payload", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodOptions, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Options(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodOptions, "/", http.StatusInternalServerError, internalServerError).
			Build()

		resp, err := httpx.Options(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Options(ctx, badURL)
		assertRawTransportError(t, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			OnRoute(http.MethodOptions, "/", http.StatusOK, goodResponse).
			Build()

		resp, err := httpx.Options(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestWithOptions(t *testing.T) {
	ctx := context.Background()

	t.Run("fail with timeout config error", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, goodResponse, httpxtest.WithDelay(10*time.Millisecond)).
			Build()

		resp, err := httpx.Get[testResponse](ctx, ts.URL, httpx.WithTimeout(0))
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, httpx.ErrOptions)
		assert.ErrorContains(t, err, "apply options: invalid options: WithTimeout: invalid timeout, must be positive")
	})

	t.Run("fail with timeout", func(t *testing.T) {
		ts := httpxtest.NewServerBuilder(t).
			On(http.StatusOK, goodResponse, httpxtest.WithDelay(10*time.Millisecond)).
			Build()

		resp, err := httpx.Get[testResponse](ctx, ts.URL, httpx.WithTimeout(1*time.Millisecond))
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, httpx.ErrTransport)
		require.ErrorIs(t, err, context.DeadlineExceeded)
		assert.ErrorContains(t, err, "context deadline exceeded")
	})
}

func TestWithSkipVerify(t *testing.T) {
	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, goodResponse).
		BuildTLS()

	result, err := httpx.Get[testResponse](context.Background(), ts.URL, httpx.WithInsecureSkipVerify())
	assertStatusOkGoodResponse(t, err, result)
}

// TestGet_PreservesExistingURLQueryString demonstrates that a query string
// embedded directly in the request URL is dropped when no WithParam/WithParams
// option is supplied.
func TestGet_PreservesExistingURLQueryString(t *testing.T) {
	ctx := context.Background()

	var capturedRawQuery string

	ts := httpxtest.NewServerBuilder(t).
		OnRouteFunc(http.MethodGet, "/search", func(w http.ResponseWriter, r *http.Request) {
			capturedRawQuery = r.URL.RawQuery

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"Name":"Test","Age":18}`))
		}).
		Build()

	type result struct {
		Name string
		Age  int
	}

	_, err := httpx.Get[result](ctx, ts.URL+"/search?q=hello&page=2")
	require.NoError(t, err)

	// Expected: the server should see the query string the caller put on the URL.
	// Actual (bug): capturedRawQuery is "", because DoRequest rebuilds RawQuery
	// from the (empty) params option set instead of merging with what was there.
	assert.Equal(t, "q=hello&page=2", capturedRawQuery)
}

// TestGet_WithParamAlsoDropsExistingURLQueryString shows the same bug still
// bites even when the caller does use WithParam: the option's params replace
// the URL's original query string entirely rather than merging with it.
func TestGet_WithParamAlsoDropsExistingURLQueryString(t *testing.T) {
	ctx := context.Background()

	var capturedRawQuery string

	ts := httpxtest.NewServerBuilder(t).
		OnRouteFunc(http.MethodGet, "/search", func(w http.ResponseWriter, r *http.Request) {
			capturedRawQuery = r.URL.RawQuery

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"Name":"Test","Age":18}`))
		}).
		Build()

	type result struct {
		Name string
		Age  int
	}

	_, err := httpx.Get[result](ctx, ts.URL+"/search?q=hello", httpx.WithParam("page", "2"))
	require.NoError(t, err)

	// Expected (if merging correctly): both q=hello and page=2 present.
	// Actual (bug): only page=2 survives; q=hello from the URL is gone.
	assert.Equal(t, "page=2&q=hello", capturedRawQuery)
}

func assertStatusOkGoodResponse(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
}

func assertStatusOkNoContent(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Empty(t, resp.Result)
}

func assertRawInvalidPayload(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrEncoding)
	require.ErrorContains(t, err, "json: unsupported type")

	var encodingError *httpx.EncodingError
	require.ErrorAs(t, err, &encodingError)
	assert.ErrorContains(t, err, "encoding failed")
}

func assertStatusOkInvalidResponse(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrDecoding)

	var decodingError *httpx.DecodingError
	require.ErrorAs(t, err, &decodingError)
	require.ErrorContains(t, err, "decoding failed")
}

func assertNon2xxStatus(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrResponse)

	var respError *httpx.ResponseError
	require.ErrorAs(t, err, &respError)
	require.ErrorContains(t, err, "unexpected response")
}

func assertTransportError(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrTransport)
	require.ErrorContains(t, err, "transport failure")
}

func assertNilContext(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrTransport)
	require.ErrorContains(t, err, "invalid options: nil context")
}

func assertNewRequestError(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrTransport)
	require.ErrorContains(t, err, "transport failure")
}

func assertInvalidPayload(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrEncoding)
	require.ErrorContains(t, err, "json: unsupported type")

	var encodingError *httpx.EncodingError
	require.ErrorAs(t, err, &encodingError)
	assert.ErrorContains(t, err, "encoding failed")
}

func assertRawStatusOk(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func assertRawStatusOkNoContent(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func assertRawStatusOkInvalidResponse(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func assertRawNon2xxStatusCode(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func assertRawTransportError(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrTransport)
	assert.ErrorContains(t, err, "transport failure")
}

func assertRawNilContext(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrTransport)
	assert.ErrorContains(t, err, "invalid options: nil context")
}
