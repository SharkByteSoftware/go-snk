package helpers_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/helpers"
	"github.com/stretchr/testify/assert"
)

type ptrTestType struct {
	name string
	num  int
}

func TestNil(t *testing.T) {
	intPtr := helpers.Nil[int]()
	assert.Nil(t, intPtr)
	assert.IsType(t, (*int)(nil), helpers.Nil[int]())

	strPtr := helpers.Nil[string]()
	assert.Nil(t, strPtr)
	assert.IsType(t, (*string)(nil), helpers.Nil[string]())

	testTypePtr := helpers.Nil[ptrTestType]()
	assert.Nil(t, testTypePtr)
	assert.IsType(t, (*ptrTestType)(nil), helpers.Nil[ptrTestType]())
}

func TestAsPtr(t *testing.T) {
	intPtr := helpers.AsPtr(5)
	assert.IsType(t, (*int)(nil), intPtr)
	assert.Equal(t, 5, *intPtr)

	strPtr := helpers.AsPtr("hello")
	assert.IsType(t, (*string)(nil), strPtr)
	assert.Equal(t, "hello", *strPtr)

	testTypePtr := helpers.AsPtr(ptrTestType{name: "test", num: 5})
	assert.IsType(t, (*ptrTestType)(nil), testTypePtr)
	assert.Equal(t, "test", testTypePtr.name)
	assert.Equal(t, testTypePtr.num, 5)
}

func TestAsValue(t *testing.T) {
	ptr := &ptrTestType{"test", 10}

	result := helpers.AsValue(ptr)
	assert.Equal(t, *ptr, result)

	ptr = nil
	result = helpers.AsValue(ptr)
	assert.Equal(t, helpers.Empty[ptrTestType](), result)

	assert.Equal(t, 0, helpers.AsValue[int](nil))
	assert.Equal(t, "", helpers.AsValue[string](nil))
}

func TestAsValueOr(t *testing.T) {
	value := &ptrTestType{"test", 10}
	fallback := ptrTestType{"fallback", 20}

	result := helpers.AsValueOr(value, fallback)
	assert.Equal(t, *value, result)

	value = nil
	result = helpers.AsValueOr(value, fallback)
	assert.Equal(t, fallback, result)
}
