// Package helpers provides various helpers such as pointer helpers.
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

// Nil returns a nil pointer of a type.
func Nil[T any]() *T {
	return nil
}

// AsPtr returns a pointer copy of a value.
func AsPtr[T any](value T) *T {
	return &value
}

// AsValue returns the value of the pointer.  If the pointer
// is nil, it returns the types default value.
func AsValue[T any](ptr *T) T {
	return AsValueOr(ptr, Empty[T]())
}

// AsValueOr returns the value of the pointer.  If the pointer
// is nil, it returns the specified fallback value.
func AsValueOr[T any](ptr *T, fallback T) T {
	if ptr == nil {
		return fallback
	}

	return *ptr
}
