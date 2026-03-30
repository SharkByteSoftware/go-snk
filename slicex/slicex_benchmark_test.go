package slicex_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/SharkByteSoftware/go-snk/containers/sets"
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
				_ = slicex.Bind(ints, func(item []int) []int { return item })
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
		for range b.N {
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

func BenchmarkZip(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			left := generateIntSlice(size)
			right := generateIntSlice(size)

			for b.Loop() {
				_ = slicex.Zip(left, right)
			}
		})
	}
}

func BenchmarkWindow(b *testing.B) {
	windowSizes := []int{2, 10, 100}

	for _, size := range startingSize {
		for _, windowSize := range windowSizes {
			if windowSize > size {
				continue
			}

			b.Run(fmt.Sprintf("slice size: %d window size: %d", size, windowSize), func(b *testing.B) {
				ints := generateIntSlice(size)
				for b.Loop() {
					_ = slicex.Window(ints, windowSize)
				}
			})
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Rotate(ints, size/2)
			}
		})
	}
}

func BenchmarkDifference(b *testing.B) {
	otherSizes := []int{0, 10, 100, 1000}

	for _, size := range startingSize {
		for _, otherSize := range otherSizes {
			b.Run(fmt.Sprintf("len: %d other: %d", size, otherSize), func(b *testing.B) {
				set := sets.New(generateIntSlice(size)...)
				other := sets.New(generateIntSlice(otherSize)...)

				for b.Loop() {
					_ = set.Difference(other)
				}
			})
		}
	}
}

func BenchmarkPartition(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_, _ = slicex.Partition(ints, func(i int) bool { return i%2 == 0 })
			}
		})
	}
}

func BenchmarkPartition_SplitRatio(b *testing.B) {
	const size = 10000

	ratios := []struct {
		name      string
		predicate func(int, int) bool
	}{
		{"50/50", func(i, idx int) bool { return idx%2 == 0 }},
		{"90/10", func(i, idx int) bool { return idx%10 != 0 }},
		{"all true", func(i, idx int) bool { return true }},
		{"all false", func(i, idx int) bool { return false }},
	}

	ints := generateIntSlice(size)

	for _, ratio := range ratios {
		b.Run(ratio.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = slicex.Partition(ints, func(i int) bool { return ratio.predicate(i, 0) })
			}
		})
	}
}
