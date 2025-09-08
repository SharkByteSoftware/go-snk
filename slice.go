// Package sink provides various slice utilities.
package sink

func sliceFilterAdapter[T any](f func(item T) bool) func(T, int) bool {
	return func(item T, idx int) bool {
		return f(item)
	}
}

// Filter filters a slice using a filter function.
func Filter[T any, S ~[]T](slice S, filter func(item T) bool) []T {
	return FilterWithIndex(slice, sliceFilterAdapter(filter))
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

func sliceMapperAdapter[T any, R any](mapper func(T) R) func(T, int) R {
	return func(item T, index int) R {
		return mapper(item)
	}
}

// Map maps a slice to a slice of another type using a mapper function.
func Map[T, R any](slice []T, mapper func(item T) R) []R {
	return MapWithIndex(slice, sliceMapperAdapter(mapper))
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

// Unique returns a slice with all duplicate values removed.
func Unique[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	set := NewSet[T]()

	for _, value := range slice {
		if set.Contains(value) {
			continue
		}

		result = append(result, value)
		set.Add(value)
	}

	return result
}
