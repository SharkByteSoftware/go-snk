package adapt

// ItemIndexAdapter is a function adapter that ignores the index.
func ItemIndexAdapter[T any, R any](f func(item T) R) func(T, int) R {
	return func(item T, _ int) R {
		return f(item)
	}
}

// ValueAdapter is a function adapter that ignores the index and returns the item.
func ValueAdapter[T any](item T) T {
	return item
}

// ItemEqualsAdapter is a function adapter that compares two items for equality.
func ItemEqualsAdapter[T comparable](item T) func(T) bool {
	return func(other T) bool { return item == other }
}

// KeySelectorAdapter is a function adapter that ignores the value and returns the key.
func KeySelectorAdapter[K comparable, V any](key K, _ V) K {
	return key
}

// ValueSelectorAdapter is a function adapter that ignores the key and returns the value.
func ValueSelectorAdapter[K comparable, V any](_ K, value V) V {
	return value
}
