// Package mapx provides helper functions for working with maps.
package mapx

import (
	"cmp"
	"slices"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

// Keys returns a slice of all keys in the map. The order of the result is not guaranteed.
func Keys[M ~map[K]V, K comparable, V any](collection M) []K {
	keys := make([]K, 0, len(collection))

	Apply(collection, func(key K, _ V) {
		keys = append(keys, key)
	})

	return keys
}

// Values returns a slice of all values in the map. The order of the result is not guaranteed.
func Values[M ~map[K]V, K comparable, V any](collection M) []V {
	values := make([]V, 0, len(collection))

	Apply(collection, func(_ K, value V) {
		values = append(values, value)
	})

	return values
}

// UniqueValues returns a slice of the map's values with duplicates removed. The order of the result is not guaranteed.
func UniqueValues[M ~map[K]V, K comparable, V comparable](collection M) []V {
	return slicex.Unique(Values(collection))
}

// Contains returns true if the map contains all of the specified keys, false otherwise.
func Contains[M ~map[K]V, K comparable, V any](collection M, keys ...K) bool {
	for _, key := range keys {
		if _, ok := collection[key]; !ok {
			return false
		}
	}

	return true
}

// ValueOr returns the value for the given key, or fallback if the key is not present.
func ValueOr[M ~map[K]V, K comparable, V any](collection M, key K, fallback V) V {
	value, ok := collection[key]
	return conditional.If(ok, value, fallback)
}

// Invert returns a new map with keys and values swapped.
// If multiple keys map to the same value, the resulting key in the inverted map is non-deterministic.
func Invert[M ~map[K]V, K comparable, V comparable](collection M) map[V]K {
	result := make(map[V]K, len(collection))

	Apply(collection, func(key K, value V) {
		result[value] = key
	})

	return result
}

// Combine merges all provided maps into a single new map.
// If the same key appears in multiple maps, the resulting value is non-deterministic.
// Use Merge for explicit control over conflict resolution.
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

// Merge combines two maps into a single new map, using the resolver function to determine
// the value when a key exists in both maps. Keys that appear in only one map are included as-is.
// Unlike Combine, which uses non-deterministic last-write-wins, Merge gives the caller explicit
// control over conflict resolution.
func Merge[M ~map[K]V, K comparable, V any](left, right M, resolver func(key K, left, right V) V) M {
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

// ToSlice converts a map into a slice by applying the mapper function to each key-value pair.
// The order of the result is not guaranteed.
func ToSlice[M ~map[K]V, K comparable, V any, R any](collection M, mapper func(key K, value V) R) []R {
	result := make([]R, 0, len(collection))

	Apply(collection, func(key K, value V) {
		result = append(result, mapper(key, value))
	})

	return result
}

// Filter returns a new map containing only the entries for which the predicate returns true.
func Filter[M ~map[K]V, K comparable, V any](collection M, predicate func(key K, value V) bool) M {
	result := make(M, len(collection))

	Apply(collection, func(key K, value V) {
		if predicate(key, value) {
			result[key] = value
		}
	})

	return result
}

// Apply calls the provided function on each key-value pair in the map for side effects. No new map is returned.
func Apply[M ~map[K]V, K comparable, V any](collection M, apply func(key K, value V)) {
	for key, value := range collection {
		apply(key, value)
	}
}

// MapKeys returns a new map with each key transformed by the mapper function.
// Values are preserved unchanged. If the mapper produces duplicate keys, the resulting value is non-deterministic.
// The original map is not modified.
func MapKeys[M ~map[K]V, K comparable, V any, R comparable](collection M, mapper func(key K) R) map[R]V {
	result := make(map[R]V, len(collection))

	Apply(collection, func(key K, value V) {
		result[mapper(key)] = value
	})

	return result
}

// Partition splits a map into two maps based on a predicate.
// Entries for which the predicate returns true are placed in the first map;
// the remaining entries are placed in the second.
func Partition[M ~map[K]V, K comparable, V any](collection M, predicate func(key K, value V) bool) (M, M) {
	const halfDivisor = 2

	half := len(collection) / halfDivisor
	trueMap := make(M, half)
	falseMap := make(M, half)

	for key, value := range collection {
		if predicate(key, value) {
			trueMap[key] = value
			continue
		}

		falseMap[key] = value
	}

	return trueMap, falseMap
}

// Count returns the number of entries in the map whose value equals the given candidate.
// It is the direct-equality counterpart to CountBy.
func Count[M ~map[K]V, K comparable, V comparable](collection M, candidate V) int {
	n := 0

	Apply(collection, func(_ K, value V) {
		if value == candidate {
			n++
		}
	})

	return n
}

// CountBy returns a map of counts grouped by the result of the classifier function.
// The classifier is called for each entry; the returned map tracks how many entries produced each key.
func CountBy[M ~map[K]V, K comparable, V any, R comparable](collection M, classifier func(key K, value V) R) map[R]int {
	result := make(map[R]int)

	Apply(collection, func(key K, value V) {
		result[classifier(key, value)]++
	})

	return result
}

// MapValues returns a new map with each value transformed by the mapper function.
// Keys are preserved unchanged. The original map is not modified.
func MapValues[M ~map[K]V, K comparable, V, W any](m M, mapper func(V) W) map[K]W {
	out := make(map[K]W, len(m))

	for k, v := range m {
		out[k] = mapper(v)
	}

	return out
}

// Any returns true if any entry in the map satisfies the predicate. Returns false for an empty map.
func Any[M ~map[K]V, K comparable, V any](m M, predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}

	return false
}

// All returns true if every entry in the map satisfies the predicate.
// Returns true for an empty map.
func All[M ~map[K]V, K comparable, V any](m M, predicate func(K, V) bool) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}

	return true
}

// SortedKeys returns the keys of the map sorted in ascending order.
// The key type must satisfy cmp.Ordered. Use SortedKeysByFunc for custom ordering.
func SortedKeys[M ~map[K]V, K cmp.Ordered, V any](m M) []K {
	keys := Keys(m)
	slices.Sort(keys)

	return keys
}

// SortedKeysByFunc returns the keys of the map sorted using the provided comparison function.
// cmpFn should return a negative number when a < b, zero when a == b, and a positive number when a > b,
// consistent with the cmp.Compare convention.
func SortedKeysByFunc[M ~map[K]V, K comparable, V any](m M, cmpFn func(a, b K) int) []K {
	keys := Keys(m)
	slices.SortFunc(keys, cmpFn)

	return keys
}
