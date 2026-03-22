package helpers

// Empty returns the empty/zero value of a type.
func Empty[T any]() T {
	var empty T
	return empty
}

// IsEmpty returns true of the value is the zero value of the type.
func IsEmpty[T comparable](value T) bool {
	var empty T
	return empty == value
}
