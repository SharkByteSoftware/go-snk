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
