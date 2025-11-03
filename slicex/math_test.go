package slicex_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/slicex"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	intResult := slicex.Sum([]int{1, 2, 3})
	assert.Equal(t, 6, intResult)

	floatResult := slicex.Sum([]float32{1.0, 2.0, 3.0})
	assert.Equal(t, float32(6.0), floatResult)

	float64Result := slicex.Sum([]float64{1.0, 2.0, 3.0})
	assert.Equal(t, 6.0, float64Result)
}

func TestSumBy(t *testing.T) {
	result := slicex.SumBy([]int{1, 2, 3}, func(i int) int { return i * i })
	assert.Equal(t, 14, result)

	result = slicex.SumBy([]int{}, func(i int) int { return i * i })
	assert.Equal(t, 0, result)
}

func TestProduct(t *testing.T) {
	product := slicex.Product(numberList)
	assert.Equal(t, 10229760, product)

	product = slicex.Product(duplicateList)
	assert.Equal(t, 104647989657600, product)

	product = slicex.Product([]int{1, 2, -3})
	assert.Equal(t, -6, product)

	product = slicex.Product([]int{})
	assert.Equal(t, 1, product)
}

func TestProductBy(t *testing.T) {
	stringList := []string{"a", "aa", "aaa"}

	result := slicex.ProductBy(stringList, func(s string) int { return len(s) })
	assert.Equal(t, 6, result)

	result = slicex.ProductBy([]string{}, func(s string) int { return len(s) })
	assert.Equal(t, 1, result)
}

func TestMean(t *testing.T) {
	mean := slicex.Mean(numberList)
	assert.Equal(t, 86, mean)

	mean = slicex.Mean(duplicateList)
	assert.Equal(t, 86, mean)

	mean = slicex.Mean([]int{1, 2, 10, -3})
	assert.Equal(t, 2, mean)

	mean = slicex.Mean([]int{-1, -2, -3, -123})
	assert.Equal(t, -32, mean)

	mean = slicex.Mean([]int{})
	assert.Equal(t, 0, mean)
}

func TestMax(t *testing.T) {
	max := slicex.Max(numberList)
	assert.Equal(t, 333, max)

	max = slicex.Max(duplicateList)
	assert.Equal(t, 333, max)

	max = slicex.Max([]int{})
	assert.Equal(t, 0, max)
}

func TestMaxBy(t *testing.T) {
	result := slicex.MaxBy(numberList, func(a int, b int) bool { return a < b })
	assert.Equal(t, 333, result)

	result = slicex.MaxBy(duplicateList, func(a int, b int) bool { return a < b })
	assert.Equal(t, 333, result)

	result = slicex.MaxBy([]int{}, func(a int, b int) bool { return a < b })
	assert.Equal(t, 0, result)
}

func TestMin(t *testing.T) {
	min := slicex.Min(numberList)
	assert.Equal(t, 1, min)

	min = slicex.Min(duplicateList)
	assert.Equal(t, 1, min)

	min = slicex.Min(slicex.Reverse(numberList))
	assert.Equal(t, 1, min)

	min = slicex.Min([]int{})
	assert.Equal(t, 0, min)
}
