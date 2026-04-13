<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# stringx

`stringx` provides small helpers for common string operations.

It is designed to help you replace repetitive string handling code with small functions for:

- checking whether a string is blank
- returning the first non-empty string from a list of candidates
- trimming a string to a maximum length
- wrapping a string with a prefix and suffix
- padding a string on the left or right to a target length

## Overview

Use `stringx` when you want string manipulation code to be more concise and readable.

It is especially useful when:

- you need a fallback string value from several candidates
- you are formatting output that requires fixed-width columns or aligned text
- you want to guard against blank user input without writing `strings.TrimSpace` inline

## When to use it

Use `stringx` when:

- a small helper makes the intent of a string operation explicit
- the same string handling pattern appears in multiple places

Prefer standard library calls when:

- the operation is a single `strings` function call that already reads clearly
- the surrounding code already makes the intent obvious

## API reference

### Blank checks

| Function  | Purpose                                                         |
|-----------|-----------------------------------------------------------------|
| `IsBlank` | Returns true if the string is empty or contains only whitespace |

### Fallback selection

| Function        | Purpose                                                              |
|-----------------|----------------------------------------------------------------------|
| `Coalesce`      | Returns the first non-empty string from the provided values          |
| `CoalesceFunc`  | Returns the first string satisfying a caller-provided predicate      |

### Length limiting

| Function   | Purpose                                                              |
|------------|----------------------------------------------------------------------|
| `Truncate` | Returns the string trimmed to a maximum length (Unicode-safe)        |

### Formatting

| Function   | Purpose                                                              |
|------------|----------------------------------------------------------------------|
| `Wrap`     | Surrounds the string with a given prefix and suffix                  |
| `PadLeft`  | Left-pads the string with a character to the specified length        |
| `PadRight` | Right-pads the string with a character to the specified length       |

## Notes

- `Truncate` operates on runes, not bytes, so it is safe for multi-byte Unicode characters.
- `PadLeft` and `PadRight` also operate on runes, so padding counts by character rather than byte.
- `Coalesce` returns an empty string if all provided values are empty.
- `CoalesceFunc` is useful when the definition of "non-empty" is caller-defined — for example, skipping blank strings rather than just empty ones.

## Examples

- [Examples](../stringx/stringx_example_test.go)
- [Unit tests](../stringx/stringx_test.go)
