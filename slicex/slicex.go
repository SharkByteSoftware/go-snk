// Package slicex provides various slice utilities.
package slicex

import (
	"slices"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/containers/sets"
	"github.com/SharkByteSoftware/go-snk/helpers"
)

// FirstOr returns the first item in the slice or a fallback value
// if the slice is empty.
func FirstOr[T any](slice []T, fallback T) T {
	if len(slice) == 0 {
		return fallback
	}

	return slice[0]
}

// FirstOrEmpty returns the first item in the slice or the empty value if
// the slice is empty.
func FirstOrEmpty[T any](slice []T) T {
	return FirstOr(slice, helpers.Empty[T]())
}

// Filter filters a slice using a predicate function.
func Filter[S ~[]T, T any](slice S, predicate func(item T) bool) []T {
	return FilterWithIndex(slice, func(item T, _ int) bool { return predicate(item) })
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
	return MapWithIndex(slice, func(item T, idx int) R { return mapper(item) })
}

// MapWithIndex is like Map, but it accepts a mapper function that takes an index as well.
func MapWithIndex[S ~[]T, T any, R any](slice S, mapper func(item T, idx int) R) []R {
	result := make([]R, len(slice))

	ApplyWithIndex(slice, func(item T, idx int) {
		result[idx] = mapper(item, idx)
	})

	return result
}

// FilterMap filters and transforms a slice to a slice of another type using a mapper function.
func FilterMap[S ~[]T, T any, R any](slice S, mapper func(item T) (R, bool)) []R {
	return FilterMapWithIndex(slice, func(item T, _ int) (R, bool) {
		return mapper(item)
	})
}

// FilterMapWithIndex filters and transforms a slice to a slice of another type using a mapper function.
func FilterMapWithIndex[S ~[]T, T any, R any](slice S, mapper func(item T, index int) (R, bool)) []R {
	result := make([]R, 0, len(slice))

	ApplyWithIndex(slice, func(item T, idx int) {
		if value, ok := mapper(item, idx); ok {
			result = append(result, value)
		}
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
	return FindBy(slice, func(item T) bool { return item == candidate })
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
	return FindOrBy(slice, func(item T) bool { return item == candidate }, fallback)
}

// FindOrBy returns the first item in the slice that satisfies the predicate,
// or the fallback value if not found.
func FindOrBy[S ~[]T, T any](slice S, predicate func(item T) bool, fallback T) T {
	item, found := FindBy(slice, predicate)
	return conditional.If(found, item, fallback)
}

// Contains returns true if the slice contains the given candidate.
func Contains[S ~[]T, T comparable](slice S, candidate T) bool {
	_, found := Find(slice, candidate)
	return found
}

// Any returns true if any item in the slice satisfies the predicate.
func Any[S ~[]T, T comparable](slice S, candidate T) bool {
	return AnyBy(slice, func(item T) bool { return item == candidate })
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
	_, found := FindBy(slice, func(item T) bool { return !predicate(item) })

	return !found
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

// UniqueBy returns a slice with unique values determined by a predicate function.
func UniqueBy[S ~[]T, T any, R comparable](slice S, predicate func(item T) R) []T {
	result := make([]T, 0, len(slice))
	set := sets.New[R]()

	Apply(slice, func(item T) {
		key := predicate(item)
		if !set.Contains(key) {
			set.Add(key)

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

// Reverse returns a slice with the reverse of the slice.
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

// GroupBy returns a map of slices grouped by a key produced by a key selector function.
func GroupBy[S ~[]T, T any, R comparable](slice S, predicate func(item T) R) map[R][]T {
	result := make(map[R][]T, len(slice))

	Apply(slice, func(item T) {
		key := predicate(item)
		result[key] = append(result[key], item)
	})

	return result
}

// Partition splits a slice into two slices based on a predicate.
func Partition[S ~[]T, T any](slice S, predicate func(item T) bool) (S, S) {
	half := len(slice) / 2
	part1 := make(S, 0, half)
	part2 := make(S, 0, half)

	Apply(slice, func(item T) {
		if predicate(item) {
			part1 = append(part1, item)
			return
		}

		part2 = append(part2, item)
	})

	return part1, part2
}

// Intersect returns a slice with the intersection of the two slices.
func Intersect[S ~[]T, T comparable](slice S, other S) S {
	return sets.New[T](slice...).
		Intersect(sets.New[T](other...)).
		Values()
}

// Union returns a slice with the union of the two slices.
func Union[S ~[]T, T comparable](slice S, other S) S {
	return sets.New[T](slice...).
		Union(sets.New[T](other...)).
		Values()
}

// Difference returns a slice with the difference of the two slices.
func Difference[S ~[]T, T comparable](slice S, other S) S {
	return sets.New[T](slice...).
		Difference(sets.New[T](other...)).
		Values()
}

// Zip combines two slices into a slice of Pairs, pairing elements by position.
// The result length is equal to the shorter of the two input slices.
func Zip[A any, B any](left []A, right []B) []Pair[A, B] {
	size := min(len(left), len(right))
	result := make([]Pair[A, B], size)

	for i := range size {
		result[i] = Pair[A, B]{Left: left[i], Right: right[i]}
	}

	return result
}

// Window returns a slice of overlapping sub-slices of the given size,
// advancing one position at a time. If size is less than 1 or greater
// than the length of the slice, an empty slice is returned.
func Window[S ~[]T, T any](slice S, size int) []S {
	if size < 1 || size > len(slice) {
		return []S{}
	}

	result := make([]S, len(slice)-size+1)

	for i := range result {
		result[i] = slice[i : i+size]
	}

	return result
}

// Rotate returns a copy of the slice with elements shifted left by n positions.
// Elements shifted off the front are wrapped to the back.
// A negative n shifts right instead. If the slice is empty, it is returned as-is.
func Rotate[S ~[]T, T any](slice S, slen int) S {
	if len(slice) == 0 {
		return slices.Clone(slice)
	}

	slen %= len(slice)
	if slen < 0 {
		slen += len(slice)
	}

	result := make(S, len(slice))
	copy(result, slice[slen:])
	copy(result[len(slice)-slen:], slice[:slen])

	return result
}
