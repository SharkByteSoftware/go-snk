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

	result, err := httpx.Get[testResponse](
		ctx,
		ts.URL,
		httpx.WithHTTPClient(http.DefaultClient), // Optional: use a custom HTTP client
		httpx.WithTimeout(5*time.Second),
		httpx.WithHeader("X-Custom-Header", "Custom Value"),
		httpx.WithParam("key", "value"),
	)

	fmt.Println(result.Result, err)
	// Output: &{Test 18} <nil>
}
