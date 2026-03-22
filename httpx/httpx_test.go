package httpx_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SharkByteSoftware/go-snk/httpx"
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

type testResponse struct {
	Name string
	Age  int
}

type errorResponse struct {
	Message string
	Code    int
}

type testPayload struct {
	Name string
	Age  int
}

func TestGet(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.Get[testResponse](ctx, ts.URL)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Get[testResponse](ctx, badURL)
		assertTransportError(t, http.MethodGet, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Get[testResponse](nil, ts.URL)
		assertNilContext(t, err, resp)
	})
}

func TestGetRawResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.GetRawResponse(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.GetRawResponse(ctx, badURL)
		assertRawTransportError(t, http.MethodGet, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.GetRawResponse(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestPost(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.Post[testResponse](ctx, "http://example.com", complex(1, 2))
		assertInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.Post[testResponse](ctx, ts.URL, payload)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Post[testResponse](ctx, badURL, payload)
		assertTransportError(t, http.MethodPost, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Post[testResponse](nil, ts.URL, payload)
		assertNilContext(t, err, resp)
	})
}

func TestPostRawResponse(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		require.NoError(t, err)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.PostRawResponse(ctx, "http://example.com", complex(1, 2))
		assertRawInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.PostRawResponse(ctx, ts.URL, payload)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.PostRawResponse(ctx, badURL, payload)
		assertRawTransportError(t, http.MethodPost, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.PostRawResponse(nil, ts.URL, payload)
		assertRawNilContext(t, err, resp)
	})
}

func TestPut(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.Put[testResponse](ctx, "http://example.com", complex(1, 2))
		assertInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.Put[testResponse](ctx, ts.URL, payload)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Put[testResponse](ctx, badURL, payload)
		assertTransportError(t, http.MethodPut, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Put[testResponse](nil, ts.URL, payload)
		assertNilContext(t, err, resp)
	})
}
func TestPutRawResponse(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		require.NoError(t, err)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.PutRawResponse(ctx, "http://example.com", complex(1, 2))
		assertRawInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.PutRawResponse(ctx, ts.URL, payload)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.PutRawResponse(ctx, badURL, payload)
		assertRawTransportError(t, http.MethodPut, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.PutRawResponse(nil, ts.URL, payload)
		assertRawNilContext(t, err, resp)
	})
}

func TestPatch(t *testing.T) {
	ctx := context.Background()
	payload := testPayload{Name: "Test", Age: 18}

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		resp, err := httpx.Patch[testResponse](ctx, "http://example.com", complex(1, 2))
		assertInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.Patch[testResponse](ctx, ts.URL, payload)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Patch[testResponse](ctx, badURL, payload)
		assertTransportError(t, http.MethodPatch, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Patch[testResponse](nil, ts.URL, payload)
		assertNilContext(t, err, resp)
	})
}

func TestPatchRawResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, complex(1, 2))
		assertRawInvalidPayload(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.PatchRawResponse(ctx, badURL, testPayload{Name: "Test", Age: 18})
		assertRawTransportError(t, http.MethodPatch, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.PatchRawResponse(nil, ts.URL, testPayload{Name: "Test", Age: 18})
		assertRawNilContext(t, err, resp)
	})
}

func TestDelete(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertStatusOkGoodResponse(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.Delete[testResponse](ctx, ts.URL)
		assertNon2xxStatus(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Delete[testResponse](ctx, badURL)
		assertTransportError(t, http.MethodDelete, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Delete[testResponse](nil, ts.URL)
		assertNilContext(t, err, resp)
	})
}

func TestDeleteRawResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.DeleteRawResponse(ctx, badURL)
		assertRawTransportError(t, http.MethodDelete, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.DeleteRawResponse(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestHead(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("status code no content", func(t *testing.T) {
		ts := setupTestServer(http.StatusNoContent, "")
		defer ts.Close()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawStatusOkNoContent(t, err, resp)
	})

	t.Run("invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.Head(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Head(ctx, badURL)
		assertRawTransportError(t, http.MethodHead, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Head(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestOptions(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Options(ctx, ts.URL)
		assertRawStatusOk(t, err, resp)
	})

	t.Run("options invalid response payload", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, badResponse)
		defer ts.Close()

		resp, err := httpx.Options(ctx, ts.URL)
		assertRawStatusOkInvalidResponse(t, err, resp)
	})

	t.Run("non 2xx status code", func(t *testing.T) {
		ts := setupTestServer(http.StatusInternalServerError, internalServerError)
		defer ts.Close()

		resp, err := httpx.Options(ctx, ts.URL)
		assertRawNon2xxStatusCode(t, err, resp)
	})

	t.Run("transport error", func(t *testing.T) {
		resp, err := httpx.Options(ctx, badURL)
		assertRawTransportError(t, http.MethodOptions, err, resp)
	})

	t.Run("nil context", func(t *testing.T) {
		ts := setupTestServer(http.StatusOK, goodResponse)
		defer ts.Close()

		resp, err := httpx.Options(nil, ts.URL)
		assertRawNilContext(t, err, resp)
	})
}

func TestWithOptions(t *testing.T) {
	ctx := context.Background()

	t.Run("fail with timeout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(10 * time.Millisecond)
		}))
		defer ts.Close()

		resp, err := httpx.Get[testResponse](ctx, ts.URL, httpx.WithTimeout(1*time.Millisecond))
		require.Error(t, err)
		require.Nil(t, resp)
		assert.ErrorContains(t, err, "context deadline exceeded")
	})
}

func assertStatusOkGoodResponse(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
	assert.Empty(t, resp.RawBody)
}

func assertStatusOkNoContent(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Empty(t, resp.Result)
	assert.Empty(t, resp.RawBody)
}

func assertRawInvalidPayload(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")
}

func assertStatusOkInvalidResponse(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "failed to decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)
}

func assertNon2xxStatus(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "non-2xx status code")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(internalServerError), resp.RawBody)
}

func assertTransportError(t *testing.T, method string, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, fmt.Sprintf("failed to send %s request:", method))
}

func assertNilContext(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorIs(t, err, httpx.ErrContextIsNil)
}

func assertInvalidPayload(t *testing.T, err error, resp *httpx.Response[testResponse]) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")
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

func assertRawTransportError(t *testing.T, method string, err error, resp *http.Response) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, fmt.Sprintf("failed to send %s request:", method))
}

func assertRawNilContext(t *testing.T, err error, resp *http.Response) {
	t.Helper()

	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, httpx.ErrContextIsNil)
}

func setupTestServer(statusCode int, body string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(body))
	}))

	return ts
}
