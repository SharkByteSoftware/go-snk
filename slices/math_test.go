package slices_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/slices"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	intResult := slices.Sum([]int{1, 2, 3})
	assert.Equal(t, 6, intResult)

	floatResult := slices.Sum([]float32{1.0, 2.0, 3.0})
	assert.Equal(t, float32(6.0), floatResult)

	float64Result := slices.Sum([]float64{1.0, 2.0, 3.0})
	assert.Equal(t, 6.0, float64Result)
}

func TestProduct(t *testing.T) {
	product := slices.Product(numberList)
	assert.Equal(t, 10229760, product)

	product = slices.Product(duplicateList)
	assert.Equal(t, 104647989657600, product)

	product = slices.Product([]int{1, 2, -3})
	assert.Equal(t, -6, product)

	product = slices.Product([]int{})
	assert.Equal(t, 1, product)
}

func TestMean(t *testing.T) {
	mean := slices.Mean(numberList)
	assert.Equal(t, 86, mean)

	mean = slices.Mean(duplicateList)
	assert.Equal(t, 86, mean)

	mean = slices.Mean([]int{1, 2, 10, -3})
	assert.Equal(t, 2, mean)

	mean = slices.Mean([]int{-1, -2, -3, -123})
	assert.Equal(t, -32, mean)

	mean = slices.Mean([]int{})
	assert.Equal(t, 0, mean)
}

func TestMax(t *testing.T) {
	max := slices.Max(numberList)
	assert.Equal(t, 333, max)

	max = slices.Max(duplicateList)
	assert.Equal(t, 333, max)

	max = slices.Max([]int{})
	assert.Equal(t, 0, max)
}

func TestMin(t *testing.T) {
	min := slices.Min(numberList)
	assert.Equal(t, 1, min)

	min = slices.Min(duplicateList)
	assert.Equal(t, 1, min)

	min = slices.Min(slices.Reverse(numberList))
	assert.Equal(t, 1, min)

	min = slices.Min([]int{})
	assert.Equal(t, 0, min)
}
