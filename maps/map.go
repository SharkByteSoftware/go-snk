// Package maps provides functions for working with maps.
package maps

import "github.com/SharkByteSoftware/go-sink/slices"

// Keys returns an array of the map keys.
func Keys[K comparable, V any](collection map[K]V) []K {
	keys := make([]K, 0, len(collection))

	for key := range collection {
		keys = append(keys, key)
	}

	return keys
}

// Values returns an array of the map values.
func Values[K comparable, V any](collection map[K]V) []V {
	values := make([]V, 0, len(collection))

	for _, value := range collection {
		values = append(values, value)
	}

	return values
}

func UniqueValues[K comparable, V comparable](collection map[K]V) []V {
	// TODO: Implement
	return nil
}

// Contains returns true/false if the map contains the specified key.
func Contains[K comparable, V any](collection map[K]V, key K) bool {
	_, ok := collection[key]
	return ok
}

// Value returns the value for a key or the specified default.
func Value[K comparable, V any](collection map[K]V, key K, fallback V) V {
	if value, ok := collection[key]; ok {
		return value
	}

	return fallback
}

func Invert[K comparable, V comparable](collection map[K]V) map[V]K {
	result := make(map[V]K, len(collection))

	for key, value := range collection {
		result[value] = key
	}

	return result
}

func Combine[K comparable, V any](maps ...map[K]V) map[K]V {
	size := slices.SumBy(maps, func(item map[K]V) int { return len(item) })
	result := make(map[K]V, size)

	slices.Apply(maps, func(item map[K]V) {
		for key, value := range item {
			result[key] = value
		}
	})

	return result
}

func ToSlice[K comparable, V any, R any](collection map[K]V, mapper func(item V) R) []R {
	// TODO: Implement
	return nil
}
