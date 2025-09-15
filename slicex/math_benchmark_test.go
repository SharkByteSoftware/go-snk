package slicex_test

import (
	"fmt"
	"testing"

	"github.com/SharkByteSoftware/go-snk/slicex"
)

func BenchmarkSum(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Sum(ints)
			}
		})
	}
}

func BenchmarkProduct(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Product(ints)
			}
		})
	}
}

func BenchmarkMean(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Mean(ints)
			}
		})
	}
}

func BenchmarkMax(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Max(ints)
			}
		})
	}
}

func BenchmarkMin(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("slice size: %d", size), func(b *testing.B) {
			ints := generateIntSlice(size)
			for b.Loop() {
				_ = slicex.Min(ints)
			}
		})
	}
}
