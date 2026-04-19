package httpx_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SharkByteSoftware/go-snk/httpx"
)

func ExampleGet() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// Example basic usage of httpx.Get
	result, err := httpx.Get[testResponse](ctx, ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Example with custom HTTP client, timeout, headers, and query parameters
	result, err = httpx.Get[testResponse](ctx, ts.URL,
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)

	fmt.Println(result.Result, result.StatusCode, err)

	// Output:
	// &{Test 18} 200 <nil>
	// &{Test 18} 200 <nil>
}

func ExamplePost() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// Example basic usage of httpx.Post
	result, err := httpx.Post[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Example with custom HTTP client, timeout, headers, and query parameters
	result, err = httpx.Post[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18},
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// setup test server to return 200 but a different payload than expected
	ts = setupTestServer(http.StatusOK, errResponse)
	defer ts.Close()

	_, err = httpx.Post[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18},
		httpx.StrictDecoding(),
	)

	// if the server returns a 200 status code, but the payload is not the expected type,
	// the error will be returned.
	var decodingError *httpx.DecodingError
	if errors.As(err, &decodingError) {
		fmt.Println(decodingError.Error())
	}

	// Output:
	// &{Test 18} 200 <nil>
	// &{Test 18} 200 <nil>
	// decoding failed: text/plain; charset=utf-8: decode: json: unknown field "Message"
}

func ExamplePut() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// Example basic usage of httpx.Put
	result, err := httpx.Put[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Example with custom HTTP client, timeout, headers, and query parameters
	result, err = httpx.Put[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18},
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Output:
	// &{Test 18} 200 <nil>
	// &{Test 18} 200 <nil>
}

func ExamplePatch() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// Example basic usage of httpx.Patch
	result, err := httpx.Patch[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Example with custom HTTP client, timeout, headers, and query parameters
	result, err = httpx.Patch[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18},
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Output:
	// &{Test 18} 200 <nil>
	// &{Test 18} 200 <nil>
}

func ExampleDelete() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// Example basic usage of httpx.Delete
	result, err := httpx.Delete[testResponse](ctx, ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Example with custom HTTP client, timeout, headers, and query parameters
	result, err = httpx.Delete[testResponse](ctx, ts.URL,
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Result, result.StatusCode, err)

	// Output:
	// &{Test 18} 200 <nil>
	// &{Test 18} 200 <nil>
}

func ExampleHead() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// Example basic usage of httpx.Head
	result, err := httpx.Head(ctx, ts.URL)
	if err != nil {
		panic(err)
	}

	defer func() { _ = result.Body.Close() }()

	fmt.Println(result.StatusCode, err)

	// Example with custom HTTP client, timeout, headers, and query parameters
	result, err = httpx.Head(ctx, ts.URL,
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)
	if err != nil {
		panic(err)
	}

	defer func() { _ = result.Body.Close() }()

	fmt.Println(result.StatusCode, err)

	// Output:
	// 200 <nil>
	// 200 <nil>
}

func ExampleOptions() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	// Example basic usage of httpx.Options
	result, err := httpx.Options(ctx, ts.URL)
	if err != nil {
		panic(err)
	}

	defer func() { _ = result.Body.Close() }()

	fmt.Println(result.StatusCode, err)

	// Example with custom HTTP client, timeout, headers, and query parameters
	result, err = httpx.Options(ctx, ts.URL,
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)
	if err != nil {
		panic(err)
	}

	defer func() { _ = result.Body.Close() }()

	fmt.Println(result.StatusCode, err)

	// Output:
	// 200 <nil>
	// 200 <nil>
}

func ExampleDoRawRequest() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	resp, err := httpx.DoRawRequest(ctx, http.MethodGet, ts.URL, nil)
	if err != nil {
		panic(err)
	}

	defer func() { _ = resp.Body.Close() }()

	fmt.Println(resp.StatusCode, err)
	// Output: 200 <nil>
}

func ExampleDoRequest() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	resp, err := httpx.DoRequest[testResponse](ctx, http.MethodGet, ts.URL, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Result, resp.StatusCode, err)
	// Output: &{Test 18} 200 <nil>
}

func ExampleDecodeResponse() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ts := setupTestServer(http.StatusOK, goodResponse)
	defer ts.Close()

	raw, err := httpx.GetRawResponse(ctx, ts.URL)
	if err != nil {
		panic(err)
	}

	defer func() { _ = raw.Body.Close() }()

	resp, err := httpx.DecodeResponse[testResponse](raw, httpx.NewOptions())
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Result, resp.StatusCode, err)
	// Output: &{Test 18} 200 <nil>
}
