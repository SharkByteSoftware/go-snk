// Package mapx provides helper functions for working with maps.
package mapx

import (
	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

// Keys returns a slice of the map keys.
func Keys[M map[K]V, K comparable, V any](collection M) []K {
	keys := make([]K, 0, len(collection))

	Apply(collection, func(key K, _ V) {
		keys = append(keys, key)
	})

	return keys
}

// Values returns a slice of the map values.
func Values[M map[K]V, K comparable, V any](collection M) []V {
	values := make([]V, 0, len(collection))

	Apply(collection, func(_ K, value V) {
		values = append(values, value)
	})

	return values
}

// UniqueValues returns a slice of all the unique values.
func UniqueValues[M map[K]V, K comparable, V comparable](collection M) []V {
	return slicex.Unique(Values(collection))
}

// Contains returns true/false if the map contains the specified key.
func Contains[M map[K]V, K comparable, V any](collection M, keys ...K) bool {
	for _, key := range keys {
		if _, ok := collection[key]; !ok {
			return false
		}
	}

	return true
}

// ValueOr returns the value for a key or a fallback value.
func ValueOr[M map[K]V, K comparable, V any](collection M, key K, fallback V) V {
	value, ok := collection[key]
	return conditional.If(ok, value, fallback)
}

// Invert inverts the map keys and values.  When there are duplicate values
// no guarantee is made about which key will be used.
func Invert[M map[K]V, K comparable, V comparable](collection M) map[V]K {
	result := make(map[V]K, len(collection))

	Apply(collection, func(key K, value V) {
		result[value] = key
	})

	return result
}

// Combine returns a single combined map of all the provided maps.  When there are duplicate
// keys there is no guarantee about which value will be used.
func Combine[M map[K]V, K comparable, V any](maps ...M) M {
	size := slicex.SumBy(maps, func(item M) int { return len(item) })
	result := make(M, size)

	slicex.Apply(maps, func(item M) {
		for key, value := range item {
			result[key] = value
		}
	})

	return result
}

// CombineWithSelect returns a single combined map of all the provided maps and uses select function
// when there is a key collision.   If select returns true, then the previous value is overwritten.
func CombineWithSelect[M map[K]V, K comparable, V any](selector func(V, V) bool, maps ...M) M {
	size := slicex.SumBy(maps, func(item M) int { return len(item) })
	result := make(M, size)

	slicex.Apply(maps, func(item M) {
		for key, value := range item {
			if !Contains(result, key) {
				result[key] = value
				continue
			}

			if selector(result[key], value) {
				result[key] = value
			}
		}
	})

	return result
}

// ToSlice returns a slice using a mapper function.
func ToSlice[M map[K]V, K comparable, V any, R any](collection M, mapper func(key K, value V) R) []R {
	result := make([]R, 0, len(collection))

	Apply(collection, func(key K, value V) {
		result = append(result, mapper(key, value))
	})

	return result
}

// Filter returns a map filtered by the predicate.
func Filter[M map[K]V, K comparable, V any](collection M, predicate func(key K, value V) bool) M {
	result := make(M, len(collection))

	Apply(collection, func(key K, value V) {
		if predicate(key, value) {
			result[key] = value
		}
	})

	return result
}

// Apply applies a function to each item in the map.
func Apply[M map[K]V, K comparable, V any](collection M, apply func(key K, value V)) {
	for key, value := range collection {
		apply(key, value)
	}
}
