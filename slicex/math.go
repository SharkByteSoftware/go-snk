package slicex

import (
	"cmp"

	"github.com/SharkByteSoftware/go-snk/internal/adapt"
	"github.com/SharkByteSoftware/go-snk/internal/constraint"
)

// Sum returns the sum of all the values of the slice.
func Sum[T constraint.Numeric](slice []T) T {
	return SumBy(slice, adapt.ValueAdapter)
}

// SumBy returns the sum of all the values of the slice as determined by the provided sum function.
func SumBy[S ~[]T, T any, R constraint.Numeric](slice S, sumFunc func(item T) R) R {
	var sum R

	Apply(slice, func(item T) {
		sum += sumFunc(item)
	})

	return sum
}

// Product returns the product of the values of the slice.
func Product[S ~[]T, T constraint.Numeric](slice S) T {
	return ProductBy(slice, adapt.ValueAdapter)
}

// ProductBy returns the product of the values in the slice as determined by the provided product function.
func ProductBy[S ~[]T, T any, R constraint.Numeric](slice S, productFunc func(item T) R) R {
	var product R = 1

	if len(slice) == 0 {
		return product
	}

	Apply(slice, func(item T) {
		product *= productFunc(item)
	})

	return product
}

// Mean returns the mean of the values of the slice.
func Mean[S ~[]T, T constraint.Numeric](slice S) T {
	return MeanBy(slice, adapt.ValueAdapter)
}

// MeanBy returns the mean of the values of the slice as determined by the provided value function.
func MeanBy[S ~[]T, T any, R constraint.Numeric](slice S, valueFunc func(item T) R) R {
	count := R(len(slice))

	if count == 0 {
		return count
	}

	return SumBy(slice, valueFunc) / count
}

// Max provides the maximum value of the slice.
func Max[S ~[]T, T cmp.Ordered](slice S) T {
	return MaxBy(slice, func(a T, b T) bool { return a < b })
}

// MaxBy returns the maximum value of the slice as determined by the provided maximum function.
func MaxBy[S ~[]T, T any](slice S, maxFunc func(a T, b T) bool) T {
	var maxValue T

	Apply(slice, func(item T) {
		if maxFunc(maxValue, item) {
			maxValue = item
		}
	})

	return maxValue
}

// Min returns the minimum value of the slice.
func Min[S ~[]T, T cmp.Ordered](slice S) T {
	return MinBy(slice, func(a T, b T) bool { return a > b })
}

// MinBy returns the minimum value of the slice as determined by the provided minimum function.
func MinBy[S ~[]T, T any](slice S, minFunc func(a T, b T) bool) T {
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
