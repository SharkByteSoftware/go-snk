package helpers_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/helpers"
	"github.com/stretchr/testify/assert"
)

type valueTestType struct {
	name string
	num  int
}

func TestEmpty(t *testing.T) {
	assert.Empty(t, helpers.Empty[int]())
	assert.Empty(t, helpers.Empty[string]())
	assert.Equal(t, ptrTestType{}, helpers.Empty[ptrTestType]())
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

	testValue := valueTestType{name: "test", num: 1}
	assert.False(t, helpers.IsEmpty(testValue))

	testValue = valueTestType{name: "test"}
	assert.False(t, helpers.IsEmpty(testValue))

	testValue = valueTestType{}
	assert.True(t, helpers.IsEmpty(testValue))
}
