package ds_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	list := ds.NewList[int]()
	assert.Equal(t, 0, list.Len())
	assert.Nil(t, list.Front())
	assert.Nil(t, list.Back())

	list = ds.NewList(1)
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, 1, list.Front().Value)
	assert.Equal(t, 1, list.Back().Value)

	list = ds.NewList(1, 2, 4, 6)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, 1, list.Front().Value)
	assert.Equal(t, 6, list.Back().Value)
}

func TestList_PushFront(t *testing.T) {
	list := ds.NewList[int](1)
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, 1, list.Front().Value)
	assert.Equal(t, 1, list.Back().Value)

	list.PushFront(2)
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, 2, list.Front().Value)
	assert.Equal(t, 1, list.Back().Value)

	list.PushFront(12, 33, 2, 4, 6)
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, 12, list.Front().Value)
	assert.Equal(t, 1, list.Back().Value)
}

func TestList_PushBack(t *testing.T) {
	list := ds.NewList[int](1)
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, 1, list.Front().Value)
	assert.Equal(t, 1, list.Back().Value)

	list.PushBack(2)
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, 1, list.Front().Value)
	assert.Equal(t, 2, list.Back().Value)

	list.PushBack(12, 33, 2, 4, 6)
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, 1, list.Front().Value)
	assert.Equal(t, 6, list.Back().Value)
}

func TestList_IsEmpty_Len(t *testing.T) {
	list := ds.NewList[int]()
	assert.True(t, list.IsEmpty())
	assert.Equal(t, 0, list.Len())

	list = ds.NewList(1, 2, 4, 6)
	assert.False(t, list.IsEmpty())
	assert.Equal(t, 4, list.Len())
}
