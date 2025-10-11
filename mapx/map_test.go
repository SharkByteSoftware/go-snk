package mapx_test

import (
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

var numberMapSame = map[int]string{
	0:   "same",
	8:   "same",
	2:   "same",
	3:   "same",
	12:  "same",
	256: "same",
}

var numberMapOther = map[int]string{
	0:   "other",
	8:   "other",
	2:   "other",
	3:   "other",
	12:  "other",
	256: "other",
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
	for k, _ := range numberMap {
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
	for key, _ := range numberMap {
		assert.Contains(t, intResult, key)
	}
}

func TestMap_Filter(t *testing.T) {
	result := mapx.Filter(numberMap, func(k int, v string) bool { return true })
	assert.Equal(t, numberMap, result)

	result = mapx.Filter(numberMap, func(k int, v string) bool { return v == "zero" })
	assert.Equal(t, map[int]string{0: "zero"}, result)

	result = mapx.Filter(numberMap, func(k int, v string) bool { return k%2 == 0 })
	assert.Equal(t, map[int]string{0: "zero", 2: "two", 8: "one", 12: "four", 256: "five"}, result)

	result = mapx.Filter(map[int]string{}, func(k int, v string) bool { return true })
	assert.Equal(t, map[int]string{}, result)
}
