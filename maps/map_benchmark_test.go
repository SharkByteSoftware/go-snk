package maps_test

import (
	"fmt"
	"testing"

	"github.com/SharkByteSoftware/go-snk/maps"
	"github.com/SharkByteSoftware/go-snk/slices"
)

var startingSize = []int{0, 1, 100, 1000, 10000}

//func BenchmarkKeys(b *testing.B) {
//	for _, size := range startingSize {
//		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
//			intMap := generateMap(size)
//			for b.Loop() {
//				_ = maps.Keys(intMap)
//			}
//		})
//	}
//}

func BenchmarkValues(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			intMap := generateMap(size)
			for b.Loop() {
				_ = maps.Values(intMap)
			}
		})
	}
}

func generateMap(size int) map[int]int {
	return slices.ToMap(generateIntSlice(size), slices.ValueAdapter[int]())
}

func generateIntSlice(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = i
	}

	return result
}
