package ds_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/SharkByteSoftware/go-snk/ds"
)

const (
	sliceSize  = 1000
	sliceCount = 5
	maxRandInt = 1000
)

var startingSize = []int{0, 1, 10, 100, 1000, 10000}

func BenchmarkNewSet(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			for b.Loop() {
				_ = ds.NewSet(generateIntSlice(size)...)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			set := ds.NewSet[int](generateIntSlice(size)...)
			for b.Loop() {
				set.Add(generateIntSlice(sliceSize)...)
			}
		})
	}
}

func BenchmarkContains(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			set := ds.NewSet(generateIntSlice(size)...)
			for b.Loop() {
				_ = set.Contains(rand.Int())
			}
		})
	}
}

func BenchmarkRemove(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			set := ds.NewSet(generateIntSlice(size)...)
			for b.Loop() {
				set.Remove(rand.Int())
			}
		})
	}
}

func BenchmarkSize(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			set := ds.NewSet(generateIntSlice(size)...)
			for b.Loop() {
				_ = set.Size()
			}
		})
	}
}

func BenchmarkClear(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("size: %d", size), func(b *testing.B) {
			set := ds.NewSet(generateIntSlice(size)...)
			for b.Loop() {
				set.Clear()
			}
		})
	}
}

func BenchmarkValues(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("values: %d", size), func(b *testing.B) {
			set := ds.NewSet(generateIntSlice(size)...)
			for b.Loop() {
				_ = set.Values()
			}
		})
	}
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
