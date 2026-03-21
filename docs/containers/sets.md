<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# Sets

This document covers the set implementation in `containers`.

## Overview

The set type provides a generic collection for unique items.

It is useful when you need:
- uniqueness
- set comparisons
- standard set operations
- simple membership checks

## Common operations

- create a new set
- add items
- remove items
- check membership
- check whether the set is empty
- compare two sets
- clear the set
- return all values
- intersect sets
- union sets
- compute differences
- compute symmetric differences
- check subset relationships
- apply a function to each item

## API reference

### `Set`

| Method                | Purpose                                         |
|-----------------------|-------------------------------------------------|
| `New`                 | Creates a new set from initial values           |
| `Add`                 | Adds one or more items                          |
| `IsEmpty`             | Reports whether the set contains no items       |
| `Equals`              | Reports whether two sets contain the same items |
| `Contains`            | Reports whether the set contains an item        |
| `Remove`              | Removes an item from the set                    |
| `Size`                | Returns the number of items                     |
| `Clear`               | Removes all items                               |
| `Values`              | Returns all values as a slice                   |
| `Intersect`           | Returns items common to both sets               |
| `Union`               | Returns items from both sets                    |
| `Difference`          | Returns items in one set but not the other      |
| `SymmetricDifference` | Returns items in either set but not both        |
| `Subset`              | Reports whether one set is a subset of another  |
| `Apply`               | Applies a function to each item                 |

## Notes

- Sets are a good fit when uniqueness matters more than ordering.
- If ordering is important, consider whether a different collection is a better fit.

## When to use it

Use the set when:
- you need to avoid duplicates
- you want to compare membership across collections
- you need common set algebra operations

## Examples
