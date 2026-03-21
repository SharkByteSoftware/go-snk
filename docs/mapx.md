<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# mapx

`mapx` provides generic helpers for working with maps in a simple and reusable way.

It is useful when you want to:
- inspect map keys or values
- check for key presence
- transform map data into other shapes
- filter entries
- combine multiple maps
- invert mappings where that makes sense

## Overview

`mapx` focuses on common map operations that often require short custom loops. The goal is to make those operations concise while keeping the code easy to understand.

## Common capabilities

- extract keys
- extract values
- collect unique values
- check whether keys are present
- return a value or fallback
- invert maps
- combine maps
- convert a map into a slice
- filter entries
- apply a function to each entry

## When to use it

Use `mapx` when:
- your code repeatedly extracts or transforms map data
- you want a compact helper for map-specific operations
- readability improves when a common map pattern becomes a named function

## API reference

| Function       | Purpose                                             |
|----------------|-----------------------------------------------------|
| `Keys`         | Returns all keys from a map                         |
| `Values`       | Returns all values from a map                       |
| `UniqueValues` | Returns unique values from a map                    |
| `Contains`     | Reports whether the map contains the specified keys |
| `ValueOr`      | Returns a value for a key or a fallback value       |
| `Invert`       | Swaps keys and values in a map                      |
| `Combine`      | Merges multiple maps into one                       |
| `ToSlice`      | Converts a map into a slice using a mapper          |
| `Filter`       | Returns a map containing only matching entries      |
| `Apply`        | Applies a function to each key-value pair           |

## Notes

- Some helpers assume the map shape is suitable for the requested transformation.
- Keep an eye on whether ordering matters in your calling code, since map iteration order is not guaranteed.

## Examples

See [`docs/examples.md`](examples.md) for usage-oriented examples.