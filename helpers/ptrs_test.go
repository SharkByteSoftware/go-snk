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

func TestIsEmpty(t *testing.T) {
	intValue := 1
	assert.False(t, helpers.IsEmpty(intValue))

	var emptyInt int
	assert.True(t, helpers.IsEmpty(emptyInt))

	strValue := "test"
	assert.False(t, helpers.IsEmpty(strValue))

	var emptyStr string
	assert.True(t, helpers.IsEmpty(emptyStr))

	testValue := testType{name: "test", num: 1}
	assert.False(t, helpers.IsEmpty(testValue))

	testValue = testType{name: "test"}
	assert.False(t, helpers.IsEmpty(testValue))

	testValue = testType{}
	assert.True(t, helpers.IsEmpty(testValue))
}

func TestNil(t *testing.T) {
	intPtr := helpers.Nil[int]()
	assert.Nil(t, intPtr)
	assert.IsType(t, (*int)(nil), helpers.Nil[int]())

	strPtr := helpers.Nil[string]()
	assert.Nil(t, strPtr)
	assert.IsType(t, (*string)(nil), helpers.Nil[string]())

	testTypePtr := helpers.Nil[testType]()
	assert.Nil(t, testTypePtr)
	assert.IsType(t, (*testType)(nil), helpers.Nil[testType]())
}

func TestAsPtr(t *testing.T) {
	intPtr := helpers.AsPtr(5)
	assert.IsType(t, (*int)(nil), intPtr)
	assert.Equal(t, 5, *intPtr)

	strPtr := helpers.AsPtr("hello")
	assert.IsType(t, (*string)(nil), strPtr)
	assert.Equal(t, "hello", *strPtr)

	testTypePtr := helpers.AsPtr(testType{name: "test", num: 5})
	assert.IsType(t, (*testType)(nil), testTypePtr)
	assert.Equal(t, "test", testTypePtr.name)
	assert.Equal(t, testTypePtr.num, 5)
}

func TestAsValue(t *testing.T) {
	ptr := &testType{"test", 10}

	result := helpers.AsValue(ptr)
	assert.Equal(t, *ptr, result)

	ptr = nil
	result = helpers.AsValue(ptr)
	assert.Equal(t, helpers.Empty[testType](), result)

	assert.Equal(t, 0, helpers.AsValue[int](nil))
	assert.Equal(t, "", helpers.AsValue[string](nil))
}

func TestAsValueOr(t *testing.T) {
	value := &testType{"test", 10}
	fallback := testType{"fallback", 20}

	result := helpers.AsValueOr(value, fallback)
	assert.Equal(t, *value, result)

	value = nil
	result = helpers.AsValueOr(value, fallback)
	assert.Equal(t, fallback, result)
}
