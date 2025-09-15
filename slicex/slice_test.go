package slicex_test

import (
	"strconv"
	"testing"

	"github.com/SharkByteSoftware/go-snk/adapt"
	"github.com/SharkByteSoftware/go-snk/conditionals"
	"github.com/SharkByteSoftware/go-snk/slicex"
	"github.com/stretchr/testify/assert"
)

var numberList = []int{1, 2, 3, 4, 5, 333, 256}
var duplicateList = []int{1, 2, 3, 4, 5, 333, 256, 1, 2, 3, 4, 5, 333, 256}
var allSame = []int{1, 1, 1, 1, 1, 1}

var nestedNumberList = [][]int{
	numberList,
	numberList,
}

func TestFilter(t *testing.T) {
	result := slicex.Filter(numberList, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, []int{2, 4, 256}, result)

	result = slicex.Filter(numberList, func(n int) bool { return false })
	assert.Equal(t, []int{}, result)

	result = slicex.Filter(numberList, func(n int) bool { return true })
	assert.Equal(t, numberList, result)

	result = slicex.Filter([]int{}, func(n int) bool { return true })
	assert.Equal(t, []int{}, result)
}

func TestMap(t *testing.T) {
	result := slicex.Map(numberList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = slicex.Map([]int{}, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))
}

func TestUniqueMap(t *testing.T) {
	result := slicex.UniqueMap(duplicateList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Len(t, result, 7)
}

func TestBind(t *testing.T) {
	x := [][]int{{1, 2}, {3, 4}}
	_ = slicex.Bind(x, adapt.ValueAdapter)

	result := slicex.Bind(nestedNumberList, adapt.ValueAdapter)
	assert.Equal(t, append(numberList, numberList...), result)
}

func TestFold(t *testing.T) {
	accumulator := func(agg int, item []int) int {
		return agg + item[0]
	}

	result := slicex.Fold(nestedNumberList, accumulator, 0)
	assert.Equal(t, 2, result)
}

func TestFind(t *testing.T) {
	result, found := slicex.Find(numberList, 88)
	assert.False(t, found)
	assert.Equal(t, 0, result)

	result, found = slicex.Find(numberList, 256)
	assert.True(t, found)
	assert.Equal(t, 256, result)
}

func TestFindOr(t *testing.T) {
	result := slicex.FindOr(numberList, 88, 8192)
	assert.Equal(t, 8192, result)

	result = slicex.FindOr(numberList, 256, -1)
	assert.Equal(t, 256, result)
}

func TestAny(t *testing.T) {
	result := slicex.Any(numberList, 0)
	assert.False(t, result)

	result = slicex.Any(numberList, 1)
	assert.True(t, result)

	result = slicex.Any(duplicateList, 256)
	assert.True(t, result)
}

func TestAll(t *testing.T) {
	result := slicex.All(numberList, 1)
	assert.False(t, result)

	result = slicex.All(allSame, 1)
	assert.True(t, result)
}

func TestUnique(t *testing.T) {
	result := slicex.Unique([]int{})
	assert.Equal(t, []int{}, result)

	result = slicex.Unique([]int{1})
	assert.Equal(t, []int{1}, result)

	result = slicex.Unique(numberList)
	assert.Equal(t, numberList, result)

	result = slicex.Unique(duplicateList)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 333, 256}, result)

	result = slicex.Unique([]int{1, 1, 1, 1, 1, 1, 1, 2})
	assert.Equal(t, []int{1, 2}, result)
}

func TestReverse(t *testing.T) {
	var orderedList = []int{1, 2, 3, 4, 5, 256}
	var oddNumberedOrderedList = []int{1, 2, 3, 4, 5, 256, 333}

	result := slicex.Reverse(orderedList)
	assert.IsDecreasing(t, result)
	assert.IsIncreasing(t, orderedList)

	result = slicex.Reverse(oddNumberedOrderedList)
	assert.IsDecreasing(t, result)
	assert.IsIncreasing(t, oddNumberedOrderedList)

	result = slicex.Reverse(allSame)
	assert.IsNonDecreasing(t, result)

	result = slicex.Reverse([]int{})
	assert.IsDecreasing(t, result)
}

func TestApply(t *testing.T) {
	var nums string
	slicex.Apply(numberList, func(n int) { nums += strconv.Itoa(n) })
	assert.Equal(t, "12345333256", nums)
}

func TestToMap(t *testing.T) {
	mapperFunc := func(item int) string { return strconv.Itoa(item) }

	result := slicex.ToMap([]int{}, mapperFunc)
	assert.Len(t, result, 0)
	assert.Equal(t, map[string]int{}, result)

	result = slicex.ToMap(numberList, mapperFunc)
	assert.Len(t, result, len(numberList))
	assert.Equal(t, numberList[0], result[strconv.Itoa(numberList[0])])
}

func TestGroupBy(t *testing.T) {
	result := slicex.GroupBy(numberList, func(item int) int { return item })
	assert.Len(t, result, len(numberList))

	result = slicex.GroupBy(duplicateList, func(item int) int { return item })
	assert.Len(t, result, 7)

	stringMap := slicex.GroupBy(duplicateList, func(item int) string {
		return conditionals.If(item%2 == 0, "even", "odd")
	})
	assert.Len(t, stringMap, 2)
	assert.Len(t, stringMap["even"], 6)
	assert.Len(t, stringMap["odd"], 8)
}

func TestPartition(t *testing.T) {
	r1, r2 := slicex.Partition([]int{}, func(item int) bool { return true })
	assert.Len(t, r1, 0)
	assert.Len(t, r2, 0)

	r1, r2 = slicex.Partition(numberList, func(item int) bool { return true })
	assert.Len(t, r1, 7)
	assert.Len(t, r2, 0)

	r1, r2 = slicex.Partition(numberList, func(item int) bool { return false })
	assert.Len(t, r1, 0)
	assert.Len(t, r2, 7)

	r1, r2 = slicex.Partition(numberList, func(item int) bool { return item%2 == 0 })
	assert.Len(t, r1, 3)
	assert.Len(t, r2, 4)
}
