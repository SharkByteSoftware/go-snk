package slices_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/SharkByteSoftware/go-snk/slices"
)

func generateIntSlice(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = rand.Intn(100_000)
	}

	return result
}

func BenchmarkFilter(b *testing.B) {
	size := 1000
	ints := generateIntSlice(size)

	b.Run(fmt.Sprintf("ints_%d", size), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Filter(ints, func(i int) bool { return i%2 == 0 })
		}
	})
}
