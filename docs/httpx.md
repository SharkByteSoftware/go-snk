<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# httpx

`httpx` provides lightweight helpers for making HTTP requests with less boilerplate.

It is designed to help you replace repetitive HTTP client code with small functions for:

- sending requests for common HTTP methods
- accessing raw responses when needed
- configuring request behavior
- decoding response bodies into typed values

## Overview

Use `httpx` when you want HTTP client code to be easier to read, reuse, and test.

It is especially useful when:

- the same request pattern appears in multiple places
- a helper makes the intent of the code clearer
- you want a consistent pattern for common HTTP methods instead of custom one-off clients

## When to use it

Use `httpx` when:

- you are writing client code that makes repeated HTTP requests
- you want to reduce request boilerplate
- you want clearer request and response handling around typed values

Prefer a simpler local implementation when:

- the request is one-off and unlikely to be reused
- you need fine-grained control over the HTTP client behavior
- a helper would obscure important details at the call site

## API reference

### Send requests

| Function | Purpose                                                 |
|----------|---------------------------------------------------------|
| `Get`    | Sends a GET request and returns a processed response    |
| `Post`   | Sends a POST request and returns a processed response   |
| `Put`    | Sends a PUT request and returns a processed response    |
| `Patch`  | Sends a PATCH request and returns a processed response  |
| `Delete` | Sends a DELETE request and returns a processed response |

### Access raw responses

| Function            | Purpose                                               |
|---------------------|-------------------------------------------------------|
| `GetRawResponse`    | Sends a GET request and returns the raw response      |
| `PostRawResponse`   | Sends a POST request and returns the raw response     |
| `PutRawResponse`    | Sends a PUT request and returns the raw response      |
| `PatchRawResponse`  | Sends a PATCH request and returns the raw response    |
| `DeleteRawResponse` | Sends a DELETE request and returns the raw response   |
| `Head`              | Sends a HEAD request and returns the raw response     |
| `Options`           | Sends an OPTIONS request and returns the raw response |

### Low-level access

| Function          | Purpose                                                                          |
|-------------------|----------------------------------------------------------------------------------|
| `DoRawRequest`    | Sends an HTTP request with a given method and body; returns the raw response     |
| `DoRequest`       | Sends an HTTP request with a given method and body; returns a decoded response   |
| `DecodeResponse`  | Decodes a raw `*http.Response` into a typed `Response[T]`                        |

### Configure

| Type            | Purpose                                           |
|-----------------|---------------------------------------------------|
| `ConfigOptions` | Configures request behavior and response handling |

#### Options

| Option               | Purpose                                                                  |
|----------------------|--------------------------------------------------------------------------|
| `WithHTTPClient`     | Sets the HTTP client used to send requests                               |
| `WithHeader`         | Adds a single header to the request                                      |
| `WithHeaders`        | Merges a set of headers into the request                                 |
| `WithTimeout`        | Sets the request timeout; must be positive                               |
| `WithParam`          | Adds a single query parameter to the request                             |
| `WithParams`         | Merges a set of query parameters into the request                        |
| `StrictDecoding`     | Enables strict JSON decoding; unknown fields cause an error              |
| `WithParseURLFunc`   | Overrides the function used to parse the request URL                     |
| `WithBearerToken`    | Sets `Authorization: Bearer <token>`; token must not be empty            |
| `WithBasicAuth`      | Sets `Authorization` using HTTP Basic auth (RFC 7617); username required |
| `WithUserAgent`      | Sets the `User-Agent` header; value must not be empty                    |

## Errors

`httpx` returns sentinel errors for common failure cases. Use `errors.Is` to check which kind of error occurred:

- `ErrResponse` — the server responded with a non-2xx status code.
- `ErrDecoding` — the response body could not be decoded into the target type.
- `ErrTransport` — the request failed to send or the response failed to read.
- `ErrEncoding` — request payload serialization failed before the request was sent.
- `ErrOptions` — one or more request options were invalid.

### Typed errors

Each sentinel has a corresponding typed error that carries additional context. Use `errors.As` to access the fields:

| Type             | Sentinel      | Fields                                          |
|------------------|---------------|-------------------------------------------------|
| `*ResponseError` | `ErrResponse` | `Method`, `URL`, `StatusCode`, `Status`, `Body` |
| `*DecodingError` | `ErrDecoding` | `ContentType`, `Err`                            |
| `*EncodingError` | `ErrEncoding` | `PayloadType`, `Err`                            |
| `*OptionsError`  | `ErrOptions`  | `Option`, `Message`, `Err`                      |

Example:
```go
resp, err := httpx.Get[MyType](ctx, url)
if err != nil {
    var respErr *httpx.ResponseError
    if errors.As(err, &respErr) {
        log.Printf("server returned %d for %s %s", respErr.StatusCode, respErr.Method, respErr.URL)
    }
}
```
## Notes

- Prefer the function that most clearly expresses your intent.
- Use raw response helpers when you need direct access to the underlying response.
- Use decoding helpers when you want to map response bodies into Go values.
- Keep request options focused so the call site stays readable.

## Examples

- [Examples](../httpx/httpx_example_test.go)
- [Unit tests](../httpx/httpx_test.go)