package slicex_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/SharkByteSoftware/go-snk/adapt"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

const (
	sliceSize  = 1000
	sliceCount = 5
	maxRandInt = 1000
)

var startingSize = []int{0, 1, 100, 1000, 10000}

func BenchmarkFilter(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Filter(ints, func(i int) bool { return i%2 == 0 })
			}
		})
	}
}

func BenchmarkMap(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Map(ints, func(i int) bool { return i%2 == 0 })
			}
		})
	}
}

func BenchmarkUniqueMap(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.UniqueMap(ints, func(i int) bool { return i%2 == 0 })
			}
		})
	}
}

func BenchmarkBind(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateNestedIntSlices(sliceCount, size)
			for b.Loop() {
				_ = slicex.Bind(ints, adapt.ValueAdapter)
			}
		})
	}
}

func BenchmarkReduce(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateNestedIntSlices(size, size)
			for b.Loop() {
				_ = slicex.Reduce(ints, accumulator, 0)
			}
		})
	}
}

func BenchmarkFind(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_, _ = slicex.Find(ints, rand.Int())
			}
		})
	}
}

func BenchmarkAny(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Any(ints, 32)
			}
		})
	}
}

func BenchmarkAll(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.All(ints, 32)
			}
		})
	}
}

func BenchmarkUnique(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Unique(ints)
			}
		})
	}
}

func BenchmarkGroupBy(b *testing.B) {
	b.Run(fmt.Sprintf("size: %d", sliceSize), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = generateNestedIntSlices(sliceCount, sliceSize)
		}
	})
}

func BenchmarkReverse(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Reverse(ints)
			}
		})
	}
}

func BenchmarkApply(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				var sum int
				slicex.Apply(ints, func(n int) { sum += n })
			}
		})
	}
}

func BenchmarkToMap(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.ToMap(ints, func(item int) int {
					return item
				})
			}
		})
	}
}

func accumulator(agg int, item []int) int {
	if len(item) == 0 {
		return agg
	}

	return agg + item[0]
}

func generateIntSlice(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = rand.Intn(maxRandInt)
	}

	return result
}

func generateNestedIntSlices(count int, size int) [][]int {
	result := make([][]int, count)
	for i := range result {
		result[i] = generateIntSlice(size)
	}

	return result
}
