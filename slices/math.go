package slices

import "github.com/SharkByteSoftware/go-sink/constraints"

func Sum[T constraints.Numeric](slice []T) T {
	return SumBy(slice, valueAdapter[T]())
}

func SumBy[T any, R constraints.Numeric](slice []T, sumFunc func(item T) R) R {
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

func Mean[T constraints.Numeric](slice []T) T {
	// TODO: Implement
	return 0
}

func Max[T constraints.Numeric](slice []T) T {
	// TODO: Implement
	return 0
}

func Min[T constraints.Numeric](slice []T) T {
	// TODO: Implement
	return 0
}

func valueAdapter[T any]() func(T) T {
	return func(item T) T { return item }
}
