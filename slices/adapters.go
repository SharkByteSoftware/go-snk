package slices

func ItemIndexAdapter[T any, R any](f func(item T) R) func(T, int) R {
	return func(item T, _ int) R {
		return f(item)
	}
}

func ValueAdapter[T any]() func(T) T {
	return func(item T) T { return item }
}

func ItemEqualsAdapter[T comparable](item T) func(T) bool {
	return func(other T) bool { return item == other }
}
