package queues

import (
	"fmt"
	"testing"
)

var startingSize = []int{0, 1, 100, 1000, 10000}

func generateIntSlice(size int) *[]int {
	result := make([]int, size)
	for i := range result {
		result[i] = i
	}

	return &result
}

func compareInt(prev int, curr int) int {
	if prev < curr {
		return -1
	}
	if prev > curr {
		return 1
	}

	return 0
}

func BenchmarkPush(b *testing.B) {
	b.ResetTimer()

	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			for b.Loop() {
				pq := NewPriorityQueueWithDefault(compareInt)
				for i := range size {
					pq.Enqueue(i)
				}
			}
		})
	}
}

func BenchmarkPop(b *testing.B) {
	b.ResetTimer()

	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				pq := NewPriorityQueue(ints, compareInt)
				for range size {
					pq.Dequeue()
				}
			}
		})
	}
}
