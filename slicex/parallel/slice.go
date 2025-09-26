package parallel

// Map transforms a slice to a slice of another type using a mapper function.
func Map[S ~[]T, T any, R any](slice S, mapper func(item T) R) []R {
	// TODO: implement
	panic("implement me")
}

// Apply applies a function to each item in the slice.
func Apply[S ~[]T, T any](slice S, apply func(item T)) {
	// TODO: implement
	panic("implement me")
}

// GroupBy returns a map of slices grouped by a key produced by a predicate function.
func GroupBy[S ~[]T, T any, R comparable](slice S, predicate func(item T) R) map[R][]T {
	// TODO: implement
	panic("implement me")
}

// Partition splits a slice into two slices based on a predicate.
func Partition[S ~[]T, T any](slice S, predicate func(item T) bool) (S, S) {
	// TODO: implement
	panic("implement me")
}
