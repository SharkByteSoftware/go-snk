<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# slicex

`slicex` provides generic helpers for working with slices in a clear and reusable way.

It is useful when you want to:
- filter values
- map slices to new types
- search for items
- group and partition data
- remove duplicates
- perform simple numeric aggregation

## Overview

`slicex` is designed to replace repetitive slice loops with small helper functions that are easy to read and reuse.

## Common capabilities

### Selection and transformation
- first-item helpers
- filtering
- mapping
- filter-then-map workflows
- flattening and reduction-style operations

### Search and predicates
- find values
- check whether any item matches
- check whether all items match

### Deduplication and ordering
- unique values
- unique-by-key behavior
- reversing slices
- compacting slices

### Grouping and map conversion
- grouping by key
- converting a slice to a map
- partitioning into two groups

### Set-like operations
- intersection
- union
- difference

### Numeric helpers
- sum
- product
- mean
- min
- max
- value-based variants

## When to use it

Use `slicex` when:
- the same slice loop appears in multiple places
- a helper makes the code easier to scan
- you want type-safe generic utilities instead of custom ad hoc helpers

## API reference

### Slice selection and transformation

| Function       | Purpose                                                       |
|----------------|---------------------------------------------------------------|
| `FirstOr`      | Returns the first element of a slice or a fallback value      |
| `FirstOrEmpty` | Returns the first element of a slice or the zero value        |
| `Filter`       | Returns a slice containing only values that match a predicate |
| `Map`          | Transforms each element into a new slice                      |
| `FilterMap`    | Filters and transforms a slice in one pass                    |
| `UniqueMap`    | Maps a slice and removes duplicates from the result           |
| `Reduce`       | Accumulates or flattens slice values into another result      |
| `Apply`        | Applies a function to each element in a slice                 |
| `Reverse`      | Returns a slice with elements in reverse order                |
| `Compact`      | Returns a slice with zero values removed                      |

### Search and predicates

| Function | Purpose                                               |
|----------|-------------------------------------------------------|
| `Find`   | Returns the first element that matches a condition    |
| `FindBy` | Returns the first element that matches a predicate    |
| `FindOr` | Returns the first matching element or a default value |
| `Any`    | Reports whether any element matches a condition       |
| `AnyBy`  | Reports whether any element matches a predicate       |
| `All`    | Reports whether all elements equal a given value      |
| `AllBy`  | Reports whether all elements satisfy a predicate      |

### Uniqueness and grouping

| Function    | Purpose                                             |
|-------------|-----------------------------------------------------|
| `Unique`    | Removes duplicate values from a slice               |
| `UniqueBy`  | Removes duplicates using a predicate or key rule    |
| `ToMap`     | Converts a slice into a map using a key selector    |
| `GroupBy`   | Groups elements by key                              |
| `Partition` | Splits a slice into two groups based on a predicate |

### Set-like operations

| Function     | Purpose                                               |
|--------------|-------------------------------------------------------|
| `Intersect`  | Returns values common to two slices                   |
| `Union`      | Returns values from both slices                       |
| `Difference` | Returns values present in one slice but not the other |

### Numeric helpers

| Function    | Purpose                                               |
|-------------|-------------------------------------------------------|
| `Sum`       | Calculates the sum of numeric values                  |
| `SumBy`     | Calculates the sum using a value selector             |
| `Product`   | Calculates the product of numeric values              |
| `ProductBy` | Calculates the product using a value selector         |
| `Mean`      | Calculates the arithmetic mean                        |
| `MeanBy`    | Calculates the mean using a value selector            |
| `Max`       | Returns the maximum value                             |
| `MaxBy`     | Returns the maximum value using a comparison function |
| `Min`       | Returns the minimum value                             |
| `MinBy`     | Returns the minimum value using a comparison function |

## Notes

- Choose the function that most clearly expresses your intent.
- Prefer the simplest helper that matches the operation.
- For very large workloads, consider whether a specialized approach is more appropriate.

## Examples
