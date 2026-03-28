<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# sets

`sets` provides a generic, unordered collection of unique values backed by a map.

It is designed to help you replace repetitive set-logic code with small functions for:

- adding and removing items
- checking membership and equality
- performing set algebra operations
- iterating over items

## Overview

Use `sets` when you want set behavior that is easy to read, reuse, and test.

It is especially useful when:

- you need to track unique values without duplicates
- you need to compute intersections, unions, or differences between collections
- membership checks are a core part of the logic

## When to use it

Use `sets` when:

- uniqueness is a requirement, not just a preference
- set algebra operations like union or intersection are needed
- you want a named abstraction over map-based deduplication

Prefer a slice when:

- order matters
- duplicate values are intentional
- you need index-based access

## API reference

### Create a set

| Function | Purpose                                            |
|----------|----------------------------------------------------|
| `New`    | Creates a new set, optionally seeded with items    |

### Add or remove items

| Method   | Purpose                              |
|----------|--------------------------------------|
| `Add`    | Adds one or more items to the set    |
| `Remove` | Removes an item from the set         |

### Check membership

| Method     | Purpose                                        |
|------------|------------------------------------------------|
| `Contains` | Returns true if the item exists in the set     |
| `Equals`   | Returns true if both sets contain the same items |
| `Subset`   | Returns true if the set is a subset of another |

### Set algebra

| Method                | Purpose                                                  |
|-----------------------|----------------------------------------------------------|
| `Intersect`           | Returns a new set with items common to both sets         |
| `Union`               | Returns a new set with all items from both sets          |
| `Difference`          | Returns a new set with items in this set but not the other |
| `SymmetricDifference` | Returns a new set with items in either set but not both  |

### Inspect the set

| Method    | Purpose                               |
|-----------|---------------------------------------|
| `IsEmpty` | Returns true if the set has no items  |
| `Size`    | Returns the number of items in the set |
| `Values`  | Returns all items as a slice          |
| `Clone`   | Returns a shallow copy of the set     |
| `Clear`   | Removes all items from the set        |

### Iterate

| Method  | Purpose                                    |
|---------|--------------------------------------------|
| `Apply` | Calls a function for each item in the set  |

## Notes

- Iteration order is not guaranteed; sets are backed by a Go map.
- `Intersect` and `Union` optimize by iterating the smaller set.
- `New` accepts optional seed items, making initialization concise.
- `Clone` produces an independent copy; mutations do not affect the original.

## Examples

Examples can be found in the [test suite](../../containers/sets/set_test.go).
