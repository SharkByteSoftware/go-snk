// Package mapx provides helper functions for working with maps.
package mapx

import (
	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

// Keys returns a slice of the map keys.
func Keys[M ~map[K]V, K comparable, V any](collection M) []K {
	keys := make([]K, 0, len(collection))

	Apply(collection, func(key K, _ V) {
		keys = append(keys, key)
	})

	return keys
}

// Values returns a slice of the map values.
func Values[M ~map[K]V, K comparable, V any](collection M) []V {
	values := make([]V, 0, len(collection))

	Apply(collection, func(_ K, value V) {
		values = append(values, value)
	})

	return values
}

// UniqueValues returns a slice of all the unique values.
func UniqueValues[M ~map[K]V, K comparable, V comparable](collection M) []V {
	return slicex.Unique(Values(collection))
}

// Contains returns true/false if the map contains the specified key.
func Contains[M ~map[K]V, K comparable, V any](collection M, keys ...K) bool {
	for _, key := range keys {
		if _, ok := collection[key]; !ok {
			return false
		}
	}

	return true
}

// ValueOr returns the value for a key or a fallback value.
func ValueOr[M ~map[K]V, K comparable, V any](collection M, key K, fallback V) V {
	value, ok := collection[key]
	return conditional.If(ok, value, fallback)
}

// Invert inverts the map keys and values.  When there are duplicate values
// no guarantee is made about which key will be used.
func Invert[M ~map[K]V, K comparable, V comparable](collection M) map[V]K {
	result := make(map[V]K, len(collection))

	Apply(collection, func(key K, value V) {
		result[value] = key
	})

	return result
}

// Combine returns a single combined map of all the provided maps.  When there are duplicate
// keys there is no guarantee about which value will be used.
func Combine[M ~map[K]V, K comparable, V any](maps ...M) M {
	size := slicex.SumBy(maps, func(item M) int { return len(item) })
	result := make(M, size)

	slicex.Apply(maps, func(item M) {
		Apply(item, func(key K, value V) {
			result[key] = value
		})
	})

	return result
}

// Merge combines two maps into one, using the resolver function to determine
// the value when a key exists in both maps.
// Unlike Combine, which uses last-write-wins, Merge gives the caller explicit
// control over how conflicts are resolved.
func Merge[M ~map[K]V, K comparable, V any](left M, right M, resolver func(key K, left V, right V) V) M {
	result := make(M, len(left)+len(right))

	Apply(left, func(key K, value V) {
		result[key] = value
	})

	Apply(right, func(key K, value V) {
		if existing, ok := result[key]; ok {
			result[key] = resolver(key, existing, value)
		} else {
			result[key] = value
		}
	})

	return result
}

// ToSlice returns a slice using a mapper function.
func ToSlice[M ~map[K]V, K comparable, V any, R any](collection M, mapper func(key K, value V) R) []R {
	result := make([]R, 0, len(collection))

	Apply(collection, func(key K, value V) {
		result = append(result, mapper(key, value))
	})

	return result
}

// Filter returns a map filtered by the predicate.
func Filter[M ~map[K]V, K comparable, V any](collection M, predicate func(key K, value V) bool) M {
	result := make(M, len(collection))

	Apply(collection, func(key K, value V) {
		if predicate(key, value) {
			result[key] = value
		}
	})

	return result
}

// Apply applies a function to each item in the map.
func Apply[M ~map[K]V, K comparable, V any](collection M, apply func(key K, value V)) {
	for key, value := range collection {
		apply(key, value)
	}
}

// MapKeys returns a new map with each key transformed by the mapper function.
// Values are preserved as-is. When the mapper produces duplicate keys,
// no guarantee is made about which value will be used.
func MapKeys[M ~map[K]V, K comparable, V any, R comparable](collection M, mapper func(key K) R) map[R]V {
	result := make(map[R]V, len(collection))

	Apply(collection, func(key K, value V) {
		result[mapper(key)] = value
	})

	return result
}

// Partition splits a map into two maps based on a predicate.
// Entries for which the predicate returns true are placed in the first map;
// all other entries are placed in the second.
func Partition[M ~map[K]V, K comparable, V any](collection M, predicate func(key K, value V) bool) (M, M) {
	const halfDivisor = 2

	half := len(collection) / halfDivisor
	trueMap := make(M, half)
	falseMap := make(M, half)

	Apply(collection, func(key K, value V) {
		conditional.IfCall(predicate(key, value),
			func() { trueMap[key] = value },
			func() { falseMap[key] = value },
		)
	})

	return trueMap, falseMap
}

// CountBy returns a map of counts keyed by the result of the classifier function.
// Each call to a classifier produces a key; the returned map tracks how many
// entries produced each key.
func CountBy[M ~map[K]V, K comparable, V any, R comparable](collection M, classifier func(key K, value V) R) map[R]int {
	result := make(map[R]int)

	Apply(collection, func(key K, value V) {
		result[classifier(key, value)]++
	})

	return result
}
