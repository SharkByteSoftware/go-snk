package httpx_test

import (
	"context"
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

	ts = setupTestServer(http.StatusUnprocessableEntity, errResponse)
	defer ts.Close()

	result, err = httpx.Post[testResponse](ctx, ts.URL, testPayload{Name: "Test", Age: 18})
	errResult, err := httpx.DecodeRawBody[errorResponse](result)

	fmt.Println(errResult, err)

	// Output:
	// &{Test 18} 200 <nil>
	// &{Test 18} 200 <nil>
	// &{custom error message 400} <nil>
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

	fmt.Println(result.StatusCode, err)

	// Output:
	// 200 <nil>
	// 200 <nil>
}
