package slices

import "github.com/SharkByteSoftware/go-sink/constraints"

func Sum[T any](slice []T) int {
	// TODO: Implement
	return 0
}

func SumBy[T any, R int](slice []T, sumFunc func(T) R) R {
	var sum R

	for _, value := range slice {
		sum += sumFunc(value)
	}

	return sum
}

func Product[T constraints.Numeric](slice []T) T {
	// TODO: Implement
	return 0
}

func ProductBy[T any, R int](slice []T, prodFunc func(T) R) R {
	// TODO: Implement
	return 0
}

func Mean[T constraints.Numeric](slice []T) T {
	// TODO: Implement
	return 0
}

func MeanBy[T any, R int](slice []T, meanFunc func(T) R) R {
	// TODO: Implement
	return 0
}

func Max[T constraints.Numeric](slice []T) T {
	// TODO: Implement
	return 0
}

func MaxBy[T any, R int](slice []T, maxFunc func(T) R) R {
	// TODO: Implement
	return 0
}

func Min[T constraints.Numeric](slice []T) T {
	// TODO: Implement
	return 0
}

func MinBy[T any, R int](slice []T, minFunc func(T) R) R {
	// TODO: Implement
	return 0
}
