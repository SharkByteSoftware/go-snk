# Overview

`go-snk` is a Go utility library built around small, reusable, type-safe helpers.

It brings together a focused set of packages for common tasks such as:
- working with slices
- working with maps
- making HTTP requests
- expressing simple conditional logic
- using reusable container data structures

## Why go-snk?

The goal of `go-snk` is to reduce repetitive boilerplate while keeping everyday code easy to read.

Use it when you want:
- small helper functions instead of custom loops scattered throughout your codebase
- generic utilities that remain type-safe
- focused packages rather than a single large framework
- consistent patterns across common helper areas

## Package areas

### `slicex`
Helpers for slice operations such as filtering, mapping, reducing, grouping, deduplicating, partitioning, and basic numeric aggregation.

### `slicex/parallel`
Parallel versions of selected slice operations for independent work where concurrency can improve throughput.

### `mapx`
Helpers for common map operations such as extracting keys or values, filtering entries, transforming maps, and combining maps.

### `httpx`
Lightweight HTTP helpers that reduce client-side boilerplate and support typed response handling.

### `conditional`
Small helpers for concise conditional expressions and callback-style branching.

### `containers`
Reusable collection types including lists, sets, stacks, and queues.

## Design principles

The library aims to stay:
- small and focused
- easy to read
- type-safe
- practical for everyday use
- consistent across packages

## Recommended starting point

If you're new to the project, start with:
1. [`docs/examples.md`](examples.md)
2. the package document for the area you want to use
3. the package summary in the root README

## Notes

This project is intended as a utility toolkit rather than a full application framework. Most packages are intentionally narrow in scope so they can be adopted independently.