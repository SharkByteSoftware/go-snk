package slices_test

import (
	"fmt"
	"testing"

	"github.com/SharkByteSoftware/go-snk/slices"
)

func BenchmarkSum(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slices.Sum(ints)
			}
		})
	}
}

func BenchmarkProduct(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slices.Product(ints)
			}
		})
	}
}

func BenchmarkMean(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slices.Mean(ints)
			}
		})
	}
}

func BenchmarkMax(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slices.Max(ints)
			}
		})
	}
}

func BenchmarkMin(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slices.Min(ints)
			}
		})
	}
}
