package slices

func ItemIndexAdapter[T any, R any](f func(item T) R) func(T, int) R {
	return func(item T, _ int) R {
		return f(item)
	}
}
