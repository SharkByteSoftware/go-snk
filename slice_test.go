package sink_test

import (
	"strconv"
	"testing"

	"github.com/SharkByteSoftware/go-sink"
	"github.com/stretchr/testify/assert"
)

var numberList = []int{1, 2, 3, 4, 5, 333, 256}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		filter   func(int) bool
		expected []int
	}{
		{
			name:     "filter even numbers",
			input:    numberList,
			filter:   func(n int) bool { return n%2 == 0 },
			expected: []int{2, 4, 256},
		},
		{
			name:     "filter nothing",
			input:    numberList,
			filter:   func(n int) bool { return false },
			expected: []int{},
		},
		{
			name:     "filter everything",
			input:    numberList,
			filter:   func(n int) bool { return true },
			expected: numberList,
		},
		{
			name:     "empty slice",
			input:    []int{},
			filter:   func(n int) bool { return true },
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := sink.Filter(test.input, test.filter)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, len(test.input), cap(result))
		})
	}
}

func TestMap(t *testing.T) {
	result := sink.Map(numberList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = sink.Map([]int{}, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))
}
