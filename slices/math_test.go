package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	intResult := Sum([]int{1, 2, 3})
	assert.Equal(t, 6, intResult)

	floatResult := Sum([]float32{1.0, 2.0, 3.0})
	assert.Equal(t, float32(6.0), floatResult)

	float64Result := Sum([]float64{1.0, 2.0, 3.0})
	assert.Equal(t, 6.0, float64Result)

	complexResult := Sum([]complex128{1.0, 2.0, 3.0})
	assert.Equal(t, complex(6.0, 0.0), complexResult)
}

func TestMax(t *testing.T) {
}

func TestMaxBy(t *testing.T) {
}

func TestMean(t *testing.T) {
}

func TestMeanBy(t *testing.T) {
}

func TestMin(t *testing.T) {
}

func TestMinBy(t *testing.T) {
}

func TestProduct(t *testing.T) {
}

func TestProductBy(t *testing.T) {
}
