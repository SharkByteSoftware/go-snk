<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# jsonx

`jsonx` provides lightweight helpers for decoding JSON from common sources with less boilerplate.

It is designed to replace repetitive decoder setup with small typed functions for:

- decoding from a reader, byte slice, or string
- configuring decoding behaviour consistently across sources

## Overview

Use `jsonx` when you want JSON decoding to be easier to read and reuse without reaching for `encoding/json` directly every time.

It is especially useful when:

- the same decode pattern appears in multiple places
- you want consistent option handling across different input sources
- you want typed results without manually managing a `json.Decoder`

## When to use it

Use `jsonx` when:

- you are decoding JSON from an `io.Reader`, `[]byte`, or `string`
- you want a consistent pattern for common decode cases
- you want to configure decoding behaviour such as strict field checking or number handling

Prefer a direct `encoding/json` implementation when:

- you need fine-grained control over the decoder
- you are streaming large payloads token by token
- a helper would obscure important details at the call site

## API reference

### Decode

| Function        | Purpose                                        |
|-----------------|------------------------------------------------|
| `Decode`        | Decodes JSON from an `io.Reader` into T        |
| `DecodeBytes`   | Decodes JSON from a `[]byte` into T            |
| `DecodeString`  | Decodes JSON from a `string` into T            |

### Configure

| Option               | Purpose                                                                           |
|----------------------|-----------------------------------------------------------------------------------|
| `WithStrictDecoding` | Returns an error if the JSON contains fields not present in the target type       |
| `WithUseNumber`      | Decodes JSON numbers as `json.Number` instead of `float64` for `any` typed fields |

## Notes

- `Decode` is the base case — `DecodeBytes` and `DecodeString` delegate to it, so all options apply consistently across all three.
- `WithUseNumber` only has a visible effect when the target type contains `any` or `interface{}` fields. Concrete typed fields such as `int` or `float64` are unaffected.
- The caller is responsible for closing the reader passed to `Decode` if applicable.
- `jsonx` returns plain errors from `encoding/json`. Wrapping into typed errors is left to the caller.

## Examples

Examples can be found in the [test suite](../jsonx/decode_test.go).
