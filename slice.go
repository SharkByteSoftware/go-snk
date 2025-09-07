// Package sink provides various slice utilities.
package sink

type SliceFilter[T any] func(T) bool
type SliceFilterIndex[T any] func(T, int) bool

func SliceFilterIndexAdapter[T any](f func(T) bool) SliceFilterIndex[T] {
	return func(item T, index int) bool {
		return f(item)
	}
}

// Filter filters a slice using a filter function.
func Filter[T any, S ~[]T](slice S, filter SliceFilter[T]) S {
	return FilterI(slice, SliceFilterIndexAdapter(filter))
}

// FilterI is like Filter, but it accepts a filter function that takes an index as well.
func FilterI[T any, S ~[]T](slice S, filter SliceFilterIndex[T]) S {
	result := make(S, 0, len(slice))
	for index, value := range slice {
		if filter(value, index) {
			result = append(result, value)
		}
	}

	return result
}
