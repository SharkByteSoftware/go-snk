package mapx_test

import (
	"fmt"
	"testing"

	"github.com/SharkByteSoftware/go-snk/internal/adapt"
	"github.com/SharkByteSoftware/go-snk/mapx"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

var startingSize = []int{0, 1, 100, 1000, 10000}

func BenchmarkKeys(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			intMap := generateMap(size)
			for b.Loop() {
				_ = mapx.Keys(intMap)
			}
		})
	}
}

func BenchmarkValues(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			intMap := generateMap(size)
			for b.Loop() {
				_ = mapx.Values(intMap)
			}
		})
	}
}

func BenchmarkCombine(b *testing.B) {
}

func generateMap(size int) map[int]int {
	return slicex.ToMap(generateIntSlice(size), adapt.ValueAdapter)
}

func generateIntSlice(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = i
	}

	return result
}
