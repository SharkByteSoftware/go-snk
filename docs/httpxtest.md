<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# httpxtest

`httpxtest` provides a small builder for spinning up `httptest` servers with less boilerplate.

It is designed to help you stand up predictable HTTP servers in tests by:

- defining a default response for any request
- defining responses for specific method/route pairs
- responding with typed values, raw bytes, strings, or JSON
- shaping responses with headers, cookies, content types, and delays
- serving over plain HTTP or TLS

## Overview

Use `httpxtest` when you want test server setup to be easy to read and reuse.

It wraps `net/http/httptest` so you can describe what a server should return instead of wiring up handlers,
mux routing, and cleanup by hand. The builder registers `t.Cleanup` for you, so the server is closed
automatically when the test finishes.

It is especially useful when:

- you are testing client code that needs a server to respond in a specific way
- the same server setup appears across multiple tests
- you want to assert behavior around status codes, headers, cookies, or latency

## When to use it

Use `httpxtest` when:

- you are writing tests that need a real HTTP endpoint
- you want to reduce `httptest` setup boilerplate
- you want a consistent pattern for canned responses across a test suite

Prefer a plain `httptest.Server` when:

- you need full control over the handler or routing logic
- the test depends on behavior the builder does not model
- a builder would obscure important details of the test

## API reference

### Build a server

| Function           | Purpose                                          |
|--------------------|--------------------------------------------------|
| `NewServerBuilder` | Creates a new `ServerBuilder` bound to a `*testing.T` |
| `Build`            | Starts an HTTP test server and registers cleanup |
| `BuildTLS`         | Starts an HTTPS test server and registers cleanup |

### Define responses

| Method        | Purpose                                                                       |
|---------------|-------------------------------------------------------------------------------|
| `On`          | Sets the default handler to return a status code and response                 |
| `OnFunc`      | Sets the default handler to a custom `http.HandlerFunc`                       |
| `OnRoute`     | Registers a handler for a method/route that returns a status code and response |
| `OnRouteFunc` | Registers a custom `http.HandlerFunc` for a method/route                      |

`OnRoute` and `OnRouteFunc` panic if the method is empty or if a handler is already registered for the
same method/route pair. Requests that do not match a registered route fall through to the default handler,
which returns `500 Internal Server Error` unless overridden with `On` or `OnFunc`.

### Response values

`On` and `OnRoute` accept any value and write it based on its type:

| Value type        | Behavior                                                        |
|-------------------|-----------------------------------------------------------------|
| `nil`             | Writes `204 No Content`                                         |
| `string`          | Written verbatim as the body                                    |
| `[]byte`          | Written verbatim as the body                                    |
| `json.RawMessage` | Written verbatim as the body                                    |
| any other value   | JSON-encoded into the body; encoding failure yields a `500`     |

### Options

Options shape the response and can be passed at three levels: to `NewServerBuilder` (server level), to
`On`/`OnFunc` (default handler), or to `OnRoute`/`OnRouteFunc` (route level).

| Option            | Purpose                                          |
|-------------------|--------------------------------------------------|
| `WithHeader`      | Sets a single response header                    |
| `WithHeaders`     | Adds multiple response headers from `http.Header` |
| `WithContentType` | Sets the `Content-Type` header                   |
| `WithCookie`      | Sets a response cookie                           |
| `WithDelay`       | Delays the response by a fixed duration          |

Server-level options run on every request before the matched handler, then the handler's own options run.
This means route- or handler-level options take precedence over server-level options for the same header.

## Example

```go
func TestClient(t *testing.T) {
    ts := httpxtest.NewServerBuilder(t, httpxtest.WithHeader("X-Svr-Lvl", "server-level")).
        On(http.StatusOK, myStruct{Name: "default"}).
        OnRoute(http.MethodGet, "/v1/horton", http.StatusOK, myStruct{Name: "Horton"},
            httpxtest.WithContentType("application/json"),
            httpxtest.WithDelay(10*time.Millisecond)).
        Build()

    result, err := httpx.Get[myStruct](context.Background(), ts.URL+"/v1/horton")
    require.NoError(t, err)
    assert.Equal(t, "Horton", result.Result.Name)
}
```

## Notes

- The server is closed automatically via `t.Cleanup`; you do not need to close it yourself.
- Use `BuildTLS` when testing clients that must talk to an HTTPS endpoint.
- Use `On`/`OnRoute` for canned responses and `OnFunc`/`OnRouteFunc` when you need custom handler logic.
- Keep options focused so the test setup stays readable.

## Examples

- [Unit tests](../httpxtest/server_test.go)
