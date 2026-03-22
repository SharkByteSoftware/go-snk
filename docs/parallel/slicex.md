<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# slicex/parallel

`parallel` provides parallel versions of selected `slicex` helpers for workloads where each item can be processed independently.

It is designed to help you replace sequential slice loops with concurrent functions for:

- mapping slice values using parallel work
- applying a function to each item concurrently
- grouping and partitioning values using parallel predicate evaluation
- limiting concurrency for better control over resource usage

## Overview

Use `slicex/parallel` when you want concurrent slice processing that follows the same patterns as `slicex`.

It is especially useful when:

- the work done per item is independent and expensive enough to benefit from concurrency
- you want results returned in the same order as the input
- you want to cap the number of goroutines to avoid excessive resource usage

## When to use it

Use `slicex/parallel` when:

- each item can be processed independently
- the work per item is expensive enough to benefit from concurrency
- you want to keep results ordered
- you want to cap concurrency to avoid excessive goroutine usage

Prefer sequential `slicex` helpers when:

- the work per item is cheap and concurrency overhead would outweigh the benefit
- items cannot be processed independently
- shared mutable state makes concurrent access difficult to reason about

## API reference

### Transform items in parallel

| Function       | Purpose                                                             |
|----------------|---------------------------------------------------------------------|
| `Map`          | Transforms a slice using a mapper function in parallel              |
| `MapWithLimit` | Transforms a slice using a mapper function with limited concurrency |

### Apply a function to each item in parallel

| Function         | Purpose                                                  |
|------------------|----------------------------------------------------------|
| `Apply`          | Applies a function to each item in parallel              |
| `ApplyWithLimit` | Applies a function to each item with limited concurrency |

### Group or split items in parallel

| Function             | Purpose                                                          |
|----------------------|------------------------------------------------------------------|
| `GroupBy`            | Groups items by a computed key in parallel                       |
| `GroupByWithLimit`   | Groups items by a computed key with limited concurrency          |
| `Partition`          | Splits a slice into two slices based on a predicate in parallel  |
| `PartitionWithLimit` | Splits a slice into two slices based on a predicate with a limit |

## Notes

- Prefer the function that most clearly expresses your intent.
- Use the `WithLimit` variants when you want to control how many goroutines run at once.
- Results are returned in the same order as the input slice where ordering applies.
- Avoid shared mutable state across parallel calls unless it is properly synchronized.

## Examples

Examples can be found in the [test suite](../../slicex/parallel/slice_test.go).