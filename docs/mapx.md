<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# mapx

`mapx` provides generic helpers for working with Go maps in a clear and reusable way.

It is designed to help you replace repetitive map loops with small functions for:

- extracting keys and values
- checking for key presence
- transforming maps into other shapes
- filtering entries
- combining or inverting maps

## Overview

Use `mapx` when you want map logic to be easier to read, reuse, and test.

It is especially useful when:

- the same map loop appears in multiple places
- a helper makes the intent of the code clearer
- you want type-safe generic utilities instead of custom one-off helpers

## When to use it

Use `mapx` when:

- you need a common map operation expressed clearly
- you want to avoid repeating small loops throughout the codebase
- you prefer reusable generic helpers over ad hoc implementations

Prefer a simpler local loop when:

- the operation is tiny and only used once
- a helper would make the code less obvious
- performance or allocation behavior needs a specialized implementation

## API reference

### Extract keys or values

| Function       | Purpose                                  |
|----------------|------------------------------------------|
| `Keys`         | Returns all keys from a map              |
| `Values`       | Returns all values from a map            |
| `UniqueValues` | Returns only unique values from a map    |

### Look up or check entries

| Function   | Purpose                                                         |
|------------|-----------------------------------------------------------------|
| `Contains` | Returns true if the map contains all of the specified keys      |
| `ValueOr`  | Returns the value for a key, or a fallback if the key is absent |

### Transform or reshape data

| Function  | Purpose                                                         |
|-----------|-----------------------------------------------------------------|
| `ToSlice` | Converts a map into a slice using a mapper function             |
| `Invert`  | Swaps map keys and values                                       |
| `Combine` | Merges multiple maps into one; last writer wins on key conflict |

### Filter or visit entries

| Function | Purpose                                                         |
|----------|-----------------------------------------------------------------|
| `Filter` | Returns a map containing only entries that satisfy a predicate  |
| `Apply`  | Runs a function on each map entry for side effects              |

## Notes

- Prefer the function that most clearly expresses your intent.
- Prefer the simplest helper that matches the operation.
- Map iteration order is not guaranteed; keep this in mind when ordering matters in your calling code.
- Check each function's documentation for details such as key-conflict behavior and zero-value handling.

## Examples

Examples can be found in the [mapx examples](../mapx/map_example_test.go) and in the [test suite](../mapx/map_test.go).