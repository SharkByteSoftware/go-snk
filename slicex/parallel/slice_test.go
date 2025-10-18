package parallel_test

import (
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/slicex/parallel"
	"github.com/stretchr/testify/assert"
)

var numberList = []int{1, 2, 3, 4, 5, 333, 256}
var duplicateList = []int{1, 2, 3, 4, 5, 333, 256, 1, 2, 3, 4, 5, 333, 256}
var allSame = []int{1, 1, 1, 1, 1, 1}

func TestParallelSlice_Map(t *testing.T) {
	result := parallel.Map(numberList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = parallel.Map([]int{}, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))
}

func TestParallelSlice_MapWithLimit(t *testing.T) {
	result := parallel.MapWithLimit(numberList, func(n int) string { return strconv.Itoa(n) }, 2)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = parallel.MapWithLimit([]int{}, func(n int) string { return strconv.Itoa(n) }, 1)
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))

	result = parallel.MapWithLimit(numberList, func(n int) string { return strconv.Itoa(n) }, len(numberList)+1)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))
}

func TestParallelSlice_Apply(t *testing.T) {
	var counter int32
	parallel.Apply(numberList, func(n int) { atomic.AddInt32(&counter, 1) })
	assert.Equal(t, int32(7), counter)
}

func TestParallelSlice_ApplyWithLimit(t *testing.T) {
	var counter int32
	parallel.ApplyWithLimit(numberList, func(n int) { atomic.AddInt32(&counter, 1) }, 2)
	assert.Equal(t, int32(7), counter)
}

func TestParallelSlice_GroupBy(t *testing.T) {
	result := parallel.GroupBy(numberList, func(item int) int { return item })
	assert.Len(t, result, len(numberList))

	result = parallel.GroupBy(duplicateList, func(item int) int { return item })
	assert.Len(t, result, 7)

	stringMap := parallel.GroupBy(duplicateList, func(item int) string {
		return conditional.If(item%2 == 0, "even", "odd")
	})
	assert.Len(t, stringMap, 2)
	assert.Len(t, stringMap["even"], 6)
	assert.Len(t, stringMap["odd"], 8)
}

func TestParallelSlice_GroupByWithLimit(t *testing.T) {
	result := parallel.GroupByWithLimit(numberList, func(item int) int { return item }, 2)
	assert.Len(t, result, len(numberList))

	result = parallel.GroupByWithLimit(duplicateList, func(item int) int { return item }, 2)
	assert.Len(t, result, 7)

	stringMap := parallel.GroupByWithLimit(duplicateList, func(item int) string {
		return conditional.If(item%2 == 0, "even", "odd")
	}, 5)
	assert.Len(t, stringMap, 2)
	assert.Len(t, stringMap["even"], 6)
	assert.Len(t, stringMap["odd"], 8)
}

func TestParallelSlice_Partition(t *testing.T) {
	r1, r2 := parallel.Partition([]int{}, func(item int) bool { return true })
	assert.Len(t, r1, 0)
	assert.Len(t, r2, 0)

	r1, r2 = parallel.Partition(numberList, func(item int) bool { return true })
	assert.Len(t, r1, 7)
	assert.Len(t, r2, 0)

	r1, r2 = parallel.Partition(numberList, func(item int) bool { return false })
	assert.Len(t, r1, 0)
	assert.Len(t, r2, 7)

	r1, r2 = parallel.Partition(numberList, func(item int) bool { return item%2 == 0 })
	assert.Len(t, r1, 3)
	assert.Len(t, r2, 4)
}

func TestParallelSlice_PartitionWithLimit(t *testing.T) {
	r1, r2 := parallel.PartitionWithLimit([]int{}, func(item int) bool { return true }, 1)
	assert.Len(t, r1, 0)
	assert.Len(t, r2, 0)

	r1, r2 = parallel.PartitionWithLimit(numberList, func(item int) bool { return true }, 2)
	assert.Len(t, r1, 7)
	assert.Len(t, r2, 0)

	r1, r2 = parallel.PartitionWithLimit(numberList, func(item int) bool { return false }, 4)
	assert.Len(t, r1, 0)
	assert.Len(t, r2, 7)

	r1, r2 = parallel.PartitionWithLimit(numberList, func(item int) bool { return item%2 == 0 }, 2)
	assert.Len(t, r1, 3)
	assert.Len(t, r2, 4)
}
