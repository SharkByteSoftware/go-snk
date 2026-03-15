package httpx_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	goodResponse        = `{"Name": "Test","Age": 18}`
	badResponse         = "bad response"
	internalServerError = "internal server error: something went wrong"
	badURL              = "snk://localhost:1234"
)

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

	// happy path
	ts := setupTestServer(http.StatusOK, goodResponse)
	resp, err := httpx.Get[testResponse](ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Equal(t, "Test", resp.Result.Name)
	assert.Equal(t, 18, resp.Result.Age)
	assert.Empty(t, resp.RawBody)

	// bad response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.Get[testResponse](ctx, ts.URL)
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "failed to decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.Get[testResponse](ctx, ts.URL)
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "non-2xx status code")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(internalServerError), resp.RawBody)

	// transport error
	resp, err = httpx.Get[testResponse](ctx, badURL)
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send GET request:")
}

func TestGetRawResponse(t *testing.T) {
	ctx := context.Background()

	// happy path
	ts := setupTestServer(http.StatusOK, goodResponse)
	resp, err := httpx.GetRawResponse(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// bad response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.GetRawResponse(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.GetRawResponse(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// transport error
	resp, err = httpx.GetRawResponse(ctx, badURL)
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send GET request:")
}

func TestPost(t *testing.T) {
	ctx := context.Background()

	// happy path
	ts := setupTestServer(http.StatusOK, goodResponse)
	resp, err := httpx.Post[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid payload
	resp, err = httpx.Post[testResponse](ctx, ts.URL, complex(1, 2))
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.Post[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "failed to decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.Post[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "non-2xx status code")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(internalServerError), resp.RawBody)

	// transport error
	resp, err = httpx.Post[testResponse](ctx, badURL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send POST request:")
}

func TestPostRawResponse(t *testing.T) {
	ctx := context.Background()
	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// happy path
	resp, err := httpx.PostRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid payload
	resp, err = httpx.PostRawResponse(ctx, ts.URL, complex(1, 2))
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.PostRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.PostRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// transport error
	resp, err = httpx.PostRawResponse(ctx, badURL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send POST request:")
}

func TestPut(t *testing.T) {
	ctx := context.Background()

	// happy path
	ts := setupTestServer(http.StatusOK, goodResponse)
	resp, err := httpx.Put[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid payload
	resp, err = httpx.Put[testResponse](ctx, ts.URL, complex(1, 2))
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.Put[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "failed to decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.Put[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "non-2xx status code")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(internalServerError), resp.RawBody)

	// transport error
	resp, err = httpx.Put[testResponse](ctx, badURL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send PUT request:")
}

func TestPutRawResponse(t *testing.T) {
	ctx := context.Background()
	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// happy path
	resp, err := httpx.PutRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid payload
	resp, err = httpx.PutRawResponse(ctx, ts.URL, complex(1, 2))
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.PutRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.PutRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// transport error
	resp, err = httpx.PutRawResponse(ctx, badURL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send PUT request:")
}

func TestPatch(t *testing.T) {
	ctx := context.Background()

	// happy path
	ts := setupTestServer(http.StatusOK, goodResponse)
	resp, err := httpx.Patch[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid payload
	resp, err = httpx.Patch[testResponse](ctx, ts.URL, complex(1, 2))
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.Patch[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "failed to decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.Patch[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "non-2xx status code")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(internalServerError), resp.RawBody)

	// transport error
	resp, err = httpx.Patch[testResponse](ctx, badURL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send PATCH request:")
}

func TestPatchRawResponse(t *testing.T) {
	ctx := context.Background()
	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// happy path
	resp, err := httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid payload
	resp, err = httpx.PatchRawResponse(ctx, ts.URL, complex(1, 2))
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "json: unsupported type")

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.PatchRawResponse(ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// transport error
	resp, err = httpx.PatchRawResponse(ctx, badURL, testPayload{Name: "Test", Age: 18})
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send PATCH request:")
}

func TestDelete(t *testing.T) {
	ctx := context.Background()

	// happy path
	ts := setupTestServer(http.StatusOK, goodResponse)
	resp, err := httpx.Delete[testResponse](ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.Delete[testResponse](ctx, ts.URL)
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "failed to decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(badResponse), resp.RawBody)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.Delete[testResponse](ctx, ts.URL)
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.ErrorContains(t, err, "non-2xx status code")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Nil(t, resp.Result)
	assert.Equal(t, []byte(internalServerError), resp.RawBody)

	// transport error
	resp, err = httpx.Delete[testResponse](ctx, badURL)
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send DELETE request:")
}

func TestDeleteRawResponse(t *testing.T) {
	ctx := context.Background()
	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// happy path
	resp, err := httpx.DeleteRawResponse(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.DeleteRawResponse(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.DeleteRawResponse(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// transport error
	resp, err = httpx.DeleteRawResponse(ctx, badURL)
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send DELETE request:")
}

func TestHead(t *testing.T) {
	ctx := context.Background()
	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// happy path
	resp, err := httpx.Head(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.Head(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.Head(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// transport error
	resp, err = httpx.Head(ctx, badURL)
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send HEAD request:")
}

func TestOptions(t *testing.T) {
	ctx := context.Background()
	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// happy path
	resp, err := httpx.Options(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// invalid response payload
	ts = setupTestServer(http.StatusOK, badResponse)
	resp, err = httpx.Options(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// non 2xx status code
	ts = setupTestServer(http.StatusInternalServerError, internalServerError)
	resp, err = httpx.Options(ctx, ts.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	// transport error
	resp, err = httpx.Options(ctx, badURL)
	require.Error(t, err)
	require.Nil(t, resp)
	assert.ErrorContains(t, err, "failed to send OPTIONS request:")
}

//func TestGet_EmptyContext(t *testing.T) {
//	resp, err := httpx.Get[testResponse](nil, "http://localhost")
//
//	require.Error(t, err)
//	require.Nil(t, resp)
//
//	assert.ErrorIs(t, err, httpx.ErrContextIsNil)
//}
//
//func TestGet_FailConfigWithAppliedOptions(t *testing.T) {
//	ctx := context.Background()
//
//	resp, err := httpx.Get[testResponse](ctx, "http://localhost", httpx.WithHTTPClient(nil))
//
//	require.Error(t, err)
//	require.Nil(t, resp)
//
//	assert.ErrorIs(t, err, httpx.ErrHTTPClientIsNil)
//}
//
//func TestGetEmptyResponse(t *testing.T) {
//	ctx := context.Background()
//
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte("{}"))
//	}))
//	defer ts.Close()
//
//	resp, err := httpx.Get[testResponse](ctx, ts.URL)
//	require.NoError(t, err)
//	require.NotNil(t, resp)
//	require.Equal(t, http.StatusOK, resp.StatusCode)
//
//	assert.Equal(t, "", resp.Result.Name)
//	assert.Equal(t, 0, resp.Result.Age)
//	assert.Empty(t, resp.RawBody)
//}
//
//func TestGet_InvalidURL(t *testing.T) {
//	ctx := context.Background()
//
//	resp, err := httpx.Get[testResponse](ctx, "http://invalid url")
//	require.Error(t, err)
//	require.Nil(t, resp)
//	assert.ErrorContains(t, err, "invalid url")
//
//	resp, err = httpx.Get[testResponse](ctx, "file://localhost")
//	require.Error(t, err)
//	require.Nil(t, resp)
//	assert.ErrorContains(t, err, "localhost")
//
//	resp, err = httpx.Get[testResponse](ctx, "")
//	require.Error(t, err)
//	require.Nil(t, resp)
//	assert.ErrorContains(t, err, "unsupported protocol scheme")
//}
//
//func TestGet_NoContent(t *testing.T) {
//	ctx := context.Background()
//
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusNoContent)
//	}))
//	defer ts.Close()
//
//	resp, err := httpx.Get[testResponse](ctx, ts.URL)
//	require.NoError(t, err)
//	require.NotNil(t, resp)
//
//	require.Equal(t, http.StatusNoContent, resp.StatusCode)
//	assert.Nil(t, resp.Result)
//	assert.Empty(t, resp.RawBody)
//}
//func TestGet_BadResponseBody(t *testing.T) {
//	ctx := context.Background()
//
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(badResponse))
//	}))
//	defer ts.Close()
//
//	response, err := httpx.Get[testResponse](ctx, ts.URL)
//	require.Error(t, err)
//	require.NotNil(t, response)
//	assert.Contains(t, err.Error(), "failed to decode response body")
//
//	assert.Equal(t, "200 OK", response.Status)
//	assert.Equal(t, http.StatusOK, response.StatusCode)
//	assert.Nil(t, response.Result)
//	assert.Equal(t, []byte(badResponse), response.RawBody)
//}
//
//func TestGet_BadRequest(t *testing.T) {
//	ctx := context.Background()
//
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusBadRequest)
//	}))
//	defer ts.Close()
//
//	response, err := httpx.Get[testResponse](ctx, ts.URL)
//	require.Error(t, err)
//	require.NotNil(t, response)
//
//	assert.Equal(t, "400 Bad Request", response.Status)
//	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
//	assert.NotEmpty(t, response.Header)
//	assert.Nil(t, response.Result)
//	assert.Empty(t, response.RawBody)
//}
//
//func TestGet_BadRequestRawBodyOnError(t *testing.T) {
//	ctx := context.Background()
//
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte(badResponse))
//	}))
//	defer ts.Close()
//
//	response, err := httpx.Get[testResponse](ctx, ts.URL)
//	require.Error(t, err)
//	require.NotNil(t, response)
//
//	assert.Equal(t, "400 Bad Request", response.Status)
//	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
//	assert.NotEmpty(t, response.Header)
//	assert.Nil(t, response.Result)
//	assert.Equal(t, []byte(badResponse), response.RawBody)
//}

func setupTestServer(statusCode int, body string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(body))
	}))

	return ts
}
