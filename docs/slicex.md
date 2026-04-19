<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# slicex

`slicex` provides generic helpers for working with Go slices in a clear and reusable way.

It is designed to help you replace repetitive slice loops with small functions for:

- filtering values
- mapping slices to new types
- searching for items
- grouping and partitioning data
- removing duplicates
- combining slices into pairs
- creating sliding windows
- performing set-like operations
- simple numeric aggregation

## Overview

Use `slicex` when you want slice logic to be easier to read, reuse, and test.

It is especially useful when:

- the same slice loop appears in multiple places
- a helper makes the intent of the code clearer
- you want type-safe generic utilities instead of custom one-off helpers

## When to use it

Use `slicex` when:

- you need a common slice operation expressed clearly
- you want to avoid repeating small loops throughout the codebase
- you prefer reusable generic helpers over ad hoc implementations

Prefer a simpler local loop when:

- the operation is tiny and only used once
- a helper would make the code less obvious
- performance or allocation behavior needs a specialized implementation

## API reference

### Get the first matching or default value

| Function       | Purpose                                                              |
|----------------|----------------------------------------------------------------------|
| `FirstOr`      | Returns the first element of a slice or a fallback value             |
| `FirstOrEmpty` | Returns the first element of a slice or the zero value               |
| `FirstBy`      | Returns the first item satisfying a predicate, and a found boolean   |
| `FirstOrBy`    | Returns the first item satisfying a predicate, or a fallback value   |
| `Find`         | Returns the first item in a slice equal to a given value             |
| `FindBy`       | Returns the first item in a slice that matches a predicate           |
| `FindOr`       | Returns the first item equal to a given value, or a fallback         |
| `FindOrBy`     | Returns the first item matching a predicate, or a fallback value     |

### Get the last matching or default value

| Function       | Purpose                                                              |
|----------------|----------------------------------------------------------------------|
| `LastOr`       | Returns the last element of a slice or a fallback value              |
| `LastOrEmpty`  | Returns the last element of a slice or the zero value                |
| `LastBy`       | Returns the last item satisfying a predicate, and a found boolean    |
| `LastOrBy`     | Returns the last item satisfying a predicate, or a fallback value    |

### Test for presence or match

| Function   | Purpose                                            |
|------------|----------------------------------------------------|
| `Contains` | Returns true if the slice contains a given value   |
| `Any`      | Returns true if any item matches a given value     |
| `AnyBy`    | Returns true if any item matches a predicate       |
| `All`      | Returns true if all items match a given value      |
| `AllBy`    | Returns true if all items satisfy a predicate      |
| `None`     | Returns true if no item equals a given value       |
| `NoneBy`   | Returns true if no item satisfies a predicate      |

### Filter and select items

| Function          | Purpose                                                              |
|-------------------|----------------------------------------------------------------------|
| `Filter`          | Returns only items that satisfy a predicate                          |
| `FilterWithIndex` | Returns only items whose predicate receives both the index and value |
| `Compact`         | Removes zero values from a slice                                     |

### Transform or reshape data

| Function             | Purpose                                                                    |
|----------------------|----------------------------------------------------------------------------|
| `Map`                | Transforms each element into a new slice                                   |
| `MapWithIndex`       | Transforms each element using a function that also receives the index      |
| `FilterMap`          | Filters and transforms elements in one pass                                |
| `FilterMapWithIndex` | Filters and transforms elements in one pass, with index access             |
| `Bind`               | Maps each item to a slice and flattens the results (flatMap)               |
| `Flatten`            | Collapses a slice of slices into a single flat slice                       |
| `Reduce`             | Folds slice values into a single accumulated result                        |
| `ToMap`              | Converts a slice into a map using a key selector                           |
| `Apply`              | Runs a function on each item for side effects; does not return a new slice |
| `ApplyWithIndex`     | Runs a function on each item for side effects, also receiving the index    |

### Remove duplicates or keep unique values

| Function    | Purpose                                          |
|-------------|--------------------------------------------------|
| `Unique`    | Removes duplicate values from a slice            |
| `UniqueBy`  | Removes duplicates by a derived key              |
| `UniqueMap` | Transforms items and returns only unique results |

### Count elements

| Function  | Purpose                                                    |
|-----------|------------------------------------------------------------|
| `Count`   | Returns the number of elements equal to a given value      |
| `CountBy` | Returns the number of elements satisfying a predicate      |

### Reorder slices

| Function  | Purpose                                                                           |
|-----------|-----------------------------------------------------------------------------------|
| `Reverse` | Returns a reversed copy of the slice                                              |
| `Rotate`  | Returns a copy with elements shifted left by n positions; negative n shifts right |
| `Sort`    | Returns a sorted copy of the slice using natural ordering                         |
| `SortBy`  | Returns a sorted copy of the slice using a comparison function                    |

### Group, split, or combine collections

| Function    | Purpose                                                                          |
|-------------|----------------------------------------------------------------------------------|
| `GroupBy`   | Groups items into a map by a computed key                                        |
| `Partition` | Splits items into two slices based on a predicate                                |
| `Chunk`     | Splits a slice into sub-slices of at most size n                                 |
| `Zip`       | Combines two slices into a slice of `Pair` values, pairing elements by index     |
| `Window`    | Returns overlapping sub-slices of a fixed size, advancing one position at a time |

### Find by index

| Function  | Purpose                                                              |
|-----------|----------------------------------------------------------------------|
| `IndexOf` | Returns the index of the first element equal to a given value, or -1 |
| `IndexBy` | Returns the index of the first element satisfying a predicate, or -1 |

### Perform set-like operations

| Function     | Purpose                                               |
|--------------|-------------------------------------------------------|
| `Intersect`  | Returns values common to both slices                  |
| `Union`      | Returns all unique values from both slices            |
| `Difference` | Returns values present in one slice but not the other |

### Aggregate numeric values

| Function    | Purpose                                               |
|-------------|-------------------------------------------------------|
| `Sum`       | Calculates the sum of numeric values                  |
| `SumBy`     | Calculates the sum using a value selector             |
| `Product`   | Calculates the product of numeric values              |
| `ProductBy` | Calculates the product using a value selector         |
| `Mean`      | Calculates the arithmetic mean                        |
| `MeanBy`    | Calculates the mean using a value selector            |
| `Min`       | Returns the minimum value                             |
| `MinBy`     | Returns the minimum value using a comparison function |
| `Max`       | Returns the maximum value                             |
| `MaxBy`     | Returns the maximum value using a comparison function |
## Notes

- Prefer the function that most clearly expresses your intent.
- Prefer the simplest helper that matches the operation.
- Check each function’s documentation for details such as ordering, stability, and zero-value behavior.
- For very large workloads, consider whether a specialized implementation would be more appropriate.
- `FirstBy` and `FirstOrBy` delegate to `FindBy` and `FindOrBy` respectively; they exist for naming symmetry with `LastBy` and `LastOrBy`.
- `Zip` returns a `Pair[A, B]` value for each position; the result length equals the shorter of the two input slices.
- `Window` returns an empty slice when `size` is less than 1 or greater than the length of the input slice.
- `Chunk` panics if the chunk size is less than 1.
- `Flatten` skips nil inner slices.
- `IndexOf` and `IndexBy` return -1 when no match is found.
- `Rotate` with a positive n shifts left; negative n shifts right. Values wrap around.

## Examples

- [Examples](../slicex/slicex_example_test.go)
- [Unit tests](../slicex/slicex_test.go)

