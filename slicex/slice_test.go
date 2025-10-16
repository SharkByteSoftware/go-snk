package slicex_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/internal/adapt"
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

func TestSlice_FirstOr(t *testing.T) {
	result := slicex.FirstOr([]int{}, 20)
	assert.Equal(t, 20, result)

	result = slicex.FirstOr([]int{1, 2}, 10)
	assert.Equal(t, 1, result)
}

func TestSlice_FirstOrEmpty(t *testing.T) {
	result := slicex.FirstOrEmpty([]int{})
	assert.Equal(t, 0, result)

	result = slicex.FirstOrEmpty([]int{1, 2, 4})
	assert.Equal(t, 1, result)
}

func TestSlice_Filter(t *testing.T) {
	result := slicex.Filter(numberList, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, []int{2, 4, 256}, result)

	result = slicex.Filter(numberList, func(n int) bool { return false })
	assert.Equal(t, []int{}, result)

	result = slicex.Filter(numberList, func(n int) bool { return true })
	assert.Equal(t, numberList, result)

	result = slicex.Filter([]int{}, func(n int) bool { return true })
	assert.Equal(t, []int{}, result)
}

func TestSlice_Map(t *testing.T) {
	result := slicex.Map(numberList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = slicex.Map([]int{}, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))
}

func TestSlice_UniqueMap(t *testing.T) {
	result := slicex.UniqueMap(duplicateList, func(n int) string { return strconv.Itoa(n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Len(t, result, 7)
}

func TestSlice_Bind(t *testing.T) {
	x := [][]int{{1, 2}, {3, 4}}
	_ = slicex.Bind(x, adapt.ValueAdapter)

	result := slicex.Bind(nestedNumberList, adapt.ValueAdapter)
	assert.Equal(t, append(numberList, numberList...), result)
}

func TestSlice_Reduce(t *testing.T) {
	accumulator := func(agg int, item []int) int {
		return agg + item[0]
	}

	result := slicex.Reduce(nestedNumberList, accumulator, 0)
	assert.Equal(t, 2, result)
}

func TestSlice_Find(t *testing.T) {
	result, found := slicex.Find(numberList, 88)
	assert.False(t, found)
	assert.Equal(t, 0, result)

	result, found = slicex.Find(numberList, 256)
	assert.True(t, found)
	assert.Equal(t, 256, result)
}

func TestSlice_FindBy(t *testing.T) {
	type myStruct struct {
		Name  string
		Age   int
		Other []myStruct
	}

	p := []myStruct{{Name: "one", Age: 1}, {Name: "two", Age: 2}, {Name: "three", Age: 3}}

	result, found := slicex.FindBy(p, func(item myStruct) bool {
		return item.Name == "one"
	})
	assert.True(t, found)
	assert.Equal(t, myStruct{Name: "one", Age: 1}, result)
}

func TestSlice_FindOr(t *testing.T) {
	result := slicex.FindOr(numberList, 88, 8192)
	assert.Equal(t, 8192, result)

	result = slicex.FindOr(numberList, 256, -1)
	assert.Equal(t, 256, result)
}

func TestSlice_Any(t *testing.T) {
	result := slicex.Any(numberList, 0)
	assert.False(t, result)

	result = slicex.Any(numberList, 1)
	assert.True(t, result)

	result = slicex.Any(duplicateList, 256)
	assert.True(t, result)
}

func TestSlice_All(t *testing.T) {
	result := slicex.All(numberList, 1)
	assert.False(t, result)

	result = slicex.All(allSame, 1)
	assert.True(t, result)
}

func TestSlice_Unique(t *testing.T) {
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

func TestSlice_UniqueBy(t *testing.T) {
	type myStruct struct {
		Name string
		Age  int
	}

	myStructs := []myStruct{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"three", 3},
		{"one", 1},
	}

	result := slicex.UniqueBy(myStructs, func(item myStruct) string { return item.Name })
	assert.Len(t, result, 3)
	assert.Equal(t, myStructs[:3], result)

	result = slicex.UniqueBy(myStructs, func(item myStruct) int { return item.Age })
	assert.Len(t, result, 3)
	assert.Equal(t, myStructs[:3], result)
}

func TestSlice_Reverse(t *testing.T) {
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

func TestSlice_Compact(t *testing.T) {
	numberList := []int{1, 2, 3, 4, 5}

	result := slicex.Compact(numberList)
	assert.Equal(t, numberList, result)

	numberList = []int{0, 2, 3, 4, 0}
	result = slicex.Compact(numberList)
	assert.Equal(t, numberList[1:4], result)

	strList := []string{"", "one", "two", "three", ""}
	strResult := slicex.Compact(strList)
	assert.Equal(t, strList[1:4], strResult)
}

func TestSlice_Apply(t *testing.T) {
	var nums string
	slicex.Apply(numberList, func(n int) { nums += strconv.Itoa(n) })
	assert.Equal(t, "12345333256", nums)
}

func TestSlice_ToMap(t *testing.T) {
	mapperFunc := func(item int) string { return strconv.Itoa(item) }

	result := slicex.ToMap([]int{}, mapperFunc)
	assert.Len(t, result, 0)
	assert.Equal(t, map[string]int{}, result)

	result = slicex.ToMap(numberList, mapperFunc)
	assert.Len(t, result, len(numberList))
	assert.Equal(t, numberList[0], result[strconv.Itoa(numberList[0])])
}

func TestSlice_GroupBy(t *testing.T) {
	result := slicex.GroupBy(numberList, func(item int) int { return item })
	assert.Len(t, result, len(numberList))

	result = slicex.GroupBy(duplicateList, func(item int) int { return item })
	assert.Len(t, result, 7)

	stringMap := slicex.GroupBy(duplicateList, func(item int) string {
		return conditional.If(item%2 == 0, "even", "odd")
	})
	assert.Len(t, stringMap, 2)
	assert.Len(t, stringMap["even"], 6)
	assert.Len(t, stringMap["odd"], 8)
}

func TestSlice_Partition(t *testing.T) {
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

func TestSlice_Intersect(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{2, 3, 4, 5, 6}

	result := slicex.Intersect(slice1, slice1)
	slices.Sort(result)
	assert.Len(t, result, len(slice1))
	assert.Equal(t, slice1, result)

	result = slicex.Intersect(slice1, slice2)
	assert.Len(t, result, 4)
	assert.NotContains(t, result, 1)
	assert.NotContains(t, result, 6)

	result = slicex.Intersect(slice2, slice1)
	assert.Len(t, result, 4)
	assert.NotContains(t, result, 1)
	assert.NotContains(t, result, 6)
}

func TestSlice_Union(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 256}

	result := slicex.Union(slice1, slice1)
	slices.Sort(result)
	assert.Len(t, result, 5)
	assert.Equal(t, slice1, result)

	result = slicex.Union(slice1, slice2)
	assert.Len(t, result, 8)
	assert.Subset(t, result, slice1)
	assert.Subset(t, result, slice2)

	result = slicex.Union(slice2, slice1)
	assert.Len(t, result, 8)
	assert.Subset(t, result, slice1)
	assert.Subset(t, result, slice2)
}

func TestSlice_Difference(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 256}

	result := slicex.Difference(slice1, slice1)
	assert.Len(t, result, 0)

	result = slicex.Difference(slice1, slice2)
	assert.Len(t, result, 3)
	assert.Subset(t, result, []int{1, 2, 3})
}
