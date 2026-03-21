// Package helpers provides various helpers such as pointer helpers.
package helpers

// Nil returns a nil pointer of a type.
func Nil[T any]() *T {
	return nil
}

// IsNil returns true if the value is nil.
func IsNil[T any](value *T) bool {
	return value == nil
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
