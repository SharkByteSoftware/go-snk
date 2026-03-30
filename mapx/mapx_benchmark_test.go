package mapx_test

import (
	"fmt"
	"strconv"
	"testing"

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

func BenchmarkMapKeys(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			intMap := generateMap(size)
			for b.Loop() {
				_ = mapx.MapKeys(intMap, strconv.Itoa)
			}
		})
	}
}

func BenchmarkPartition(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			intMap := generateMap(size)
			for b.Loop() {
				_, _ = mapx.Partition(intMap, func(k int, _ int) bool { return k%2 == 0 })
			}
		})
	}
}

func BenchmarkCountBy(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			intMap := generateMap(size)
			for b.Loop() {
				_ = mapx.CountBy(intMap, func(k int, _ int) string { return strconv.Itoa(k % 10) })
			}
		})
	}
}

func generateMap(size int) map[int]int {
	return slicex.ToMap(generateIntSlice(size), func(k int) int { return k })
}

func generateIntSlice(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = i
	}

	return result
}
