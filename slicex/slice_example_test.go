package slicex_test

import (
	"fmt"

	"github.com/SharkByteSoftware/go-snk/slicex"
)

func ExampleFilter() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	values := slicex.Filter(numberList, func(item int) bool { return item%2 == 0 })

	fmt.Println(values)
	// Output: [2 4 6]
}

func ExampleFilterWithIndex() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	values := slicex.FilterWithIndex(numberList, func(_ int, idx int) bool {
		return numberList[idx]%3 == 0
	})

	fmt.Println(values)
	// Output: [3 6]
}

func ExampleMap() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	values := slicex.Map(numberList, func(item int) string { return fmt.Sprintf("'%d'", item) })

	fmt.Println(values)
	// Output: ['1' '2' '3' '4' '5' '6']
}

func ExampleMapWithIndex() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	values := slicex.MapWithIndex(numberList, func(_ int, idx int) string {
		return fmt.Sprintf("'%d'", numberList[idx])
	})

	fmt.Println(values)
	// Output: ['1' '2' '3' '4' '5' '6']
}

func ExampleUnique() {
	numberList := []int{1, 1, 2, 2, 2, 5, 5, 5, 5}

	values := slicex.Unique(numberList)

	fmt.Println(values)
	// Output [1 2 5]
}

func ExampleUniqueMap() {
	numberList := []int{1, 1, 2, 2, 2, 5, 5, 5, 5}

	values := slicex.UniqueMap(numberList, func(item int) string { return fmt.Sprintf("'%d'", item) })

	fmt.Println(values)
	// Output ['1' '2' '5']
}

func ExampleBind() {
	nestedNumberList := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	values := slicex.Bind(nestedNumberList, func(item []int) []int { return item })

	fmt.Println(values)
	// Output [1 2 3 4 5 6 7 8 9]
}

func ExampleReduce() {
	nestedNumberList := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	accumulator := func(agg int, items []int) int {
		return agg * slicex.Sum(items)
	}

	result := slicex.Reduce(nestedNumberList, accumulator, 1)

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
	numberList := []int{1, 2, 3, 4, 5, 6}

	value1 := slicex.FindOr(numberList, 22, 256)
	value2 := slicex.FindOr(numberList, 6, 256)

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
	numberList := []int{2, 2, 2, 4, 5, 6}

	result1 := slicex.Any(numberList, 2)
	result2 := slicex.Any(numberList, 22)

	fmt.Println(result1, result2)
	// Output: true false
}

func ExampleAnyBy() {
	numberList := []int{2, 2, 2, 4, 12, 6}

	result1 := slicex.AnyBy(numberList, func(item int) bool { return item%2 == 0 })
	result2 := slicex.AnyBy(numberList, func(item int) bool { return item%2 != 0 })

	fmt.Println(result1, result2)
	// Output: true false
}

func ExampleAll() {
	numberList1 := []int{2, 2, 2}
	numberList2 := []int{2, 2, 5}

	result1 := slicex.All(numberList1, 2)
	result2 := slicex.All(numberList2, 2)

	fmt.Println(result1, result2)
	// Output: true false
}

func ExampleApply() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	var sum int

	slicex.Apply(numberList, func(item int) {
		sum += item
	})

	fmt.Println(sum)
	// Output: 21
}

func ExampleApplyWithIndex() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	var sum int

	slicex.ApplyWithIndex(numberList, func(_ int, index int) {
		sum += numberList[index]
		numberList[index] = sum
	})

	fmt.Println(sum, numberList)
	// Output: 21 [1 3 6 10 15 21]
}

func ExampleReverse() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	values := slicex.Reverse(numberList)

	fmt.Println(values)
	// Output: [6 5 4 3 2 1]
}

func ExampleToMap() {
	numberList := []int{1, 2, 3, 4, 5, 6}

	result := slicex.ToMap(numberList, func(item int) string { return fmt.Sprintf("%d~key", item) })

	fmt.Println(result)
	// Output: map[1~key:1 2~key:2 3~key:3 4~key:4 5~key:5 6~key:6]
}

func ExampleGroupBy() {
	numberList := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	result := slicex.GroupBy(numberList, func(item int) int {
		return item % 4
	})

	fmt.Println(result)
	// Output: map[0:[4 8] 1:[1 5 9] 2:[2 6 10] 3:[3]]
}

func ExamplePartition() {
	numberList := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	even, odd := slicex.Partition(numberList, func(item int) bool { return item%2 == 0 })

	fmt.Println(even, odd)
	// Output: [2 4 6 8 10] [1 3 5 9]
}
