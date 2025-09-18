package ds_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	list := ds.NewList[int]()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.First())
	assert.Nil(t, list.Last())

	list = ds.NewList(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.First().Value)
	assert.Equal(t, 1, list.Last().Value)

	list = ds.NewList(1, 2, 4, 6)
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 1, list.First().Value)
	assert.Equal(t, 6, list.Last().Value)
}

func TestList_Add(t *testing.T) {
	list := ds.NewList[int](1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.First().Value)
	assert.Equal(t, 1, list.Last().Value)

	list.Add(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.First().Value)
	assert.Equal(t, 2, list.Last().Value)

	list.Add(1, 2, 4, 6)
	assert.Equal(t, 6, list.Size())
	assert.Equal(t, 1, list.First().Value)
	assert.Equal(t, 6, list.Last().Value)
}

func TestList_IsEmpty_Size(t *testing.T) {
	list := ds.NewList[int]()
	assert.True(t, list.IsEmpty())
	assert.Equal(t, 0, list.Size())

	list = ds.NewList(1, 2, 4, 6)
	assert.False(t, list.IsEmpty())
	assert.Equal(t, 4, list.Size())
}
