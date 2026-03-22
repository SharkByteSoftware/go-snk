<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="img/logo.png" />
    </a>
</div>

# go-snk - Slices, Maps, Data Structures, Channels, and more...

[![GitHub Tag](https://img.shields.io/github/v/tag/SharkByteSoftware/go-snk)](https://github.com/SharkByteSoftware/go-snk/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/SharkByteSoftware/go-snk.svg)](https://pkg.go.dev/github.com/SharkByteSoftware/go-snk)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/SharkByteSoftware/go-snk)](./go.mod)
[![Build Status](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml)
[![Lint Status](https://github.com/SharkByteSoftware/go-snk/actions/workflows/lint.yml/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/lint.yml)
[![Go report](https://goreportcard.com/badge/github.com/SharkByteSoftware/go-snk)](https://goreportcard.com/report/SharkByteSoftware/go-snk)
[![GitHub License](https://img.shields.io/github/license/SharkByteSoftware/go-snk)](./LICENSE)
[![Contributors](https://img.shields.io/github/contributors/SharkByteSoftware/go-snk)](https://github.com/SharkByteSoftware/go-snk/graphs/contributors)

`go-snk` is a collection of small, type-safe Go utilities for slices, maps, HTTP requests, conditional logic, and reusable containers.

## Getting Started

### Installation

```sh
go get github.com/SharkByteSoftware/go-snk@latest
```
Import the packages you need:

```go
package main

import "github.com/SharkByteSoftware/go-snk/slicex"
import "github.com/SharkByteSoftware/go-snk/mapx"
```


### Documentation

- [Overview](docs/overview.md)
- [slicex](docs/slicex.md)
- [mapx](docs/mapx.md)
- [httpx](docs/httpx.md)
- [conditional](docs/conditional.md)
- [containers](docs/containers/README.md)
- [Roadmap](docs/roadmap.md)

## Package Summary

- `slicex` — helpers for filtering, mapping, searching, grouping, and numeric operations on slices
- `slicex/parallel` — parallel slice helpers for independent per-item work
- `mapx` — generic helpers for common map operations
- `httpx` — lightweight helpers for making HTTP requests with less boilerplate
- `conditional` — concise helpers for branching and value selection
- `containers` — reusable data structures including lists, sets, stacks, and queues

## License

MIT License, see [LICENSE](LICENSE) for details
