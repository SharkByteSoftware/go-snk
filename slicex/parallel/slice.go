// Package parallel provides versions of slicex functions that execute in parallel.
package parallel

import (
	"github.com/SharkByteSoftware/go-snk/slicex"
)

const defaultMaxConcurrency = 1000

// Map transforms a slice to a slice of another type using a mapper function.
// The mapper function is called in parallel, and results are returned in order
// they appear in the slice.
func Map[S ~[]T, T any, R any](slice S, mapper func(item T) R) []R {
	return MapWithLimit(slice, mapper, defaultMaxConcurrency)
}

// MapWithLimit transforms a slice to a slice of another type using a mapper function
// with a max concurrency. The mapper function is called in parallel, and results are
// returned in order they appear in the slice.
func MapWithLimit[S ~[]T, T any, R any](slice S, mapper func(item T) R, concurrency int) []R {
	result := make([]R, len(slice))

	runIndexed(slice, concurrency, func(idx int, item T) {
		result[idx] = mapper(item)
	})

	return result
}

// Apply applies a function to each item in the slice. The apply function is called
// in parallel with at most 1000 concurrent calls.
// If the function panics, the panic is re-raised on the calling goroutine.
func Apply[S ~[]T, T any](slice S, apply func(item T)) {
	ApplyWithLimit(slice, apply, defaultMaxConcurrency)
}

// ApplyWithLimit applies a function to each item in the slice. The apply function is called
// in parallel with a max concurrency.
// If the function panics, the panic is re-raised on the calling goroutine.
func ApplyWithLimit[S ~[]T, T any](slice S, apply func(item T), concurrency int) {
	runIndexed(slice, concurrency, func(_ int, item T) {
		apply(item)
	})
}

// GroupBy returns a map of slices grouped by a key produced by a predicate function.
// The predicate is called in parallel, and the results are returned in the order they
// appear in the slice.
func GroupBy[S ~[]T, T any, R comparable](slice S, predicate func(item T) R) map[R]S {
	return groupByKeys(slice, Map(slice, predicate))
}

// GroupByWithLimit returns a map of slices grouped by a key produced by a predicate function.
// The predicate is called in parallel with a max concurrency, and the results are returned in the order they
// appear in the slice.
func GroupByWithLimit[S ~[]T, T any, R comparable](slice S, predicate func(item T) R, concurrency int) map[R]S {
	return groupByKeys(slice, MapWithLimit(slice, predicate, concurrency))
}

// Partition splits a slice into two slices based on a predicate. The predicate is called
// in parallel, and the results are returned in the order they appear in the slice.
func Partition[S ~[]T, T any](slice S, predicate func(item T) bool) (S, S) {
	return splitByMatch(slice, Map(slice, predicate))
}

// PartitionWithLimit splits a slice into two slices based on a predicate. The predicate is called
// in parallel with a max concurrency, and the results are returned in the order they appear in the slice.
func PartitionWithLimit[S ~[]T, T any](slice S, predicate func(item T) bool, concurrency int) (S, S) {
	return splitByMatch(slice, MapWithLimit(slice, predicate, concurrency))
}

func groupByKeys[S ~[]T, T any, R comparable](slice S, keys []R) map[R]S {
	result := make(map[R]S, len(slice))

	slicex.ApplyWithIndex(slice, func(item T, idx int) {
		result[keys[idx]] = append(result[keys[idx]], item)
	})

	return result
}

// splitByMatch sequentially routes items by the precomputed predicate results;
// the parallelism lives in the predicate evaluation, which is the expensive part.
func splitByMatch[S ~[]T, T any](slice S, matches []bool) (S, S) {
	const halfDivisor = 2

	half := len(slice) / halfDivisor
	matched := make(S, 0, half)
	unmatched := make(S, 0, half)

	for i, item := range slice {
		if matches[i] {
			matched = append(matched, item)
			continue
		}

		unmatched = append(unmatched, item)
	}

	return matched, unmatched
}
