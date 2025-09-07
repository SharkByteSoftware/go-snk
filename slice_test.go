package sink_test

import (
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := sink.Filter(test.input, test.filter)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFilterI(t *testing.T) {
	var indexes []int

	tests := []struct {
		name          string
		input         []int
		filter        func(int, int) bool
		expected      []int
		expected_idxs []int
	}{
		{
			name:  "filter even numbers",
			input: numberList,
			filter: func(n int, i int) bool {
				if n%2 != 0 {
					return false
				}

				indexes = append(indexes, i)

				return true
			},
			expected:      []int{2, 4, 256},
			expected_idxs: []int{1, 3, 6},
		},
		{
			name:  "filter nothing",
			input: numberList,
			filter: func(n int, i int) bool {
				return false
			},
			expected:      []int{},
			expected_idxs: []int{},
		},
		{
			name:  "filter everything",
			input: numberList,
			filter: func(n int, i int) bool {
				indexes = append(indexes, i)
				return true
			},
			expected:      numberList,
			expected_idxs: []int{0, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			indexes = make([]int, 0)
			result := sink.FilterI(test.input, test.filter)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expected_idxs, indexes)
		})
	}
}
