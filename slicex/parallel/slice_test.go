package parallel_test

import (
	"strconv"
	"testing"

	"github.com/SharkByteSoftware/go-snk/slicex/parallel"
	"github.com/stretchr/testify/assert"
)

var numberList = []int{1, 2, 3, 4, 5, 333, 256}

func TestParallelSlice_Map(t *testing.T) {
	result := parallel.Map(numberList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = parallel.Map([]int{}, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))
}
