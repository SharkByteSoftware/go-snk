// Package slicex provides various slice utilities.
package slicex

import (
	"slices"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/containers/sets"
	"github.com/SharkByteSoftware/go-snk/helpers"
	"github.com/SharkByteSoftware/go-snk/internal/adapt"
)

// Filter filters a slice using a predicate function.
func Filter[S ~[]T, T any](slice S, predicate func(item T) bool) []T {
	return FilterWithIndex(slice, adapt.ItemIndexAdapter(predicate))
}

// FilterWithIndex is like Filter, but it accepts a predicate function that takes an index as well.
func FilterWithIndex[S ~[]T, T any](slice S, predicate func(item T, index int) bool) []T {
	result := make(S, 0, len(slice))

	ApplyWithIndex(slice, func(item T, index int) {
		if predicate(item, index) {
			result = append(result, item)
		}
	})

	return result
}

// Map transforms a slice to a slice of another type using a mapper function.
func Map[S ~[]T, T any, R any](slice S, mapper func(item T) R) []R {
	return MapWithIndex(slice, adapt.ItemIndexAdapter(mapper))
}

// MapWithIndex is like Map, but it accepts a mapper function that takes an index as well.
func MapWithIndex[S ~[]T, T any, R any](slice S, mapper func(item T, idx int) R) []R {
	result := make([]R, len(slice))

	ApplyWithIndex(slice, func(item T, idx int) {
		result[idx] = mapper(item, idx)
	})

	return result
}

// UniqueMap maps a slice to a slice of another type using a mapper function and removes duplicate values.
func UniqueMap[S ~[]T, T any, R comparable](slice S, mapper func(item T) R) []R {
	return Unique(Map(slice, mapper))
}

// Bind transforms and flattens a slice from one type to another using a mapper
// function. Function should return a slice or `nil`, if `nil` is returned then no
// value is added to the final result.
func Bind[S ~[]T, T any, R any, RS ~[]R](slice S, mapper func(item T) RS) RS {
	result := make([]R, 0, len(slice))

	Apply(slice, func(item T) {
		result = append(result, mapper(item)...)
	})

	return result
}

// Reduce transforms and flattens a slice to another type.
func Reduce[S ~[]T, T any, R any](slice S, accumulator func(agg R, item T) R, initial R) R {
	Apply(slice, func(item T) {
		initial = accumulator(initial, item)
	})

	return initial
}

// Find returns the first item in the slice that is equal to the given candidate.
func Find[S ~[]T, T comparable](slice S, candidate T) (T, bool) {
	return FindBy(slice, adapt.ItemEqualsAdapter(candidate))
}

// FindBy returns the first item in the slice that satisfies the predicate.
func FindBy[S ~[]T, T any](slice S, predicate func(item T) bool) (T, bool) {
	for _, value := range slice {
		if predicate(value) {
			return value, true
		}
	}

	var result T

	return result, false
}

// FindOr returns the first item in the slice that is equal to the given candidate,
// or the fallback value if not found.
func FindOr[S ~[]T, T comparable](slice S, candidate T, fallback T) T {
	return FindOrBy(slice, adapt.ItemEqualsAdapter(candidate), fallback)
}

// FindOrBy returns the first item in the slice that satisfies the predicate,
// or the fallback value if not found.
func FindOrBy[S ~[]T, T any](slice S, predicate func(item T) bool, fallback T) T {
	item, found := FindBy(slice, predicate)
	return conditional.If(found, item, fallback)
}

// Any returns true if any item in the slice satisfies the predicate.
func Any[S ~[]T, T comparable](slice S, candidate T) bool {
	return AnyBy(slice, adapt.ItemEqualsAdapter(candidate))
}

// AnyBy returns true if any item in the slice satisfies the predicate.
func AnyBy[S ~[]T, T any](slice S, predicate func(item T) bool) bool {
	_, found := FindBy(slice, predicate)
	return found
}

// All returns true if all items in the slice are equal to the given candidate.
func All[S ~[]T, T comparable](slice S, candidate T) bool {
	_, found := FindBy(slice, func(item T) bool {
		return candidate != item
	})

	return !found
}

// AllBy returns true if all items in the slice satisfy the predicate.
func AllBy[S ~[]T, T any](slice S, predicate func(item T) bool) bool {
	// TODO: implement
	panic("not implemented")
}

// Unique returns a slice with all duplicate values removed.
func Unique[S ~[]T, T comparable](slice S) []T {
	result := make([]T, 0, len(slice))
	set := sets.New[T]()

	Apply(slice, func(item T) {
		if !set.Contains(item) {
			set.Add(item)
			result = append(result, item)
		}
	})

	return result
}

// Apply applies a function to each item in the slice.
func Apply[S ~[]T, T any](slice S, apply func(item T)) {
	ApplyWithIndex(slice, func(item T, _ int) { apply(item) })
}

// ApplyWithIndex applies a function to each item in the slice and provides the index of the item.
func ApplyWithIndex[S ~[]T, T any](slice S, apply func(item T, index int)) {
	for idx, value := range slice {
		apply(value, idx)
	}
}

// Reverse returns a slice with the revers of the slice.
func Reverse[S ~[]T, T any](slice S) S {
	result := slices.Clone(slice)
	slices.Reverse(result)

	return result
}

// Compact returns a slice with all the non-zero items.
func Compact[S ~[]T, T comparable](slice S) S {
	return Filter(slice, func(item T) bool {
		return !helpers.IsEmpty(item)
	})
}

// ToMap converts a slice to a map using the predicate to determine the map key.
func ToMap[S ~[]T, T any, K comparable](slice S, predicate func(item T) K) map[K]T {
	result := make(map[K]T, len(slice))

	Apply(slice, func(item T) {
		result[predicate(item)] = item
	})

	return result
}

// GroupBy returns a map of slices grouped by a key produced by a predicate function.
func GroupBy[S ~[]T, T any, R comparable](slice S, predicate func(item T) R) map[R][]T {
	result := make(map[R][]T, len(slice))

	Apply(slice, func(item T) {
		result[predicate(item)] = append(result[predicate(item)], item)
	})

	return result
}

// Partition splits a slice into two slices based on a predicate.
func Partition[S ~[]T, T any](slice S, predicate func(item T) bool) (S, S) {
	part1 := make(S, 0)
	part2 := make(S, 0)

	Apply(slice, func(item T) {
		conditional.IfCall(predicate(item),
			func() { part1 = append(part1, item) },
			func() { part2 = append(part2, item) },
		)
	})

	return part1, part2
}
