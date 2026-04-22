<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="img/logo.png" />
    </a>
</div>

# go-snk — the kitchen sink for clean Go projects

[![GitHub Tag](https://img.shields.io/github/v/tag/SharkByteSoftware/go-snk)](https://github.com/SharkByteSoftware/go-snk/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/SharkByteSoftware/go-snk.svg)](https://pkg.go.dev/github.com/SharkByteSoftware/go-snk)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/SharkByteSoftware/go-snk)](./go.mod)
[![Build Status](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml)
[![Lint Status](https://github.com/SharkByteSoftware/go-snk/actions/workflows/lint.yml/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/lint.yml)
[![Go report](https://goreportcard.com/badge/github.com/SharkByteSoftware/go-snk)](https://goreportcard.com/report/SharkByteSoftware/go-snk)
[![CodeQL](https://github.com/SharkByteSoftware/go-snk/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/github-code-scanning/codeql)
[![GitHub License](https://img.shields.io/github/license/SharkByteSoftware/go-snk)](./LICENSE)
[![Contributors](https://img.shields.io/github/contributors/SharkByteSoftware/go-snk)](https://github.com/SharkByteSoftware/go-snk/graphs/contributors)

`go-snk` is the library you pull into any Go project to clean it up and keep it clean. The name is short for
*kitchen sink* — not because it does everything, but because it collects the small, type-safe helpers you keep reaching
for across projects. No repeated loops, no one-off helpers scattered through your codebase, just consistent patterns
for the things you do all the time.

## Getting Started

### Installation

```sh
go get github.com/SharkByteSoftware/go-snk@latest
```

Import only the packages you need:

```go
package main

import "github.com/SharkByteSoftware/go-snk/slicex"
import "github.com/SharkByteSoftware/go-snk/mapx"
```

## Documentation

| Package           | Description                                                                                                                      | Docs                                   |
|-------------------|----------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| `slicex`          | Helpers for filtering, mapping, searching, grouping, zipping, windowing, rotating, and numeric operations on slices              | [slicex.md](docs/slicex.md)            |
| `slicex/parallel` | Parallel slice helpers for independent per-item work                                                                             | [slicex.md](docs/parallel/slicex.md)   |
| `mapx`            | Helpers for common map operations including partitioning, key mapping, counting by classifier, and merging maps                  | [mapx.md](docs/mapx.md)                |
| `httpx`           | Lightweight helpers for HTTP requests with less boilerplate                                                                      | [httpx.md](docs/httpx.md)              |
| `jsonx`           | Helpers for encoding and decoding JSON to and from readers, byte slices, and strings                                             | [jsonx.md](docs/jsonx.md)              |
| `conditional`     | Concise helpers for branching and value selection                                                                                | [conditional.md](docs/conditional.md)  |
| `stringx`         | Helpers for blank checks, fallback selection, truncation, wrapping, and padding strings                                          | [stringx.md](docs/stringx.md)          |
| `errorx`          | Helpers for intentional error suppression, initialization panics, multi-target error checks, and reducing multiple errors to one | [errorx.md](docs/errorx.md)            |
| `containers`      | Reusable data structures including lists, sets, stacks, and queues                                                               | [README.md](docs/containers/README.md) |
| `helpers`         | Generic utilities for pointer handling and zero value operations                                                                 | [helpers.md](docs/helpers.md)          |

For a full overview of the library, see [docs/overview.md](docs/overview.md).

## License

MIT License, see [LICENSE](LICENSE) for details
