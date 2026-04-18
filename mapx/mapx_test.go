package mapx_test

import (
	"cmp"
	"strconv"
	"strings"
	"testing"

	"github.com/SharkByteSoftware/go-snk/mapx"
	"github.com/stretchr/testify/assert"
)

var numberMap = map[int]string{
	0:   "zero",
	8:   "one",
	2:   "two",
	3:   "three",
	12:  "four",
	256: "five",
}

var contNumberMap = map[int]string{
	200: "zero",
	201: "one",
	202: "two",
	203: "three",
	204: "four",
	205: "five",
}

var dupValueMap = map[int]string{
	0:    "zero",
	8:    "zero",
	2:    "two",
	3:    "five",
	12:   "five",
	256:  "five",
	8192: "five",
}

func TestMap_Keys(t *testing.T) {
	keys := mapx.Keys(numberMap)

	assert.Len(t, keys, 6)

	for k := range numberMap {
		assert.Contains(t, keys, k)
	}
}

func TestMap_Values(t *testing.T) {
	values := mapx.Values(numberMap)

	assert.Len(t, values, 6)

	for _, v := range numberMap {
		assert.Contains(t, values, v)
	}
}

func TestMap_UniqueValues(t *testing.T) {
	values := mapx.UniqueValues(numberMap)
	assert.Len(t, values, 6)

	for _, v := range numberMap {
		assert.Contains(t, values, v)
	}

	values = mapx.UniqueValues(dupValueMap)
	assert.Len(t, values, 3)
	assert.Contains(t, values, "zero")
	assert.Contains(t, values, "two")
	assert.Contains(t, values, "five")
}

func TestMap_Contains(t *testing.T) {
	assert.True(t, mapx.Contains(numberMap, 256))

	assert.False(t, mapx.Contains(numberMap, 257))

	assert.False(t, mapx.Contains(map[int]int{}, 0))
}

func TestMap_ValueOr(t *testing.T) {
	result := mapx.ValueOr(numberMap, 0, "negative")
	assert.Equal(t, "zero", result)

	result = mapx.ValueOr(numberMap, 257, "negative")
	assert.Equal(t, "negative", result)

	result = mapx.ValueOr(map[int]string{}, 12, "negative")
	assert.Equal(t, "negative", result)
}

func TestMap_Invert(t *testing.T) {
	inverted := mapx.Invert(numberMap)
	assert.Len(t, inverted, 6)

	for k, v := range inverted {
		assert.Contains(t, numberMap, v)
		assert.Equal(t, k, numberMap[v])
	}

	inverted = mapx.Invert(dupValueMap)
	assert.Len(t, inverted, 3)
	assert.Contains(t, inverted, "five", "zero", "two")
}

func TestMap_Combine(t *testing.T) {
	result := mapx.Combine(numberMap, numberMap)
	assert.Len(t, result, 6)
	assert.Equal(t, numberMap, result)

	result = mapx.Combine(numberMap, dupValueMap)
	assert.Len(t, result, 7)
	assert.Equal(t, dupValueMap, result)

	result = mapx.Combine(numberMap, contNumberMap)
	assert.Len(t, result, 12)

	for k, v := range numberMap {
		assert.Contains(t, result, k)
		assert.Equal(t, v, result[k])
	}

	for k, v := range contNumberMap {
		assert.Contains(t, result, k)
		assert.Equal(t, v, result[k])
	}
}

func TestMap_Merge(t *testing.T) {
	keepLeft := func(_ int, left, _ string) string { return left }
	keepRight := func(_ int, _, right string) string { return right }
	concat := func(_ int, left, right string) string { return left + right }

	// no overlap — all keys from both maps appear in result
	left := map[int]string{0: "zero", 1: "one"}
	right := map[int]string{2: "two", 3: "three"}

	result := mapx.Merge(left, right, keepLeft)
	assert.Len(t, result, 4)
	assert.Equal(t, "zero", result[0])
	assert.Equal(t, "two", result[2])

	// conflict — resolver receives both values
	left = map[int]string{0: "zero", 1: "one"}
	right = map[int]string{1: "ONE", 2: "two"}

	result = mapx.Merge(left, right, keepLeft)
	assert.Len(t, result, 3)
	assert.Equal(t, "one", result[1]) // left wins

	result = mapx.Merge(left, right, keepRight)
	assert.Len(t, result, 3)
	assert.Equal(t, "ONE", result[1]) // right wins

	result = mapx.Merge(left, right, concat)
	assert.Len(t, result, 3)
	assert.Equal(t, "oneONE", result[1]) // combined

	// resolver receives the correct key
	var seenKey int
	mapx.Merge(
		map[int]string{5: "five"},
		map[int]string{5: "FIVE"},
		func(key int, left, _ string) string { seenKey = key; return left },
	)
	assert.Equal(t, 5, seenKey)

	// empty maps
	assert.Empty(t, mapx.Merge(map[int]string{}, map[int]string{}, keepLeft))

	result = mapx.Merge(map[int]string{}, right, keepLeft)
	assert.Equal(t, right, result)

	result = mapx.Merge(left, map[int]string{}, keepLeft)
	assert.Equal(t, left, result)
}

func TestMap_ToSlice(t *testing.T) {
	stringResult := mapx.ToSlice(numberMap, func(_ int, value string) string { return value })

	assert.Len(t, stringResult, 6)

	for _, value := range numberMap {
		assert.Contains(t, stringResult, value)
	}

	intResult := mapx.ToSlice(numberMap, func(key int, _ string) int { return key })
	assert.Len(t, intResult, 6)

	for key := range numberMap {
		assert.Contains(t, intResult, key)
	}
}

func TestMap_Filter(t *testing.T) {
	result := mapx.Filter(numberMap, func(_ int, _ string) bool { return true })
	assert.Equal(t, numberMap, result)

	result = mapx.Filter(numberMap, func(_ int, v string) bool { return v == "zero" })
	assert.Equal(t, map[int]string{0: "zero"}, result)

	result = mapx.Filter(numberMap, func(k int, _ string) bool { return k%2 == 0 })
	assert.Equal(t, map[int]string{0: "zero", 2: "two", 8: "one", 12: "four", 256: "five"}, result)

	result = mapx.Filter(map[int]string{}, func(_ int, _ string) bool { return true })
	assert.Equal(t, map[int]string{}, result)
}

func TestMap_MapKeys(t *testing.T) {
	// transforms keys using the mapper
	result := mapx.MapKeys(numberMap, strconv.Itoa)

	assert.Len(t, result, len(numberMap))

	for k, v := range numberMap {
		assert.Contains(t, result, strconv.Itoa(k))
		assert.Equal(t, v, result[strconv.Itoa(k)])
	}

	// empty map returns empty map
	result = mapx.MapKeys(map[int]string{}, strconv.Itoa)

	assert.Empty(t, result)

	// duplicate mapped keys: result length is <= original length
	dupKeyMap := map[int]string{1: "one", 2: "two", 3: "three"}
	collapsed := mapx.MapKeys(dupKeyMap, func(_ int) string { return "same" })

	assert.Len(t, collapsed, 1)
	assert.Contains(t, collapsed, "same")
}

func TestMap_Partition(t *testing.T) {
	// all entries match predicate
	trueMap, falseMap := mapx.Partition(numberMap, func(_ int, _ string) bool { return true })

	assert.Equal(t, numberMap, trueMap)
	assert.Empty(t, falseMap)

	// no entries match predicate
	trueMap, falseMap = mapx.Partition(numberMap, func(_ int, _ string) bool { return false })

	assert.Empty(t, trueMap)
	assert.Equal(t, numberMap, falseMap)

	// split by key parity
	trueMap, falseMap = mapx.Partition(numberMap, func(k int, _ string) bool { return k%2 == 0 })

	for k, v := range trueMap {
		assert.Equal(t, 0, k%2, "expected even key in true map, got %d", k)
		assert.Equal(t, numberMap[k], v)
	}

	for k, v := range falseMap {
		assert.NotEqual(t, 0, k%2, "expected odd key in false map, got %d", k)
		assert.Equal(t, numberMap[k], v)
	}

	assert.Len(t, trueMap, 5)
	assert.Len(t, falseMap, 1)

	// split by value
	trueMap, falseMap = mapx.Partition(numberMap, func(_ int, v string) bool { return v == "zero" })

	assert.Equal(t, map[int]string{0: "zero"}, trueMap)
	assert.Len(t, falseMap, 5)

	// empty map returns two empty maps
	trueMap, falseMap = mapx.Partition(map[int]string{}, func(_ int, _ string) bool { return true })

	assert.Empty(t, trueMap)
	assert.Empty(t, falseMap)
}

func TestMap_CountBy(t *testing.T) {
	// count by value
	counts := mapx.CountBy(numberMap, func(_ int, v string) string { return v })

	assert.Len(t, counts, len(numberMap))

	for _, v := range numberMap {
		assert.Equal(t, 1, counts[v])
	}

	// count by value with duplicates
	counts = mapx.CountBy(dupValueMap, func(_ int, v string) string { return v })

	assert.Equal(t, 2, counts["zero"])
	assert.Equal(t, 1, counts["two"])
	assert.Equal(t, 4, counts["five"])

	// count by classifier on key (even vs odd)
	keyCounts := mapx.CountBy(numberMap, func(k int, _ string) string {
		if k%2 == 0 {
			return "even"
		}

		return "odd"
	})

	assert.Equal(t, 5, keyCounts["even"])
	assert.Equal(t, 1, keyCounts["odd"])

	// all entries produce same classifier key
	sameKey := mapx.CountBy(numberMap, func(_ int, _ string) string { return "all" })

	assert.Len(t, sameKey, 1)
	assert.Equal(t, len(numberMap), sameKey["all"])

	// empty map returns empty counts
	emptyCounts := mapx.CountBy(map[int]string{}, func(_ int, v string) string { return v })

	assert.Empty(t, emptyCounts)
}

func TestMap_Merge_ResolverCalledOnce(t *testing.T) {
	left := map[int]string{1: "one", 2: "two"}
	right := map[int]string{1: "ONE", 3: "three"}

	callCount := 0

	var seenLeft, seenRight string

	mapx.Merge(left, right, func(_ int, lv, rv string) string {
		callCount++
		seenLeft = lv
		seenRight = rv

		return lv
	})

	// key 1 conflicts — resolver must be called exactly once
	assert.Equal(t, 1, callCount)

	// resolver must receive left value first, right value second
	assert.Equal(t, "one", seenLeft)
	assert.Equal(t, "ONE", seenRight)
}

func TestMap_Merge_ResolverCalledOncePerConflict(t *testing.T) {
	left := map[int]string{1: "a", 2: "b", 3: "c"}
	right := map[int]string{1: "A", 2: "B", 3: "C"}

	callCount := 0

	mapx.Merge(left, right, func(_ int, lv, _ string) string {
		callCount++
		return lv
	})

	// three conflicting keys — resolver must be called exactly three times
	assert.Equal(t, 3, callCount)
}

func Test_MapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := mapx.MapValues(m, func(v int) int { return v * 2 })

	assert.Equal(t, map[string]int{"a": 2, "b": 4, "c": 6}, got)
	// original must not be modified
	assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, m)

	// type-changing mapper
	labels := mapx.MapValues(m, func(v int) string {
		if v > 1 {
			return "big"
		}

		return "small"
	})
	assert.Equal(t, map[string]string{"a": "small", "b": "big", "c": "big"}, labels)

	// empty map
	assert.Equal(t, map[string]int{}, mapx.MapValues(map[string]int{}, func(v int) int { return v }))
}

func Test_MapAny(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	hasEven := func(_ string, v int) bool { return v%2 == 0 }
	allNeg := func(_ string, v int) bool { return v < 0 }

	assert.True(t, mapx.Any(m, hasEven))
	assert.False(t, mapx.Any(m, allNeg))
	assert.False(t, mapx.Any(map[string]int{}, hasEven))
}

func Test_MapAll(t *testing.T) {
	m := map[string]int{"a": 2, "b": 4, "c": 6}
	allEven := func(_ string, v int) bool { return v%2 == 0 }
	hasOdd := func(_ string, v int) bool { return v%2 != 0 }

	assert.True(t, mapx.All(m, allEven))
	assert.False(t, mapx.All(m, hasOdd))
	// empty map always returns true
	assert.True(t, mapx.All(map[string]int{}, allEven))
}

func Test_SortedKeys(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	assert.Equal(t, []string{"a", "b", "c"}, mapx.SortedKeys(m))
	assert.Equal(t, []string{}, mapx.SortedKeys(map[string]int{}))
}

func Test_SortedKeysByFunc(t *testing.T) {
	m := map[string]int{"banana": 2, "apple": 1, "cherry": 3}
	// sort keys by length, then lexicographically
	byLen := func(a, b string) int {
		if n := cmp.Compare(len(a), len(b)); n != 0 {
			return n
		}

		return strings.Compare(a, b)
	}

	got := mapx.SortedKeysByFunc(m, byLen)
	assert.Equal(t, []string{"apple", "banana", "cherry"}, got)
}
