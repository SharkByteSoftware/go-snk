// Package slices provides various slice utilities.
package slices

import (
	"github.com/SharkByteSoftware/go-sink/sets"
)

// Filter filters a slice using a filter function.
func Filter[T any, S ~[]T](slice S, filter func(item T) bool) []T {
	return FilterWithIndex(slice, ItemIndexAdapter(filter))
}

// FilterWithIndex is like Filter, but it accepts a filter function that takes an index as well.
func FilterWithIndex[T any, S ~[]T](slice S, filter func(item T, index int) bool) []T {
	result := make(S, 0, len(slice))
	for index, value := range slice {
		if filter(value, index) {
			result = append(result, value)
		}
	}

	return result
}

// Map maps a slice to a slice of another type using a mapper function.
func Map[T, R any](slice []T, mapper func(item T) R) []R {
	return MapWithIndex(slice, ItemIndexAdapter(mapper))
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

func Bind[T, R any](slice []T, mapper func(item T) []R) []R {
	result := make([]R, 0, len(slice))
	return result
}

func Fold[T any, R any](slice []T, reducer func(acc R, item T) R) R {
	var acc R
	return acc
}

func Any[T any](slice []T, condition func(item T) bool) bool {
	return false
}

func All[T any, R comparable](slice []T) bool {
	return false
}

func Contains[T comparable](slice []T, item T) bool {
	return false
}

// Unique returns a slice with all duplicate values removed.
func Unique[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	set := sets.NewSet[T]()

	for _, value := range slice {
		if set.Contains(value) {
			continue
		}

		result = append(result, value)
		set.Add(value)
	}

	return result
}

func Apply[T any](slice []T, f func(item T)) {
	for _, value := range slice {
		f(value)
	}
}

func GroupBy[T, R comparable, S ~[]T](slice S, groupFunc func(item T) R) map[R][]T {
	// TODO: Implement
	return nil
}

func Reverse[T any, S ~[]T](slice S) S {
	// TODO: Implement
	return nil
}

func ToMap[T any, K comparable, V any](slice []T, keyFunc func(item T) K, valueFunc func(item T) V) map[K]V {
	// TODO: Implement
	return nil
}
