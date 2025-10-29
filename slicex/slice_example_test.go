package slicex_test

import (
	"fmt"

	"github.com/SharkByteSoftware/go-snk/slicex"
)

func ExampleFilter() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	values := slicex.Filter(numbers, func(item int) bool { return item%2 == 0 })

	fmt.Println(values)
	// Output: [2 4 6]
}

func ExampleFilterWithIndex() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	values := slicex.FilterWithIndex(numbers, func(_ int, idx int) bool {
		return numbers[idx]%3 == 0
	})

	fmt.Println(values)
	// Output: [3 6]
}

func ExampleMap() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	values := slicex.Map(numbers, func(item int) string { return fmt.Sprintf("'%d'", item) })

	fmt.Println(values)
	// Output: ['1' '2' '3' '4' '5' '6']
}

func ExampleMapWithIndex() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	values := slicex.MapWithIndex(numbers, func(_ int, idx int) string {
		return fmt.Sprintf("'%d'", numbers[idx])
	})

	fmt.Println(values)
	// Output: ['1' '2' '3' '4' '5' '6']
}

func ExampleUnique() {
	numbers := []int{1, 1, 2, 2, 2, 5, 5, 5, 5}

	values := slicex.Unique(numbers)

	fmt.Println(values)
	// Output: [1 2 5]
}

func ExampleUniqueMap() {
	numbers := []int{1, 1, 2, 2, 2, 5, 5, 5, 5}

	values := slicex.UniqueMap(numbers, func(item int) string { return fmt.Sprintf("'%d'", item) })

	fmt.Println(values)
	// Output: ['1' '2' '5']
}

func ExampleBind() {
	nestedNumbers := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	values := slicex.Bind(nestedNumbers, func(item []int) []int { return item })

	fmt.Println(values)
	// Output: [1 2 3 4 5 6 7 8 9]
}

func ExampleReduce() {
	nestedNumbers := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	acc := func(agg int, items []int) int {
		return agg * slicex.Sum(items)
	}

	result := slicex.Reduce(nestedNumbers, acc, 1)

	fmt.Println(result)
	// Output: 2160
}

func ExampleFind() {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "not"}

	value, found := slicex.Find(numbers, "four")

	fmt.Println(value, found)
	// Output: four true
}

func ExampleFindBy() {
	numbers := []string{"one", "two", "not"}

	value, found := slicex.FindBy(numbers, func(item string) bool {
		return item != "one" && item != "two"
	})

	fmt.Println(value, found)
	// Output: not true
}

func ExampleFindOr() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	value1 := slicex.FindOr(numbers, 22, 256)
	value2 := slicex.FindOr(numbers, 6, 256)

	fmt.Println(value1, value2)
	// Output: 256 6
}

func ExampleFindOrBy() {
	numbers := []string{"one", "two", "not"}

	value1 := slicex.FindOrBy(numbers, func(item string) bool { return item == "not" },
		"nothing")
	value2 := slicex.FindOrBy(numbers, func(item string) bool { return item == "other" },
		"nothing")

	fmt.Println(value1, value2)
	// Output: not nothing
}

func ExampleAny() {
	numbers := []int{2, 2, 2, 4, 5, 6}

	result1 := slicex.Any(numbers, 2)
	result2 := slicex.Any(numbers, 22)

	fmt.Println(result1, result2)
	// Output: true false
}

func ExampleAnyBy() {
	numbers := []int{2, 2, 2, 4, 12, 6}

	result1 := slicex.AnyBy(numbers, func(item int) bool { return item%2 == 0 })
	result2 := slicex.AnyBy(numbers, func(item int) bool { return item%2 != 0 })

	fmt.Println(result1, result2)
	// Output: true false
}

func ExampleAll() {
	numbers1 := []int{2, 2, 2}
	numbers2 := []int{2, 2, 5}

	result1 := slicex.All(numbers1, 2)
	result2 := slicex.All(numbers2, 2)

	fmt.Println(result1, result2)
	// Output: true false
}

func ExampleApply() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var sum int

	slicex.Apply(numbers, func(item int) {
		sum += item
	})

	fmt.Println(sum)
	// Output: 21
}

func ExampleApplyWithIndex() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var sum int

	slicex.ApplyWithIndex(numbers, func(_ int, index int) {
		sum += numbers[index]
		numbers[index] = sum
	})

	fmt.Println(sum, numbers)
	// Output: 21 [1 3 6 10 15 21]
}

func ExampleReverse() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	values := slicex.Reverse(numbers)

	fmt.Println(values)
	// Output: [6 5 4 3 2 1]
}

func ExampleCompact() {
	numbers := []int{0, 2, 3, 4, 5, 0}

	values := slicex.Compact(numbers)
	fmt.Println(values)
	// Output: [2 3 4 5]
}

func ExampleToMap() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	result := slicex.ToMap(numbers, func(item int) string { return fmt.Sprintf("%d~key", item) })

	fmt.Println(result)
	// Output: map[1~key:1 2~key:2 3~key:3 4~key:4 5~key:5 6~key:6]
}

func ExampleGroupBy() {
	numbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	result := slicex.GroupBy(numbers, func(item int) int {
		return item % 4
	})

	fmt.Println(result)
	// Output: map[0:[4 8] 1:[1 5 9] 2:[2 6 10] 3:[3]]
}

func ExamplePartition() {
	numbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	even, odd := slicex.Partition(numbers, func(item int) bool { return item%2 == 0 })

	fmt.Println(even, odd)
	// Output: [2 4 6 8 10] [1 3 5 9]
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	sum := slicex.Sum(numbers)

	fmt.Println(sum)
	// Output: 48
}

func ExampleProduct() {
	numbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	product := slicex.Product(numbers)

	fmt.Println(product)
	// Output: 518400
}

func ExampleMean() {
	numbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	mean := slicex.Mean(numbers)

	fmt.Println(mean)
	// Output: 5
}

func ExampleMax() {
	numbers := []int{1, 30, 3, 4, 5, 6, -1, 9, 10}

	maximum := slicex.Max(numbers)

	fmt.Println(maximum)
	// Output: 30
}

func ExampleMaxBy() {
	numbers := []int{1, 30, 3, 4, 5, 6, -1, 9, 10}

	maximum := slicex.MaxBy(numbers, func(a int, b int) bool { return a < b })

	fmt.Println(maximum)
	// Output: 30
}

func ExampleMin() {
	numbers := []int{1, 30, 3, 4, 5, 6, -1, 9, 10}

	minimum := slicex.Min(numbers)

	fmt.Println(minimum)
	// Output: -1
}
