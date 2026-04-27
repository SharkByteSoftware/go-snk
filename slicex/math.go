package slicex

import (
	"cmp"

	"github.com/SharkByteSoftware/go-snk/helpers"
	"github.com/SharkByteSoftware/go-snk/internal/constraints"
)

// Sum returns the sum of all elements in the slice.
func Sum[S ~[]T, T constraints.Numeric](slice S) T {
	return SumBy(slice, func(item T) T { return item })
}

// SumBy returns the sum of values produced by applying the value selector to each element.
func SumBy[S ~[]T, T any, R constraints.Numeric](slice S, sumFunc func(item T) R) R {
	var sum R

	Apply(slice, func(item T) {
		sum += sumFunc(item)
	})

	return sum
}

// Product returns the product of all elements in the slice.
// Returns 1 for an empty slice.
func Product[S ~[]T, T constraints.Numeric](slice S) T {
	return ProductBy(slice, func(item T) T { return item })
}

// ProductBy returns the product of values produced by applying the value selector to each element.
// Returns 1 for an empty slice.
func ProductBy[S ~[]T, T any, R constraints.Numeric](slice S, productFunc func(item T) R) R {
	var product R = 1

	if len(slice) == 0 {
		return product
	}

	Apply(slice, func(item T) {
		product *= productFunc(item)
	})

	return product
}

// Mean returns the arithmetic mean of all elements in the slice.
// Returns the zero value of T for an empty slice.
func Mean[S ~[]T, T constraints.Numeric](slice S) T {
	return MeanBy(slice, func(item T) T { return item })
}

// MeanBy returns the arithmetic mean of values produced by applying the value selector to each element.
// Returns the zero value of R for an empty slice.
func MeanBy[S ~[]T, T any, R constraints.Numeric](slice S, valueFunc func(item T) R) R {
	count := R(len(slice))

	if count == 0 {
		return count
	}

	return SumBy(slice, valueFunc) / count
}

// Max returns the maximum element in the slice using natural ordering.
// Returns the zero value of T for an empty slice.
func Max[S ~[]T, T cmp.Ordered](slice S) T {
	return MaxBy(slice, func(a, b T) bool { return a < b })
}

// MaxBy returns the maximum element in the slice as determined by the less function.
// less(a, b) should return true when a should be considered less than b.
// Returns the zero value of T for an empty slice.
func MaxBy[S ~[]T, T any](slice S, maxFunc func(a, b T) bool) T {
	if len(slice) == 0 {
		return helpers.Empty[T]()
	}

	maxValue := slice[0]
	Apply(slice[1:], func(item T) {
		if maxFunc(maxValue, item) {
			maxValue = item
		}
	})

	return maxValue
}

// Min returns the minimum element in the slice using natural ordering.
// Returns the zero value of T for an empty slice.
func Min[S ~[]T, T cmp.Ordered](slice S) T {
	return MinBy(slice, func(a, b T) bool { return a > b })
}

// MinBy returns the minimum element in the slice as determined by the less function.
// less(a, b) should return true when a should be considered less than b.
// Returns the zero value of T for an empty slice.
func MinBy[S ~[]T, T any](slice S, minFunc func(a, b T) bool) T {
	var minValue T

	if len(slice) == 0 {
		return minValue
	}

	minValue = slice[0]

	Apply(slice, func(item T) {
		if minFunc(minValue, item) {
			minValue = item
		}
	})

	return minValue
}
