package helpers_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/helpers"
	"github.com/stretchr/testify/assert"
)

type testType struct {
	name string
	num  int
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, 0, helpers.Empty[int]())
	assert.Equal(t, "", helpers.Empty[string]())
	assert.Equal(t, testType{}, helpers.Empty[testType]())
}

func TestSafeDeref(t *testing.T) {
	value := &testType{"test", 10}

	result := helpers.SafeDeref(value)
	assert.Equal(t, *value, result)

	value = nil
	result = helpers.SafeDeref(value)
	assert.Equal(t, helpers.Empty[testType](), result)

	assert.Equal(t, 0, helpers.SafeDeref[int](nil))
	assert.Equal(t, "", helpers.SafeDeref[string](nil))
}

func TestSafeDerefOrl(t *testing.T) {
	value := &testType{"test", 10}
	fallback := testType{"fallback", 20}

	result := helpers.SafeDerefOr(value, fallback)
	assert.Equal(t, *value, result)

	value = nil
	result = helpers.SafeDerefOr(value, fallback)
	assert.Equal(t, fallback, result)
}
