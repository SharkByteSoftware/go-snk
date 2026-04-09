<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# conditional

`conditional` provides small helpers for simple branching logic and value selection.

It is designed to help you replace repetitive conditional patterns with small functions for:

- choosing between two values
- calling a function only when a pointer is non-nil
- executing one of two callbacks based on a condition
- returning a result from one of two branches
- looking up a value by key with a fallback

## Overview

Use `conditional` when you want branching logic to be easier to read, reuse, and test.

It is especially useful when:

- the same `if` pattern appears in multiple places
- a helper makes the intent of the code clearer
- you want a named function to express a small conditional decision rather than an inline expression

## When to use it

Use `conditional` when:

- a small branching helper improves readability
- the code would otherwise be repetitive
- the helper is more expressive than an inline conditional

Prefer a simpler inline expression when:

- the condition is used only once
- the branching logic is complex enough that a named helper would obscure it
- the surrounding code already makes the intent obvious

## API reference

### Select a value

| Function | Purpose                                        |
|----------|------------------------------------------------|
| `If`     | Returns one of two values based on a condition |
| `Switch` | Returns a value from a map by key, with a fallback |

### Call a function conditionally

| Function       | Purpose                                           |
|----------------|---------------------------------------------------|
| `IfNotNil`     | Calls a function only when a pointer is not nil   |
| `IfCall`       | Calls one of two functions based on a condition   |
| `IfCallReturn` | Calls one of two functions and returns the result |

### `Switch`

`Switch` returns the value associated with `key` in the `cases` map. If the key is not present, `fallback` is returned.

```go
label := conditional.Switch(status, map[int]string{
    1: "active",
    2: "inactive",
    3: "pending",
}, "unknown")
```

Use `Switch` when:

- you are selecting a value from a fixed set of keys
- a `switch` statement would otherwise be used purely for value lookup
- you want a concise, readable alternative to repeated `if/else` chains

## Notes

- Prefer the function that most clearly expresses your intent.
- This package is intentionally small. It works best for concise, readable selection logic rather than complex branching.

## Examples

Examples can be found in the [test suite](../conditional/conditionals_test.go).