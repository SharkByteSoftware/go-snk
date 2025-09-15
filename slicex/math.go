package slicex

import (
	"cmp"

	"github.com/SharkByteSoftware/go-snk/adapt"
	"github.com/SharkByteSoftware/go-snk/constraints"
)

func Sum[T constraints.Numeric](slice []T) T {
	return SumBy(slice, adapt.ValueAdapter)
}

func SumBy[S ~[]T, T any, R constraints.Numeric](slice S, sumFunc func(item T) R) R {
	var sum R

	Apply(slice, func(item T) {
		sum += sumFunc(item)
	})

	return sum
}

func Product[S ~[]T, T constraints.Numeric](slice S) T {
	return ProductBy(slice, adapt.ValueAdapter)
}

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

func Mean[S ~[]T, T constraints.Numeric](slice S) T {
	return MeanBy(slice, adapt.ValueAdapter)
}

func MeanBy[S ~[]T, T any, R constraints.Numeric](slice S, valueFunc func(item T) R) R {
	count := R(len(slice))

	if count == 0 {
		return count
	}

	return SumBy(slice, valueFunc) / count
}

func Max[S ~[]T, T cmp.Ordered](slice S) T {
	return MaxBy(slice, func(a T, b T) bool { return a < b })
}

func MaxBy[S ~[]T, T any](slice S, maxFunc func(a T, b T) bool) T {
	var maxValue T

	Apply(slice, func(item T) {
		if maxFunc(maxValue, item) {
			maxValue = item
		}
	})

	return maxValue
}

func Min[S ~[]T, T cmp.Ordered](slice S) T {
	return MinBy(slice, func(a T, b T) bool { return a > b })
}

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
