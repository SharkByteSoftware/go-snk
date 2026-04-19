package parallel_test

import (
	"fmt"
	"slices"
	"sync/atomic"

	"github.com/SharkByteSoftware/go-snk/slicex/parallel"
)

func ExampleMap() {
	numbers := []int{1, 2, 3, 4, 5}

	squares := parallel.Map(numbers, func(n int) int { return n * n })

	fmt.Println(squares)
	// Output: [1 4 9 16 25]
}

func ExampleMapWithLimit() {
	numbers := []int{1, 2, 3, 4, 5}

	squares := parallel.MapWithLimit(numbers, func(n int) int { return n * n }, 2)

	fmt.Println(squares)
	// Output: [1 4 9 16 25]
}

func ExampleApply() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var total atomic.Int64

	parallel.Apply(numbers, func(n int) {
		total.Add(int64(n))
	})

	fmt.Println(total.Load())
	// Output: 21
}

func ExampleApplyWithLimit() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var total atomic.Int64

	parallel.ApplyWithLimit(numbers, func(n int) {
		total.Add(int64(n))
	}, 2)

	fmt.Println(total.Load())
	// Output: 21
}

func ExampleGroupBy() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	groups := parallel.GroupBy(numbers, func(n int) string {
		if n%2 == 0 {
			return "even"
		}

		return "odd"
	})

	evens := groups["even"]
	odds := groups["odd"]

	slices.Sort(evens)
	slices.Sort(odds)

	fmt.Println(evens)
	fmt.Println(odds)
	// Output:
	// [2 4 6]
	// [1 3 5]
}

func ExampleGroupByWithLimit() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	groups := parallel.GroupByWithLimit(numbers, func(n int) string {
		if n%2 == 0 {
			return "even"
		}

		return "odd"
	}, 2)

	evens := groups["even"]
	odds := groups["odd"]

	slices.Sort(evens)
	slices.Sort(odds)

	fmt.Println(evens)
	fmt.Println(odds)
	// Output:
	// [2 4 6]
	// [1 3 5]
}

func ExamplePartition() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	even, odd := parallel.Partition(numbers, func(n int) bool { return n%2 == 0 })

	slices.Sort(even)
	slices.Sort(odd)

	fmt.Println(even)
	fmt.Println(odd)
	// Output:
	// [2 4 6]
	// [1 3 5]
}

func ExamplePartitionWithLimit() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	even, odd := parallel.PartitionWithLimit(numbers, func(n int) bool { return n%2 == 0 }, 2)

	slices.Sort(even)
	slices.Sort(odd)

	fmt.Println(even)
	fmt.Println(odd)
	// Output:
	// [2 4 6]
	// [1 3 5]
}
