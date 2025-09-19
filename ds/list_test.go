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

func TestList_IsEmpty_Len(t *testing.T) {
	list := ds.NewList[int]()
	assert.True(t, list.IsEmpty())
	assert.Equal(t, 0, list.Len())

	list = ds.NewList(1, 2, 4, 6)
	assert.False(t, list.IsEmpty())
	assert.Equal(t, 4, list.Len())
}

func TestList_Remove(t *testing.T) {
	list := ds.NewList(1, 556, 2, 3, 223, 5)

	value := list.Remove(list.Front())
	assert.Equal(t, 1, value)
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 556, list.Front().Value)
	assert.Equal(t, 5, list.Back().Value)

	value = list.Remove(list.Back())
	assert.Equal(t, 5, value)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, 223, list.Back().Value)

	value = list.Remove(list.Front().Next())
	assert.Equal(t, 2, value)
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, 556, list.Front().Value)
	assert.Equal(t, 223, list.Back().Value)

	list = ds.NewList(1, 2, 4, 6)
	value = list.Remove(ds.NewElement(1, nil))
	assert.Equal(t, 1, value)
	assert.Equal(t, 4, list.Len())

	list = ds.NewList(1)
	element := list.Front()
	_ = list.Remove(element)
	_ = list.Remove(element)

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

func TestList_InsertBefore(t *testing.T) {
	list := ds.NewList[int](223, 556)

	element := list.InsertBefore(1, list.Front())
	assert.Equal(t, 1, element.Value)
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, 1, list.Front().Value)

	element = list.InsertBefore(2, list.Back())
	assert.Equal(t, 2, element.Value)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, 2, list.Back().Prev().Value)

	element = list.InsertBefore(308, list.Front().Next())
	assert.Equal(t, 308, element.Value)
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 308, list.Front().Next().Value)

	list = ds.NewList(1, 2, 4, 6)
	element = list.InsertBefore(22, ds.NewElement(1, nil))
	assert.Nil(t, element)
	assert.Equal(t, 4, list.Len())
}

func TestList_InsertAfter(t *testing.T) {
	list := ds.NewList[int](223, 556)

	element := list.InsertAfter(1, list.Front())
	assert.Equal(t, 1, element.Value)
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, 1, list.Front().Next().Value)

	element = list.InsertAfter(2, list.Back())
	assert.Equal(t, 2, element.Value)
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, 2, list.Back().Value)

	element = list.InsertAfter(308, list.Front().Next())
	assert.Equal(t, 308, element.Value)
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 308, list.Front().Next().Next().Value)

	list = ds.NewList(1, 2, 4, 6)
	element = list.InsertAfter(22, ds.NewElement(1, nil))
	assert.Nil(t, element)
	assert.Equal(t, 4, list.Len())
}

func TestList_MoveToFront(t *testing.T) {
	list := ds.NewList[int](1, 223, 3, 4, 556)

	list.MoveToFront(list.Front())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 1, list.Front().Value)

	list.MoveToFront(list.Back())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 556, list.Front().Value)

	list.MoveToFront(list.Front().Next().Next())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 223, list.Front().Value)

	list.MoveToFront(ds.NewElement(1, nil))
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 223, list.Front().Value)
}

func TestList_MoveToBack(t *testing.T) {
	list := ds.NewList[int](1, 223, 3, 4, 556)

	list.MoveToBack(list.Back())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 556, list.Back().Value)

	list.MoveToBack(list.Front())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 1, list.Back().Value)

	list.MoveToBack(list.Back().Prev())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 556, list.Back().Value)

	list.MoveToBack(ds.NewElement(1, nil))
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 556, list.Back().Value)
}

func TestList_MoveBefore(t *testing.T) {
	list := ds.NewList(1, 223, 3, 4, 556)

	list.MoveBefore(list.Front(), list.Front())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 1, list.Front().Value)
	assert.Equal(t, 556, list.Back().Value)

	list.MoveBefore(list.Back(), list.Front())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 556, list.Front().Value)
	assert.Equal(t, 4, list.Back().Value)

	list = ds.NewList(1, 223, 3, 4, 556)
	list.MoveBefore(list.Front(), list.Back())
	assert.Equal(t, 5, list.Len())
	assert.Equal(t, 223, list.Front().Value)
	assert.Equal(t, 556, list.Back().Value)

	list = ds.NewList(223, 556)
	list.MoveBefore(ds.NewElement(1, nil), list.Front())
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, 223, list.Front().Value)
	assert.Equal(t, 556, list.Back().Value)

	list = ds.NewList(223, 556)
	list.MoveBefore(list.Front(), ds.NewElement(1, nil))
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, 223, list.Front().Value)
	assert.Equal(t, 556, list.Back().Value)

}
