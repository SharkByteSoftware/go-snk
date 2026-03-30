package mapx_test

import (
	"strconv"
	"testing"

	"github.com/SharkByteSoftware/go-snk/internal/adapt"
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

func TestMap_ToSlice(t *testing.T) {
	stringResult := mapx.ToSlice(numberMap, adapt.ValueSelectorAdapter)

	assert.Len(t, stringResult, 6)

	for _, value := range numberMap {
		assert.Contains(t, stringResult, value)
	}

	intResult := mapx.ToSlice(numberMap, adapt.KeySelectorAdapter)
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
