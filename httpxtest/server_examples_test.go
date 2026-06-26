package httpxtest_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/SharkByteSoftware/go-snk/httpx"
	"github.com/SharkByteSoftware/go-snk/httpxtest"
)

// The examples below construct a throwaway *testing.T so they can run as
// documentation. In your own tests, pass the *testing.T from the test
// function instead; the server is then closed automatically via t.Cleanup.

func ExampleNewServerBuilder() {
	t := &testing.T{}

	// On sets the handler for any request that does not match a route.
	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "Horton"}).
		Build()

	resp, err := httpx.Get[myStruct](context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode, resp.Result.Name)
	// Output: 200 Horton
}

func ExampleNewServerBuilder_serverLevelOptions() {
	t := &testing.T{}

	// Options passed to NewServerBuilder apply to every response.
	ts := httpxtest.NewServerBuilder(t,
		httpxtest.WithHeader("X-Server", "go-snk"),
		httpxtest.WithJSONContentType()).
		OnRoute(http.MethodGet, "/a", http.StatusOK, myStruct{Name: "a"}).
		OnRoute(http.MethodGet, "/b", http.StatusOK, myStruct{Name: "b"}).
		Build()

	respA, err := httpx.GetRawResponse(context.Background(), ts.URL+"/a")
	if err != nil {
		panic(err)
	}

	defer func() { _ = respA.Body.Close() }()

	respB, err := httpx.GetRawResponse(context.Background(), ts.URL+"/b")
	if err != nil {
		panic(err)
	}

	defer func() { _ = respB.Body.Close() }()

	fmt.Println(respA.Header.Get("X-Server"))
	fmt.Println(respA.Header.Get("Content-Type"))
	// Output: go-snk
	// application/json
}

func ExampleServerBuilder_On() {
	t := &testing.T{}

	// A nil response writes 204 No Content with an empty body.
	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusCreated, nil).
		Build()

	resp, err := httpx.GetRawResponse(context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	defer func() { _ = resp.Body.Close() }()

	fmt.Println(resp.StatusCode)
	// Output: 204
}

func ExampleServerBuilder_OnFunc() {
	t := &testing.T{}

	// OnFunc sets a custom handler for any unmatched request.
	ts := httpxtest.NewServerBuilder(t).
		OnFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(myStructReturn))
		}).
		Build()

	resp, err := httpx.Get[myStruct](context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode, resp.Result.Name)
	// Output: 200 test
}

func ExampleServerBuilder_OnRoute() {
	t := &testing.T{}

	// OnRoute responds to a specific method/route; anything else falls
	// through to the default handler set with On.
	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "default"}).
		OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"}).
		Build()

	matched, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
	if err != nil {
		panic(err)
	}

	unmatched, err := httpx.Get[myStruct](context.Background(), ts.URL+"/other")
	if err != nil {
		panic(err)
	}

	fmt.Println(matched.Result.Name)
	fmt.Println(unmatched.Result.Name)
	// Output:
	// Horton
	// default
}

func ExampleServerBuilder_OnRouteFunc() {
	t := &testing.T{}

	// OnRouteFunc registers a custom handler for a specific method/route.
	ts := httpxtest.NewServerBuilder(t).
		OnRouteFunc(http.MethodPost, "/v1/echo", func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)

			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write(body)
		}).
		Build()

	resp, err := httpx.Post[myStruct](context.Background(), ts.URL+"/v1/echo", myStruct{Name: "echo"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode, resp.Result.Name)
	// Output: 201 echo
}

func ExampleServerBuilder_BuildTLS() {
	t := &testing.T{}

	// BuildTLS serves over HTTPS. Clients must trust the test certificate;
	// here we skip verification.
	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "secure"}).
		BuildTLS()

	resp, err := httpx.Get[myStruct](context.Background(), ts.URL, httpx.WithInsecureSkipVerify())
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode, resp.Result.Name)
	// Output: 200 secure
}

func ExampleWithHeader() {
	t := &testing.T{}

	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "Horton"}, httpxtest.WithHeader("X-Custom", "go-snk")).
		Build()

	resp, err := httpx.Get[myStruct](context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Header.Get("X-Custom"))
	// Output: go-snk
}

func ExampleWithHeaders() {
	t := &testing.T{}

	headers := http.Header{}
	headers.Set("X-One", "1")
	headers.Set("X-Two", "2")

	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "Horton"}, httpxtest.WithHeaders(headers)).
		Build()

	resp, err := httpx.Get[myStruct](context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Header.Get("X-One"), resp.Header.Get("X-Two"))
	// Output: 1 2
}

func ExampleWithContentType() {
	t := &testing.T{}

	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "Horton"}, httpxtest.WithContentType("application/go-snk")).
		Build()

	resp, err := httpx.Get[myStruct](context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	// Output: application/go-snk
}

func ExampleWithCookie() {
	t := &testing.T{}

	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "Horton"}, httpxtest.WithCookie("session", "abc123")).
		Build()

	resp, err := httpx.Get[myStruct](context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Header.Get("Set-Cookie"))
	// Output: session=abc123
}

func ExampleWithDelay() {
	t := &testing.T{}

	ts := httpxtest.NewServerBuilder(t).
		On(http.StatusOK, myStruct{Name: "slow"}, httpxtest.WithDelay(10*time.Millisecond)).
		Build()

	start := time.Now()

	_, err := httpx.Get[myStruct](context.Background(), ts.URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Since(start) >= 10*time.Millisecond)
	// Output: true
}
