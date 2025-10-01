package parallel

import (
	"sync"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

// Map transforms a slice to a slice of another type using a mapper function.
// The mapper function is called in parallel, and results are returned in order
// they appear in the slice.
func Map[S ~[]T, T any, R any](slice S, mapper func(item T) R) []R {
	result := make([]R, len(slice))

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(slice))

	slicex.ApplyWithIndex(slice, func(item T, idx int) {
		go func() {
			result[idx] = mapper(item)

			waitGroup.Done()
		}()
	})

	waitGroup.Wait()

	return result
}

// Apply applies a function to each item in the slice.  The apply function is called
// in parallel.
func Apply[S ~[]T, T any](slice S, apply func(item T)) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(slice))

	slicex.Apply(slice, func(item T) {
		go func() {
			apply(item)
			waitGroup.Done()
		}()
	})

	waitGroup.Wait()
}

// GroupBy returns a map of slices grouped by a key produced by a predicate function.
// The predicate is called in parallel, and the results are returned in the order they
// appear in the slice.
func GroupBy[S ~[]T, T any, R comparable](slice S, predicate func(item T) R) map[R]S {
	result := make(map[R]S, len(slice))

	keys := Map(slice, func(item T) R {
		return predicate(item)
	})

	slicex.ApplyWithIndex(slice, func(item T, idx int) {
		result[keys[idx]] = append(result[keys[idx]], item)
	})

	return result
}

// Partition splits a slice into two slices based on a predicate.  The predicate is called
// in parallel, and the results are returned in the order they appear in the slice.
func Partition[S ~[]T, T any](slice S, predicate func(item T) bool) (S, S) {
	result1 := make(S, 0)
	result2 := make(S, 0)

	result := Map(slice, func(item T) *S {
		return conditional.If(predicate(item), &result1, &result2)
	})

	slicex.ApplyWithIndex(slice, func(item T, idx int) {
		*result[idx] = append(*result[idx], item)
	})

	return result1, result2
}
