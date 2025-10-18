package parallel_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/SharkByteSoftware/go-snk/slicex/parallel"
)

const (
	sliceCount = 5
	maxRandInt = 1000
	sliceSize  = 10000
)

var startingSize = []int{1, 10, 100, 1000, 10000}

func BenchmarkMap(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = parallel.Map(ints, func(i int) bool { return i%2 == 0 })
			}
		})
	}
}

func BenchmarkMapWithLimit(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("concurrency: %d", size), func(b *testing.B) {
			ints := generateIntSlice(sliceSize)
			for b.Loop() {
				_ = parallel.MapWithLimit(ints,
					func(n int) string {
						time.Sleep(500 * time.Nanosecond)
						return strconv.Itoa(n)
					},
					size)
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
