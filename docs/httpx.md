# httpx

`httpx` provides lightweight helpers for making HTTP requests with less boilerplate.

It is useful when you want:
- concise request helpers
- typed response handling
- simple request configuration
- a consistent pattern for common HTTP methods

## Overview

The package is intended as a small convenience layer over standard HTTP client workflows.

## What it helps with

- GET, POST, PUT, and DELETE requests
- raw response access when needed
- request configuration
- response decoding
- retaining or exposing raw response bodies
- keeping client code shorter and more focused

## When to use it

Use `httpx` when:
- you are writing client code that makes repeated HTTP requests
- you want to reduce request boilerplate
- you want clearer request/response handling around typed values

## API reference

### Request helpers

| Function | Purpose                                                 |
|----------|---------------------------------------------------------|
| `Get`    | Sends a GET request and returns a processed response    |
| `Post`   | Sends a POST request and returns a processed response   |
| `Put`    | Sends a PUT request and returns a processed response    |
| `Delete` | Sends a DELETE request and returns a processed response |

### Raw response helpers

| Function            | Purpose                                             |
|---------------------|-----------------------------------------------------|
| `GetRawResponse`    | Sends a GET request and returns the raw response    |
| `PostRawResponse`   | Sends a POST request and returns the raw response   |
| `PutRawResponse`    | Sends a PUT request and returns the raw response    |
| `DeleteRawResponse` | Sends a DELETE request and returns the raw response |

### Configuration and response handling

| Function / Type        | Purpose                                           |
|------------------------|---------------------------------------------------|
| `ConfigOptions`        | Configures request behavior and response handling |
| `AlwaysIncludeRawBody` | Configures responses to retain the raw body       |
| `DecodeRawBody`        | Decodes a raw response body into a target value   |

## Notes

- Use raw response helpers when you need direct access to the underlying response.
- Use decoding helpers when you want to map response bodies into Go values.
- Keep request options focused so the call site stays readable.

## Examples

See [`docs/examples.md`](examples.md) for short usage examples and guidance.