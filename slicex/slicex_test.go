package slicex_test

import (
	"cmp"
	"slices"
	"strconv"
	"testing"

	"github.com/SharkByteSoftware/go-snk/conditional"
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

func TestSlice_FirstBy(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	result, found := slicex.FirstBy([]int{}, isEven)
	assert.False(t, found)
	assert.Equal(t, 0, result)

	result, found = slicex.FirstBy([]int{1, 3, 5}, isEven)
	assert.False(t, found)
	assert.Equal(t, 0, result)

	result, found = slicex.FirstBy([]int{1, 2, 4, 6}, isEven)
	assert.True(t, found)
	assert.Equal(t, 2, result)
}

func TestSlice_FirstOrBy(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	assert.Equal(t, -1, slicex.FirstOrBy([]int{1, 3, 5}, isEven, -1))
	assert.Equal(t, 2, slicex.FirstOrBy([]int{1, 2, 4}, isEven, -1))
}

func TestSlice_LastOr(t *testing.T) {
	assert.Equal(t, -1, slicex.LastOr([]int{}, -1))
	assert.Equal(t, 3, slicex.LastOr([]int{1, 2, 3}, -1))
}

func TestSlice_LastOrEmpty(t *testing.T) {
	assert.Equal(t, 0, slicex.LastOrEmpty([]int{}))
	assert.Equal(t, 3, slicex.LastOrEmpty([]int{1, 2, 3}))
}

func TestSlice_LastBy(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	result, found := slicex.LastBy([]int{}, isEven)
	assert.False(t, found)
	assert.Equal(t, 0, result)

	result, found = slicex.LastBy([]int{1, 3, 5}, isEven)
	assert.False(t, found)
	assert.Equal(t, 0, result)

	result, found = slicex.LastBy([]int{2, 4, 6, 7}, isEven)
	assert.True(t, found)
	assert.Equal(t, 6, result)
}

func TestSlice_LastOrBy(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	assert.Equal(t, -1, slicex.LastOrBy([]int{1, 3, 5}, isEven, -1))
	assert.Equal(t, 6, slicex.LastOrBy([]int{2, 4, 6, 7}, isEven, -1))
}

func TestSlice_Filter(t *testing.T) {
	result := slicex.Filter(numberList, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, []int{2, 4, 256}, result)

	result = slicex.Filter(numberList, func(_ int) bool { return false })
	assert.Equal(t, []int{}, result)

	result = slicex.Filter(numberList, func(_ int) bool { return true })
	assert.Equal(t, numberList, result)

	result = slicex.Filter([]int{}, func(_ int) bool { return true })
	assert.Equal(t, []int{}, result)
}

func TestSlice_Map(t *testing.T) {
	result := slicex.Map(numberList, strconv.Itoa)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Equal(t, len(numberList), cap(result))

	result = slicex.Map([]int{}, strconv.Itoa)
	assert.Equal(t, []string{}, result)
	assert.Equal(t, 0, cap(result))
}

func TestSlice_FilterMap(t *testing.T) {
	result := slicex.FilterMap(numberList, func(item int) (string, bool) {
		if item%2 == 0 {
			return "", false
		}

		return strconv.Itoa(item), true
	})

	assert.Equal(t, []string{"1", "3", "5", "333"}, result)
}

func TestSlice_UniqueMap(t *testing.T) {
	result := slicex.UniqueMap(duplicateList, strconv.Itoa)
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "333", "256"}, result)
	assert.Len(t, result, 7)
}

func TestSlice_Bind(t *testing.T) {
	x := [][]int{{1, 2}, {3, 4}}
	_ = slicex.Bind(x, func(item []int) []int { return item })

	result := slicex.Bind(nestedNumberList, func(item []int) []int { return item })
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

func TestSlice_Contains(t *testing.T) {
	assert.True(t, slicex.Contains(numberList, 256))
	assert.False(t, slicex.Contains(numberList, 88))
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

func TestSlice_AllBy(t *testing.T) {
	type myStruct struct {
		Name string
		Age  int
	}

	p := []myStruct{{Name: "one", Age: 1}, {Name: "two", Age: 1}, {Name: "three", Age: 1}}

	assert.True(t, slicex.AllBy(p, func(item myStruct) bool { return item.Age == 1 }))
	assert.False(t, slicex.AllBy(p, func(item myStruct) bool { return item.Name == "three" }))
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
	var (
		orderedList            = []int{1, 2, 3, 4, 5, 256}
		oddNumberedOrderedList = []int{1, 2, 3, 4, 5, 256, 333}
	)

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
	result := slicex.ToMap([]int{}, strconv.Itoa)
	assert.Empty(t, result)
	assert.Equal(t, map[string]int{}, result)

	result = slicex.ToMap(numberList, strconv.Itoa)
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
	r1, r2 := slicex.Partition([]int{}, func(_ int) bool { return true })
	assert.Empty(t, r1)
	assert.Empty(t, r2)

	r1, r2 = slicex.Partition(numberList, func(_ int) bool { return true })
	assert.Len(t, r1, 7)
	assert.Empty(t, r2)

	r1, r2 = slicex.Partition(numberList, func(_ int) bool { return false })
	assert.Empty(t, r1)
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
	assert.Empty(t, result)

	result = slicex.Difference(slice1, slice2)
	assert.Len(t, result, 3)
	assert.Subset(t, result, []int{1, 2, 3})
}

func TestSlice_Zip(t *testing.T) {
	// equal length slices
	result := slicex.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
	assert.Len(t, result, 3)
	assert.Equal(t, slicex.Pair[int, string]{Left: 1, Right: "a"}, result[0])
	assert.Equal(t, slicex.Pair[int, string]{Left: 2, Right: "b"}, result[1])
	assert.Equal(t, slicex.Pair[int, string]{Left: 3, Right: "c"}, result[2])

	// left shorter than right — result is capped at left length
	result = slicex.Zip([]int{1, 2}, []string{"a", "b", "c"})
	assert.Len(t, result, 2)
	assert.Equal(t, slicex.Pair[int, string]{Left: 1, Right: "a"}, result[0])
	assert.Equal(t, slicex.Pair[int, string]{Left: 2, Right: "b"}, result[1])

	// right shorter than left — result is capped at right length
	result = slicex.Zip([]int{1, 2, 3}, []string{"a"})
	assert.Len(t, result, 1)
	assert.Equal(t, slicex.Pair[int, string]{Left: 1, Right: "a"}, result[0])

	// empty left
	result = slicex.Zip([]int{}, []string{"a", "b"})
	assert.Empty(t, result)

	// empty right
	result = slicex.Zip([]int{1, 2}, []string{})
	assert.Empty(t, result)

	// both empty
	result = slicex.Zip([]int{}, []string{})
	assert.Empty(t, result)
}

func TestSlice_Window(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	// standard window
	result := slicex.Window(input, 3)
	assert.Len(t, result, 3)
	assert.Equal(t, []int{1, 2, 3}, result[0])
	assert.Equal(t, []int{2, 3, 4}, result[1])
	assert.Equal(t, []int{3, 4, 5}, result[2])

	// window of 1 — one element per window
	result = slicex.Window(input, 1)
	assert.Len(t, result, 5)

	for i, w := range result {
		assert.Equal(t, []int{input[i]}, w)
	}

	// window equal to slice length — single window containing all elements
	result = slicex.Window(input, len(input))
	assert.Len(t, result, 1)
	assert.Equal(t, input, result[0])

	// window larger than slice — returns empty
	result = slicex.Window(input, len(input)+1)
	assert.Empty(t, result)

	// size less than 1 — returns empty
	result = slicex.Window(input, 0)
	assert.Empty(t, result)

	result = slicex.Window(input, -1)
	assert.Empty(t, result)

	// empty input
	result = slicex.Window([]int{}, 3)
	assert.Empty(t, result)
}

func TestSlice_Rotate(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	// positive rotation — shift left
	result := slicex.Rotate(input, 2)
	assert.Equal(t, []int{3, 4, 5, 1, 2}, result)

	// single step
	result = slicex.Rotate(input, 1)
	assert.Equal(t, []int{2, 3, 4, 5, 1}, result)

	// negative rotation — shift right
	result = slicex.Rotate(input, -1)
	assert.Equal(t, []int{5, 1, 2, 3, 4}, result)

	result = slicex.Rotate(input, -2)
	assert.Equal(t, []int{4, 5, 1, 2, 3}, result)

	// zero rotation — identical copy
	result = slicex.Rotate(input, 0)
	assert.Equal(t, input, result)

	// n equal to slice length — full wrap, identical copy
	result = slicex.Rotate(input, len(input))
	assert.Equal(t, input, result)

	// n larger than slice length — normalised correctly
	result = slicex.Rotate(input, len(input)+2)
	assert.Equal(t, slicex.Rotate(input, 2), result)

	// original is not modified
	_ = slicex.Rotate(input, 2)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, input)

	// empty slice
	result = slicex.Rotate([]int{}, 3)
	assert.Empty(t, result)

	// single element
	result = slicex.Rotate([]int{42}, 1)
	assert.Equal(t, []int{42}, result)
}

func TestSlice_Chunk(t *testing.T) {
	// empty slice
	assert.Empty(t, slicex.Chunk([]int{}, 3))

	// evenly divisible
	result := slicex.Chunk([]int{1, 2, 3, 4, 5, 6}, 2)
	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, result)

	// remainder chunk
	result = slicex.Chunk([]int{1, 2, 3, 4, 5}, 2)
	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, result)

	// chunk size larger than slice
	result = slicex.Chunk([]int{1, 2, 3}, 10)
	assert.Equal(t, [][]int{{1, 2, 3}}, result)

	// chunk size of 1
	result = slicex.Chunk([]int{1, 2, 3}, 1)
	assert.Equal(t, [][]int{{1}, {2}, {3}}, result)

	// panics on invalid size
	assert.Panics(t, func() { slicex.Chunk([]int{1, 2, 3}, 0) })
}

func TestSlice_Flatten(t *testing.T) {
	assert.Empty(t, slicex.Flatten([][]int{}))

	result := slicex.Flatten([][]int{{1, 2}, {3, 4}, {5}})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)

	// nil inner slices are skipped
	result = slicex.Flatten([][]int{{1, 2}, nil, {3}})
	assert.Equal(t, []int{1, 2, 3}, result)

	// single inner slice
	result = slicex.Flatten([][]int{{1, 2, 3}})
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestSlice_IndexOf(t *testing.T) {
	assert.Equal(t, -1, slicex.IndexOf([]int{}, 1))
	assert.Equal(t, -1, slicex.IndexOf([]int{1, 2, 3}, 99))
	assert.Equal(t, 0, slicex.IndexOf([]int{1, 2, 3}, 1))
	assert.Equal(t, 2, slicex.IndexOf([]int{1, 2, 3}, 3))

	// returns first match on duplicates
	assert.Equal(t, 1, slicex.IndexOf([]int{0, 5, 5, 5}, 5))
}

func TestSlice_IndexBy(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	assert.Equal(t, -1, slicex.IndexBy([]int{}, isEven))
	assert.Equal(t, -1, slicex.IndexBy([]int{1, 3, 5}, isEven))
	assert.Equal(t, 1, slicex.IndexBy([]int{1, 2, 4, 6}, isEven))
}

func Test_None(t *testing.T) {
	assert.True(t, slicex.None([]int{1, 3, 5}, 2))
	assert.False(t, slicex.None([]int{1, 2, 3}, 2))
	assert.True(t, slicex.None([]int{}, 1))
}

func Test_NoneBy(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	assert.True(t, slicex.NoneBy([]int{1, 3, 5}, isEven))
	assert.False(t, slicex.NoneBy([]int{1, 2, 3}, isEven))
	assert.True(t, slicex.NoneBy([]int{}, isEven))
}

func Test_Count(t *testing.T) {
	assert.Equal(t, 3, slicex.Count([]int{1, 2, 2, 3, 2}, 2))
	assert.Equal(t, 0, slicex.Count([]int{1, 3, 5}, 2))
	assert.Equal(t, 0, slicex.Count([]int{}, 1))
}

func Test_CountBy(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	assert.Equal(t, 3, slicex.CountBy([]int{1, 2, 4, 3, 6}, isEven))
	assert.Equal(t, 0, slicex.CountBy([]int{1, 3, 5}, isEven))
	assert.Equal(t, 0, slicex.CountBy([]int{}, isEven))
}

func Test_SortBy(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5, 9, 2}
	sorted := slicex.SortBy(nums, cmp.Compare)

	assert.Equal(t, []int{1, 1, 2, 3, 4, 5, 9}, sorted)
	// original must not be modified
	assert.Equal(t, []int{3, 1, 4, 1, 5, 9, 2}, nums)
}

func Test_Sort(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5}
	sorted := slicex.Sort(nums)

	assert.Equal(t, []int{1, 1, 3, 4, 5}, sorted)
	assert.Equal(t, []int{3, 1, 4, 1, 5}, nums)
}
