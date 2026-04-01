<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="img/logo.png" />
    </a>
</div>

# Changelog

## v1.1.1

### Bug Fixes
- **`slicex.MaxBy`** — fixed incorrect results when all values are below the type's zero value (e.g. all-negative integers). `MaxBy` now seeds from `slice[0]` consistent with `MinBy`, rather than from the zero value of `T`
- **`containers/stacks.Stack`** — relaxed generic constraint from `comparable` to `any`, consistent with all other container types. Existing code is unaffected

### Improvements
- Updated lint settings

---

## v1.1.0

### New Packages

- **`jsonx`** — helpers for decoding JSON from common sources with less boilerplate: `Decode` (from `io.Reader`), `DecodeBytes`, `DecodeString`, with `WithStrictDecoding` and `WithUseNumber` options
- **`stringx`** — helpers for common string operations: blank checks (`IsBlank`), fallback selection (`Coalesce`), truncation (`Truncate`), wrapping (`Wrap`), and padding (`PadLeft`, `PadRight`)
- **`errorx`** — helpers for common error handling patterns: intentional suppression (`Ignore`), initialization-time panics (`Must`), and multi-target matching (`IsAny`)

### New Functions

#### `slicex`
- `Zip` — combines two slices into a slice of `Pair` values paired by index
- `Window` — returns overlapping sub-slices of a fixed size
- `Rotate` — returns a copy of a slice shifted left or right by n positions

#### `mapx`
- `Partition` — splits a map into two maps based on a predicate
- `CountBy` — returns a map of counts grouped by a classifier function
- `MapKeys` — returns a new map with each key transformed by a mapper function
- `Filter` — returns a map containing only entries satisfying a predicate
- `Apply` — runs a function on each map entry for side effects

### Improvements
- **`httpx`** — overhauled error handling with typed sentinel errors (`ErrTransport`, `ErrDecoding`, `ErrEncoding`, `ErrResponse`, `ErrOptions`) and corresponding typed error structs with `errors.Is` / `errors.As` support; expanded documentation and examples
- Performance improvements across `slicex` based on compiler output analysis, reducing heap allocations
- Expanded `mapx` documentation and examples
- Test linting added to CI via a dedicated golangci-lint configuration
- General linting pass across the codebase; re-enabled `revive`, `funlen`, and `thelper` linters
- Updated GitHub Actions versions

---

## v1.0.0

### Improvements
- Linting fixes and cleanup across the codebase in preparation for stable release

---

## v0.9.9

### Improvements
- Updated library description and blurbs across README and overview docs to better reflect intent

---

## v0.9.8

### New Features
- **`httpx`** — functional options pattern for request configuration: `WithHTTPClient`, `WithHeader`, `WithHeaders`, `WithTimeout`, `WithParam`, `WithParams`, `StrictDecoding`
- Added `StrictDecoding` config option to `httpx`

### Improvements
- Expanded `httpx` test coverage with helper assertion functions and timeout failure test
- Updated golangci-lint to v2.5 with `goinstall` mode
- Expanded `httpx` examples including error handling example

---

## v0.9.7

### Improvements
- Minor allocation reduction in `slicex`
- Fixed misspelling in documentation

---

## v0.9.6

### New Functions
- **`slicex`** — `Contains`

---

## v0.9.5

### Improvements
- Expanded test coverage and examples for `slicex` math functions: `SumBy`, `MaxBy`, `MinBy`, `ProductBy`, `MeanBy`

---

## v0.9.4

### New Features
- **`slicex/parallel`** — concurrency-limited variants for all parallel helpers: `MapWithLimit`, `ApplyWithLimit`, `GroupByWithLimit`, `PartitionWithLimit`
- Concurrency is now capped to `len(slice)` across all parallel helpers

---

## v0.9.3

### New Functions
- **`slicex`** — `FilterMap`, `FilterMapWithIndex`

### Improvements
- Added `update-pkg-go-dev` target to Makefile

---

## v0.9.2

### New Functions
- **`slicex`** — `FirstOr`, `FirstOrEmpty`, `UniqueBy`

---

## v0.9.1

### Bug Fixes
- **`slicex.FindBy`** — fixed issue with `comparable` constraint that prevented use with non-comparable types

---

## v0.9.0

### New Packages
- **`containers/queues`** — double-ended queue (`DQueue`) with `Enqueue`, `EnqueueFront`, `Dequeue`, `DequeueBack`, `Peek`, `PeekBack`, `IsEmpty`, `Size`, `Clear`, `Values`

### New Functions
- **`slicex`** — `Compact`

### New Features
- `Container` interface introduced across `lists`, `sets`, and `stacks` (`IsEmpty`, `Size`, `Values`, `Clear`)

### Improvements
- Added benchmark tests for `containers/queues`
- Expanded `slicex` examples
- Internal constraint and adapter packages moved to `internal`

---

## v0.1.0

### New Packages
- **`slicex/parallel`** — parallel versions of core slice helpers: `Map`, `Apply`, `GroupBy`, `Partition`, with results returned in input order

### Improvements
- Added examples for `slicex` math functions
- Added examples for `mapx` functions
- Added examples for `containers/sets`
- Fixed shadowed variable lint errors
- Updated golangci-lint config

---

## v0.0.3

### New Packages
- **`containers/stacks`** — LIFO stack backed by a linked list with `Push`, `Pop`, `Peek`, `IsEmpty`, `Size`, `Clear`, `Values`

### New Features
- Added golangci-lint workflow to CI
- Added godoc target to Makefile
- Added project logo

### Improvements
- Updated README and workflows
- Added list benchmark tests

---

## v0.0.2

### New Packages
- **`containers/lists`** — doubly linked list with `PushFront`, `PushBack`, `Prepend`, `Append`, `InsertBefore`, `InsertAfter`, `Remove`, `MoveToFront`, `MoveToBack`, `Front`, `Back`, `Len`, `IsEmpty`, `Values`

### New Functions
- **`mapx`** — `Apply`, `Combine`
- **`containers/sets`** — `IsEmpty`, `Equals`, `ApplyWithIndex`

### Improvements
- Restructured package layout; lists and sets moved under `containers`
- Renamed factory functions for consistency
- Added JSON serialization support to `sets`
- Badge and README improvements

---

## v0.0.1

### Initial Release

- **`slicex`** — `Filter`, `FilterWithIndex`, `Map`, `MapWithIndex`, `Bind`, `Reduce`, `Find`, `FindBy`, `FindOr`, `FindOrBy`, `Any`, `AnyBy`, `All`, `AllBy`, `Unique`, `UniqueMap`, `ToMap`, `GroupBy`, `Partition`, `Reverse`, `Apply`, `ApplyWithIndex`, `Sum`, `SumBy`, `Product`, `ProductBy`, `Mean`, `MeanBy`, `Min`, `MinBy`, `Max`, `MaxBy`
- **`mapx`** — `Keys`, `Values`, `UniqueValues`, `Contains`, `ValueOr`, `ToSlice`, `MapKeys`, `Invert`, `Filter`, `Partition`, `CountBy`
- **`containers/sets`** — `New`, `Add`, `Remove`, `Contains`, `Intersect`, `Union`, `Difference`, `SymmetricDifference`, `Subset`, `Equals`, `Clone`, `IsEmpty`, `Size`, `Clear`, `Values`, `Apply`
- **`conditional`** — `If`, `IfNotNil`, `IfCall`, `IfCallReturn`
- **`httpx`** — initial implementation of `Get`, `Post`, `Put`, `Patch`, `Delete`, `Head`, `Options` with raw response variants. Note: the request configuration API was significantly expanded in v0.9.8 with the functional options pattern (`WithHTTPClient`, `WithHeader`, `WithTimeout`, etc.) and error handling was overhauled in v1.1.0
- Benchmark tests for `slicex` and `containers/sets`
- CI workflow via GitHub Actions
- 