<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# slicex/parallel

`parallel` provides parallel versions of selected `slicex` helpers for workloads where each item can be processed independently.

It is useful when you want to:
- map slice values using concurrent work
- apply a function to each item in parallel
- group values using parallel predicate evaluation
- partition values using parallel predicate evaluation
- limit concurrency for better control over resource usage

## Overview

This package is intended for situations where the work done for each slice item is independent and concurrency can improve throughput.

It follows the same general patterns as `slicex`, but evaluates the per-item work in parallel instead of sequentially.

## Common capabilities

- parallel mapping
- parallel application of functions
- concurrency-limited mapping
- concurrency-limited application
- grouping by predicate in parallel
- concurrency-limited grouping
- partitioning by predicate in parallel
- concurrency-limited partitioning

## When to use it

Use `slicex/parallel` when:
- each item can be processed independently
- the work per item is expensive enough to benefit from concurrency
- you want to keep results ordered
- you want to cap concurrency to avoid excessive goroutine usage

## API reference

| Function             | Purpose                                                             |
|----------------------|---------------------------------------------------------------------|
| `Map`                | Transforms a slice using a mapper function in parallel              |
| `MapWithLimit`       | Transforms a slice using a mapper function with limited concurrency |
| `Apply`              | Applies a function to each item in parallel                         |
| `ApplyWithLimit`     | Applies a function to each item with limited concurrency            |
| `GroupBy`            | Groups items by a computed key in parallel                          |
| `GroupByWithLimit`   | Groups items by a computed key with limited concurrency             |
| `Partition`          | Splits a slice into two slices based on a predicate in parallel     |
| `PartitionWithLimit` | Splits a slice into two slices based on a predicate with a limit    |

## Notes

- Results are returned in the same order as the input slice where ordering applies.
- Use the `WithLimit` variants when you want to control how many goroutines work at once.
- Parallel helpers are best suited to independent tasks; avoid a shared mutable state unless it is properly synchronized.

## Examples
