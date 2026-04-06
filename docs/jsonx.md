<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# jsonx

`jsonx` provides lightweight helpers for encoding and decoding JSON from common sources with less boilerplate.

It is designed to replace repetitive encoder and decoder setup with small typed functions for:

- decoding from a reader, byte slice, string, or file
- encoding to a writer, byte slice, string, or file
- configuring encoding and decoding behaviour consistently across sources

## Overview

Use `jsonx` when you want JSON encoding and decoding to be easier to read and reuse without reaching for `encoding/json` directly every time.

It is especially useful when:

- the same encode or decode pattern appears in multiple places
- you want consistent option handling across different input sources
- you want typed results without manually managing a `json.Encoder` or `json.Decoder`

## When to use it

Use `jsonx` when:

- you are decoding JSON from an `io.Reader`, `[]byte`, `string`, or file path
- you are encoding a value to JSON and want the result as a writer output, `[]byte`, `string`, or file
- you want a consistent pattern for common encode and decode cases
- you want to configure behaviour such as strict field checking, number handling, HTML escaping, or indentation

Prefer a direct `encoding/json` implementation when:

- you need fine-grained control over the encoder or decoder
- you are streaming large payloads token by token
- a helper would obscure important details at the call site

## API reference

### Decode

| Function          | Purpose                                        |
|-------------------|------------------------------------------------|
| `Decode`          | Decodes JSON from an `io.Reader` into T        |
| `DecodeBytes`     | Decodes JSON from a `[]byte` into T            |
| `DecodeString`    | Decodes JSON from a `string` into T            |
| `DecodeFromFile`  | Decodes JSON from a file path into T           |

### Encode

| Function        | Purpose                                               |
|-----------------|-------------------------------------------------------|
| `Encode`        | Encodes a value as JSON into an `io.Writer`           |
| `EncodeBytes`   | Encodes a value as JSON and returns a `[]byte`        |
| `EncodeString`  | Encodes a value as JSON and returns a `string`        |
| `EncodeToFile`  | Encodes a value as JSON into a file at a given path   |

### Configure decoding

| Option               | Purpose                                                                            |
|----------------------|------------------------------------------------------------------------------------|
| `WithStrictDecoding` | Returns an error if the JSON contains fields not present in the target type        |
| `WithUseNumber`      | Decodes JSON numbers as `json.Number` instead of `float64` for `any` typed fields |

### Configure encoding

| Option            | Purpose                                                                                    |
|-------------------|--------------------------------------------------------------------------------------------|
| `WithEscapeHTML`  | Enables escaping of HTML characters (`<`, `>`, `&`) in the output; disabled by default    |
| `WithIndent`      | Enables pretty-printing with a given indent string per level (e.g. `"\t"` or `"  "`)      |

## Notes

- `Decode` is the base case — `DecodeBytes`, `DecodeString`, and `DecodeFromFile` delegate to it, so all decode options apply consistently across all four.
- `Encode` is the base case — `EncodeBytes`, `EncodeString`, and `EncodeToFile` delegate to it, so all encode options apply consistently across all four.
- `WithUseNumber` only has a visible effect when the target type contains `any` or `interface{}` fields. Concrete typed fields such as `int` or `float64` are unaffected.
- `WithEscapeHTML` is disabled by default, unlike the standard library which enables it. Enable it when output may be embedded in HTML.
- The caller is responsible for closing the reader passed to `Decode` if applicable.
- `jsonx` returns plain errors from `encoding/json`. Wrapping into typed errors is left to the caller.

## Examples

- [Decode test suite](../jsonx/decode_test.go)
- [Encode test suite](../jsonx/encode_test.go)
