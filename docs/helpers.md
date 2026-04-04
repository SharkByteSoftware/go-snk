<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# helpers

`helpers` provides small generic utilities for working with pointers and zero values.

It is designed to help you replace repetitive pointer-handling and zero-value patterns with small functions for:

- creating and dereferencing pointers safely
- checking for nil pointers
- working with zero/empty values of any type

## Overview

Use `helpers` when you want pointer and value utilities that are easy to read, reuse, and test.

It is especially useful when:

- you need to take the address of a literal or computed value inline
- you want safe dereferencing with a fallback instead of a nil check
- you need to check whether a value is the zero value of its type

## When to use it

Use `helpers` when:

- you need a pointer to a value that cannot be addressed directly (e.g. a function return value or literal)
- you want to dereference a pointer with a default rather than writing an explicit nil check
- you need a typed nil or zero value in generic code

Prefer a direct expression when:

- the operation is a simple one-off and the intent is already clear
- the surrounding code already handles nil or zero cases explicitly

## API reference

### Pointer utilities

| Function    | Purpose                                                                          |
|-------------|----------------------------------------------------------------------------------|
| `AsPtr`     | Returns a pointer to the given value                                             |
| `AsValue`   | Dereferences a pointer; returns the type's zero value if the pointer is nil      |
| `AsValueOr` | Dereferences a pointer; returns a specified fallback value if the pointer is nil |
| `Nil`       | Returns a typed nil pointer                                                      |
| `IsNil`     | Returns true if the pointer is nil                                               |

### Zero value utilities

| Function  | Purpose                                              |
|-----------|------------------------------------------------------|
| `Empty`   | Returns the zero value of a type                     |
| `IsEmpty` | Returns true if the value equals the type's zero value |

## Notes

- `AsValue` is shorthand for `AsValueOr(ptr, Empty[T]())`.
- `IsEmpty` requires a `comparable` type constraint; `Empty` works with any type.
- `Nil` and `Empty` are useful in generic code where you need a typed zero or nil without a concrete value.

## Examples

Examples can be found in the [test suite](../helpers/ptrs_test.go).
