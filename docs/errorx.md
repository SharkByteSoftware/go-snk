<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# errorx

`errorx` provides small helpers for common error handling patterns.

It is designed to help you replace repetitive error handling code with small functions for:

- explicitly discarding errors with intent
- panicking on non-recoverable errors at initialization time
- checking an error against multiple targets in one call

## Overview

Use `errorx` when you want error handling code to be easier to read, reuse, and test.

It is especially useful when:

- you want to make intentional error suppression visible rather than silent
- a value must be valid at startup and an error represents a misconfiguration
- you need to match an error against several sentinel values without chaining `errors.Is`

## When to use it

Use `errorx` when:

- a small helper improves readability or makes intent explicit
- the same error handling pattern appears in multiple places

Prefer a direct `if err != nil` when:

- the error handling requires custom logic or recovery behavior
- the surrounding code already makes the intent obvious

## API reference

### Discard errors

| Function | Purpose                                                                  |
|----------|--------------------------------------------------------------------------|
| `Ignore` | Explicitly discards an error, documenting the suppression as intentional |

### Panic on failure

| Function | Purpose                                                               |
|----------|-----------------------------------------------------------------------|
| `Must`   | Returns the value if err is nil; panics otherwise                     |

### Check against multiple targets

| Function   | Purpose                                                                        |
|------------|--------------------------------------------------------------------------------|
| `IsAny`    | Reports whether an error matches any of the provided targets using `errors.Is` |

### Reduce multiple errors

| Function   | Purpose                                                                  |
|------------|--------------------------------------------------------------------------|
| `FirstErr` | Returns the first non-nil error from a list, or nil if all are nil       |

## Notes

- `Ignore` is a named alternative to `_ = someFunc()`. It signals to readers that the error is intentionally discarded, not overlooked.
- `Must` is intended for use at program initialization time. Avoid it in request or business logic paths where errors should be handled gracefully.
- `IsAny` uses `errors.Is` semantics for each comparison, so wrapped errors are matched correctly.
- `FirstErr` is useful for collapsing a set of validation or initialization errors into a single result without chaining multiple `if` statements.

## Examples

- [Examples](../errorx/errorx_example_test.go)
- [Unit tests](../errorx/errorx_test.go)
