package maps_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/adapt"
	"github.com/SharkByteSoftware/go-snk/maps"
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

func TestKeys(t *testing.T) {
	keys := maps.Keys(numberMap)

	assert.Len(t, keys, 6)
	for k, _ := range numberMap {
		assert.Contains(t, keys, k)
	}
}

func TestValues(t *testing.T) {
	values := maps.Values(numberMap)

	assert.Len(t, values, 6)
	for _, v := range numberMap {
		assert.Contains(t, values, v)
	}
}

func TestUniqueValues(t *testing.T) {
	values := maps.UniqueValues(numberMap)
	assert.Len(t, values, 6)
	for _, v := range numberMap {
		assert.Contains(t, values, v)
	}

	values = maps.UniqueValues(dupValueMap)
	assert.Len(t, values, 3)
	assert.Contains(t, values, "zero")
	assert.Contains(t, values, "two")
	assert.Contains(t, values, "five")
}

func TestContains(t *testing.T) {
	assert.True(t, maps.Contains(numberMap, 256))

	assert.False(t, maps.Contains(numberMap, 257))

	assert.False(t, maps.Contains(map[int]int{}, 0))
}

func TestValue(t *testing.T) {
	result := maps.Value(numberMap, 0, "negative")
	assert.Equal(t, "zero", result)

	result = maps.Value(numberMap, 257, "negative")
	assert.Equal(t, "negative", result)

	result = maps.Value(map[int]string{}, 12, "negative")
	assert.Equal(t, "negative", result)
}

func TestInvert(t *testing.T) {
	inverted := maps.Invert(numberMap)
	assert.Len(t, inverted, 6)
	for k, v := range inverted {
		assert.Contains(t, numberMap, v)
		assert.Equal(t, k, numberMap[v])
	}

	inverted = maps.Invert(dupValueMap)
	assert.Len(t, inverted, 3)
	assert.Contains(t, inverted, "five", "zero", "two")
}

func TestCombine(t *testing.T) {
	result := maps.Combine(numberMap, numberMap)
	assert.Len(t, result, 6)
	assert.Equal(t, numberMap, result)

	result = maps.Combine(numberMap, dupValueMap)
	assert.Len(t, result, 7)
	assert.Equal(t, dupValueMap, result)

	result = maps.Combine(numberMap, contNumberMap)
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

func TestToSlice(t *testing.T) {
	stringResult := maps.ToSlice(numberMap, adapt.ValueSelectorAdapter)

	assert.Len(t, stringResult, 6)
	for _, value := range numberMap {
		assert.Contains(t, stringResult, value)
	}

	intResult := maps.ToSlice(numberMap, adapt.KeySelectorAdapter)
	assert.Len(t, intResult, 6)
	for key, _ := range numberMap {
		assert.Contains(t, intResult, key)
	}
}
