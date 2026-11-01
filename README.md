<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="img/logo.png" />
    </a>
</div>

# go-snk — the kitchen sink for clean Go projects

[![GitHub Tag](https://img.shields.io/github/v/tag/SharkByteSoftware/go-snk)](https://github.com/SharkByteSoftware/go-snk/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/SharkByteSoftware/go-snk.svg)](https://pkg.go.dev/github.com/SharkByteSoftware/go-snk)
[![Build Status](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/SharkByteSoftware/go-snk/graph/badge.svg?token=HGGS1ESDYP)](https://codecov.io/gh/SharkByteSoftware/go-snk)
[![Go report](https://goreportcard.com/badge/github.com/SharkByteSoftware/go-snk)](https://goreportcard.com/report/SharkByteSoftware/go-snk)
[![GitHub License](https://img.shields.io/github/license/SharkByteSoftware/go-snk)](./LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

`go-snk` is the library you pull into any Go project to clean it up and keep it clean. The name is short for
*kitchen sink* — not because it does everything, but because it collects the small, type-safe helpers you keep reaching
for across projects. No repeated loops, no one-off helpers scattered through your codebase, just consistent patterns
for the things you do all the time.

## Getting Started

### Installation
### Quick example

```go
type Order struct {
    ID       string
    Customer string
    Shipped  bool
}

pending    := slicex.Filter(orders, func(o Order) bool { return !o.Shipped })
ids        := slicex.Map(pending, func(o Order) string { return o.ID })
byCustomer := slicex.GroupBy(pending, func(o Order) string { return o.Customer })
```

Three lines instead of three hand-rolled loops, with each step's intent explicit.

```sh
go get github.com/SharkByteSoftware/go-snk@latest
```

Import only the packages you need:

```go
package main

import (
	"github.com/SharkByteSoftware/go-snk/mapx"
	"github.com/SharkByteSoftware/go-snk/slicex"
)
```

## Why go-snk?

Stdlib first: go-snk fills the gaps `slices`/`maps` leave out (`Map`, `Filter`, `GroupBy`…)
with consistent `S ~[]T` / `M ~map[K]V` constraints, order-preserving fail-fast parallel
variants, and backend helpers (`jsonx`, `httpx`, `httpxtest`).
Compared with `samber/lo`, it's narrower and more opinionated. → [Full comparison](docs/why-go-snk.md)

## Documentation

| Package           | Description                                                                                                                      | Docs                                   |
|-------------------|----------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| `slicex`          | Helpers for filtering, mapping, searching, grouping, zipping, windowing, rotating, and numeric operations on slices              | [slicex.md](docs/slicex.md)            |
| `slicex/parallel` | Parallel slice helpers for independent per-item work                                                                             | [slicex.md](docs/parallel/slicex.md)   |
| `mapx`            | Helpers for common map operations including partitioning, key mapping, counting by classifier, and merging maps                  | [mapx.md](docs/mapx.md)                |
| `httpx`           | Lightweight helpers for HTTP requests with less boilerplate                                                                      | [httpx.md](docs/httpx.md)              |
| `httpxtest`       | Builder for `httptest` servers with canned responses, per-route handlers, and response shaping                                   | [httpxtest.md](docs/httpxtest.md)      |
| `jsonx`           | Helpers for encoding and decoding JSON to and from readers, byte slices, and strings                                             | [jsonx.md](docs/jsonx.md)              |
| `conditional`     | Concise helpers for branching and value selection                                                                                | [conditional.md](docs/conditional.md)  |
| `stringx`         | Helpers for blank checks, fallback selection, truncation, wrapping, and padding strings                                          | [stringx.md](docs/stringx.md)          |
| `errorx`          | Helpers for intentional error suppression, initialization panics, multi-target error checks, and reducing multiple errors to one | [errorx.md](docs/errorx.md)            |
| `containers`      | Reusable data structures including lists, sets, stacks, and queues                                                               | [README.md](docs/containers/README.md) |
| `helpers`         | Generic utilities for pointer handling and zero value operations                                                                 | [helpers.md](docs/helpers.md)          |

For a full overview of the library, see [docs/overview.md](docs/overview.md).

## License

MIT License, see [LICENSE](LICENSE) for details
