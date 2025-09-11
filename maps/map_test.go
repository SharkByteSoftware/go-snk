package maps_test

import (
	"testing"

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
