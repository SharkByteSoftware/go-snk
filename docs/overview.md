<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../img/logo.png" />
    </a>
</div>

# Overview

`go-snk` is a Go utility library built around small, reusable, type-safe helpers. The goal is to reduce repetitive 
boilerplate while keeping everyday code easy to read — small helper functions instead of custom loops, generic utilities that stay type-safe, and consistent patterns across common helper areas.

## Packages

### [`slicex`](slicex.md)
Helpers for slice operations such as filtering, mapping, reducing, grouping, deduplicating, partitioning, and basic 
numeric aggregation.

### [`slicex/parallel`](parallel/slicex.md)
Helpers for parallel slice operations where each item can be processed independently and concurrency can improve 
throughput.

### [`mapx`](mapx.md)
Helpers for common map operations such as extracting keys or values, filtering entries, transforming maps, and 
combining maps.

### [`httpx`](httpx.md)
Helpers for HTTP client code that reduce boilerplate and support typed response handling.

### [`conditional`](conditional.md)
Helpers for concise conditional expressions and callback-style branching.

### [`containers`](containers/README.md)
Reusable collection types including lists, sets, stacks, and queues.

## Design principles

The library aims to stay small and focused rather than grow into a full application framework. Most packages are 
intentionally narrow in scope so they can be adopted independently, and practical for everyday use rather than edge cases.