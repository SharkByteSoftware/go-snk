// Package helpers provides various helpers such as pointer helpers.
package helpers

// Empty returns the empty/zero value of a type.
func Empty[T any]() T {
	var empty T
	return empty
}

// SafeDeref returns the pointer value of or the types empty value.
func SafeDeref[T any](ptr *T) T {
	return SafeDerefOr(ptr, Empty[T]())
}

// SafeDerefOr returns the pointer value or the specified fallback value.
func SafeDerefOr[T any](ptr *T, fallback T) T {
	if ptr == nil {
		return fallback
	}

	return *ptr
}
