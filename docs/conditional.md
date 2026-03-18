# conditional

`conditional` provides small helpers for simple branching logic and value selection.

It is useful when you want to:
- choose between two values
- call a function only if something is non-nil
- execute one of two callbacks based on a condition
- express small conditional decisions more compactly

## Overview

The package exists for situations where a helper can make intent clearer than repeating the same `if` pattern.

## Common capabilities

- return one of two values based on a boolean
- call a function only when a pointer is present
- choose between two functions at runtime
- return a result from one of two branches

## When to use it

Use `conditional` when:
- a small branching helper improves readability
- the code would otherwise be repetitive
- the helper is more expressive than an inline conditional

## API reference

| Function       | Purpose                                           |
|----------------|---------------------------------------------------|
| `If`           | Returns one of two values based on a condition    |
| `IfNotNil`     | Calls a function only when a pointer is not nil   |
| `IfCall`       | Calls one of two functions based on a condition   |
| `IfCallReturn` | Calls one of two functions and returns the result |

## Notes

This package is intentionally small. It works best for concise, readable selection logic rather than complex branching.

## Examples

See [`docs/examples.md`](examples.md) for short usage-oriented examples.