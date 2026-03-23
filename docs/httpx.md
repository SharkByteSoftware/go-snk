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
| `Delete` | Sends a DELETE request and returns a processed response |

### Access raw responses

| Function            | Purpose                                             |
|---------------------|-----------------------------------------------------|
| `GetRawResponse`    | Sends a GET request and returns the raw response    |
| `PostRawResponse`   | Sends a POST request and returns the raw response   |
| `PutRawResponse`    | Sends a PUT request and returns the raw response    |
| `DeleteRawResponse` | Sends a DELETE request and returns the raw response |

### Configure and decode

| Function / Type        | Purpose                                           |
|------------------------|---------------------------------------------------|
| `ConfigOptions`        | Configures request behavior and response handling |
| `AlwaysIncludeRawBody` | Configures responses to retain the raw body       |
| `DecodeRawBody`        | Decodes a raw response body into a target value   |

## Errors

`httpx` returns sentinel errors for common failure cases:

- `ErrNon2xxStatusCode` — the server responded with a non-2xx status code. The response body is stored in 
- `Response.RawBody` for inspection.
- `ErrDecoding` — the response body could not be decoded into the target type.
- `ErrTransport` — the response body could not be read due to a transport/read failure.
- `ErrMarshaling` — request payload serialization failed before the request was sent.

### Inspecting error responses

When a request returns `ErrNon2xxStatusCode`, the typed result is not decoded. Use `Response.RawBody` or 
`DecodeRawBody` to parse the error payload into a custom type.

## Notes

- Prefer the function that most clearly expresses your intent.
- Use raw response helpers when you need direct access to the underlying response.
- Use decoding helpers when you want to map response bodies into Go values.
- Keep request options focused so the call site stays readable.

## Examples

Examples can be found in the [httpx examples](../httpx/httpx_example_test.go) and in the [test suite](../httpx/httpx_test.go).