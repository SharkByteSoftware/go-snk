// Package slices provides various slice utilities.
package slices

import (
	"github.com/SharkByteSoftware/go-snk/adapt"
	"github.com/SharkByteSoftware/go-snk/conditionals"
	"github.com/SharkByteSoftware/go-snk/ds"
)

// Filter filters a slice using a predicate function.
func Filter[T any, S ~[]T](slice S, predicate func(item T) bool) []T {
	return FilterWithIndex(slice, adapt.ItemIndexAdapter(predicate))
}

// FilterWithIndex is like Filter, but it accepts a predicate function that takes an index as well.
func FilterWithIndex[T any, S ~[]T](slice S, predicate func(item T, index int) bool) []T {
	result := make(S, 0, len(slice))
	for index, value := range slice {
		if predicate(value, index) {
			result = append(result, value)
		}
	}

	return result
}

// Map maps a slice to a slice of another type using a mapper function.
func Map[T, R any](slice []T, mapper func(item T) R) []R {
	return MapWithIndex(slice, adapt.ItemIndexAdapter(mapper))
}

// MapWithIndex is like Map, but it accepts a mapper function that takes an index as well.
func MapWithIndex[T, R any](slice []T, mapper func(item T, idx int) R) []R {
	result := make([]R, len(slice))

	for idx, value := range slice {
		result[idx] = mapper(value, idx)
	}

	return result
}

// UniqueMap maps a slice to a slice of another type using a mapper function and removes duplicate values.
func UniqueMap[T, R comparable](slice []T, mapper func(item T) R) []R {
	return Unique(Map(slice, mapper))
}

// Bind transforms and flattens a slice from one type to another using a mapper
// function.
func Bind[T any, S ~[]T, R any, RS ~[]R](slice S, mapper func(item T) RS) RS {
	result := make([]R, 0, len(slice))

	Apply(slice, func(item T) {
		result = append(result, mapper(item)...)
	})

	return result
}

// Fold reduces a slice to a value which is the accumulated result of calling an accumulate func
// for each item in the slice where each successive call is supplied by the return value of
// the previous call.
func Fold[T any, R any](slice []T, accumulator func(agg R, item T) R, initial R) R {
	Apply(slice, func(item T) {
		initial = accumulator(initial, item)
	})

	return initial
}

// Find returns the first item in the slice that is equal to the given candidate.
func Find[T comparable](slice []T, candidate T) (T, bool) {
	return FindBy(slice, adapt.ItemEqualsAdapter(candidate))
}


// FindBy returns the first item in the slice that satisfies the predicate.
func FindBy[T comparable](slice []T, predicate func(item T) bool) (T, bool) {
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
func FindOr[T comparable](slice []T, candidate T, fallback T) T {
	item, found := FindBy(slice, adapt.ItemEqualsAdapter(candidate))
	return conditionals.If(found, item, fallback)
}


func FindByOr[T comparable](slice []T, predicate func(item T) bool, fallback T) T {
	item, found := FindBy(slice, predicate)
	return conditionals.If(found, item, fallback)
}

func Any[T comparable](slice []T, candidate T) bool {
	return AnyBy(slice, adapt.ItemEqualsAdapter(candidate))
}

// AnyBy returns true if any item in the slice satisfies the predicate.
func AnyBy[T comparable](slice []T, predicate func(item T) bool) bool {
	_, found := FindBy(slice, predicate)
	return found
}

// All returns true if all items in the slice are equal to the given candidate.
func All[T comparable](slice []T, candidate T) bool {
	found := Filter(slice, func(item T) bool { return item == candidate })
	return len(found) == len(slice)
}

// Unique returns a slice with all duplicate values removed.
func Unique[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	set := ds.NewSet[T]()

	for _, value := range slice {
		if set.Contains(value) {
			continue
		}

		result = append(result, value)
		set.Add(value)
	}

	return result
}

// Apply applies a function to each item in the slice.
func Apply[T any](slice []T, predicate func(item T)) {
	for _, value := range slice {
		predicate(value)
	}
}

func GroupBy[T, R comparable, S ~[]T](slice S, groupFunc func(item T) R) map[R][]T {
	// TODO: Implement
	return nil
}

// Reverse reverses a slice.
func Reverse[T any, S ~[]T](slice S) S {
	result := make([]T, len(slice))
	sliceLen := len(slice)
	mid := sliceLen / 2

	for i := range mid {
		j := sliceLen - 1 - i
		result[i], result[j] = slice[j], slice[i]
	}

	if sliceLen%2 != 0 {
		result[mid] = slice[mid]
	}

	return result
}

// IndexOf returns the index of the first item in the slice that is equal to the given candidate.
func IndexOf[T comparable](slice []T, candidate T) (int, bool) {
	return IndexOfBy(slice, adapt.ItemEqualsAdapter(candidate))
}

// IndexOfBy returns the index of the first item in the slice that satisfies the predicate.
func IndexOfBy[T comparable](slice []T, predicate func(item T) bool) (int, bool) {
	for idx, value := range slice {
		if predicate(value) {
			return idx, true
		}
	}

	return -1, false
}

// ToMap converts a slice to a map using the given key function to determine the map key.
func ToMap[T any, K comparable](slice []T, predicate func(item T) K) map[K]T {
	result := make(map[K]T, len(slice))

	Apply(slice, func(item T) {
		result[predicate(item)] = item
	})

	return result
}
