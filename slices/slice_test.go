package slices_test

import (
	"strconv"
	"testing"

	"github.com/SharkByteSoftware/go-snk/slices"
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
	result := slices.Filter(numberList, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, []int{2, 4, 256}, result)

	result = slices.Filter(numberList, func(n int) bool { return false })
	assert.Equal(t, []int{}, result)

	result = slices.Filter(numberList, func(n int) bool { return true })
	assert.Equal(t, numberList, result)

	result = slices.Filter([]int{}, func(n int) bool { return true })
	assert.Equal(t, []int{}, result)
}

func TestMap(t *testing.T) {
	result := slices.Map(numberList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = slices.Map([]int{}, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))
}

func TestUniqueMap(t *testing.T) {
	result := slices.UniqueMap(duplicateList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Len(t, result, 7)
}

func TestBind(t *testing.T) {
	result := slices.Bind(nestedNumberList, slices.ValueAdapter[[]int]())
	assert.Equal(t, append(numberList, numberList...), result)
}

func TestFold(t *testing.T) {
	accumulator := func(agg int, item []int) int {
		return agg + item[0]
	}

	result := slices.Fold(nestedNumberList, accumulator, 0)
	assert.Equal(t, 2, result)
}

func TestFind(t *testing.T) {
	result, found := slices.Find(numberList, 88)
	assert.False(t, found)
	assert.Equal(t, 0, result)

	result, found = slices.Find(numberList, 256)
	assert.True(t, found)
	assert.Equal(t, 256, result)
}

func TestAny(t *testing.T) {
	result := slices.Any(numberList, 0)
	assert.False(t, result)

	result = slices.Any(numberList, 1)
	assert.True(t, result)

	result = slices.Any(duplicateList, 256)
	assert.True(t, result)
}

func TestAll(t *testing.T) {
	result := slices.All(numberList, 1)
	assert.False(t, result)

	result = slices.All(allSame, 1)
	assert.True(t, result)
}

func TestUnique(t *testing.T) {
	result := slices.Unique([]int{})
	assert.Equal(t, []int{}, result)

	result = slices.Unique([]int{1})
	assert.Equal(t, []int{1}, result)

	result = slices.Unique(numberList)
	assert.Equal(t, numberList, result)

	result = slices.Unique(duplicateList)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 333, 256}, result)

	result = slices.Unique([]int{1, 1, 1, 1, 1, 1, 1, 2})
	assert.Equal(t, []int{1, 2}, result)
}

func TestGroupBy(t *testing.T) {
}

func TestReverse(t *testing.T) {
	var orderedList = []int{1, 2, 3, 4, 5, 256}
	var oddNumberedOrderedList = []int{1, 2, 3, 4, 5, 256, 333}

	result := slices.Reverse(orderedList)
	assert.IsDecreasing(t, result)
	assert.IsIncreasing(t, orderedList)

	result = slices.Reverse(oddNumberedOrderedList)
	assert.IsDecreasing(t, result)
	assert.IsIncreasing(t, oddNumberedOrderedList)

	result = slices.Reverse(allSame)
	assert.IsNonDecreasing(t, result)

	result = slices.Reverse([]int{})
	assert.IsDecreasing(t, result)
}

func TestApply(t *testing.T) {
	var nums string
	slices.Apply(numberList, func(n int) { nums += strconv.Itoa(n) })
	assert.Equal(t, "12345333256", nums)
}

func TestToMap(t *testing.T) {
	mapperFunc := func(item int) string { return strconv.Itoa(item) }

	result := slices.ToMap([]int{}, mapperFunc)
	assert.Len(t, result, 0)
	assert.Equal(t, map[string]int{}, result)

	result = slices.ToMap(numberList, mapperFunc)
	assert.Len(t, result, len(numberList))
	assert.Equal(t, numberList[0], result[strconv.Itoa(numberList[0])])
}
