package slices

import (
	"cmp"

	"github.com/SharkByteSoftware/go-snk/constraints"
)

func Sum[T constraints.Numeric](slice []T) T {
	return SumBy(slice, ValueAdapter[T]())
}

func SumBy[T any, R constraints.Numeric](slice []T, sumFunc func(item T) R) R {
	var sum R

	Apply(slice, func(item T) {
		sum += sumFunc(item)
	})

	return sum
}

func Product[T constraints.Numeric](slice []T) T {
	return ProductBy(slice, ValueAdapter[T]())
}

func ProductBy[T any, R constraints.Numeric](slice []T, productFunc func(item T) R) R {
	var product R = 1

	if len(slice) == 0 {
		return product
	}

	Apply(slice, func(item T) {
		product *= productFunc(item)
	})

	return product
}

func Mean[T constraints.Numeric](slice []T) T {
	return MeanBy(slice, ValueAdapter[T]())
}

func MeanBy[T any, R constraints.Numeric](slice []T, valueFunc func(item T) R) R {
	count := R(len(slice))

	if count == 0 {
		return count
	}

	return SumBy(slice, valueFunc) / count
}

func Max[T cmp.Ordered](slice []T) T {
	return MaxBy(slice, func(a T, b T) bool { return a < b })
}

func MaxBy[T any](slice []T, maxFunc func(a T, b T) bool) T {
	var maxValue T

	Apply(slice, func(item T) {
		if maxFunc(maxValue, item) {
			maxValue = item
		}
	})

	return maxValue
}

func Min[T cmp.Ordered](slice []T) T {
	return MinBy(slice, func(a T, b T) bool { return a > b })
}

func MinBy[T any](slice []T, minFunc func(a T, b T) bool) T {
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
