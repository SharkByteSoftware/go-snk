package lists_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/container/lists"
	"github.com/stretchr/testify/assert"
)

func TestList_NewList(t *testing.T) {
	lst := lists.NewList[int]()
	assert.Equal(t, 0, lst.Len())
	assert.Nil(t, lst.Front())
	assert.Nil(t, lst.Back())

	lst = lists.NewList(1)
	assert.Equal(t, 1, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 1, lst.Back().Value)

	lst = lists.NewList(1, 2, 4, 6)
	assert.Equal(t, 4, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 6, lst.Back().Value)
}

func TestList_IsEmpty_Len(t *testing.T) {
	lst := lists.NewList[int]()
	assert.True(t, lst.IsEmpty())
	assert.Equal(t, 0, lst.Len())

	lst = lists.NewList(1, 2, 4, 6)
	assert.False(t, lst.IsEmpty())
	assert.Equal(t, 4, lst.Len())
}

func TestList_Remove(t *testing.T) {
	lst := lists.NewList(1, 556, 2, 3, 223, 5)

	value := lst.Remove(lst.Front())
	assert.Equal(t, 1, value)
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 556, lst.Front().Value)
	assert.Equal(t, 5, lst.Back().Value)

	value = lst.Remove(lst.Back())
	assert.Equal(t, 5, value)
	assert.Equal(t, 4, lst.Len())
	assert.Equal(t, 223, lst.Back().Value)

	value = lst.Remove(lst.Front().Next())
	assert.Equal(t, 2, value)
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 556, lst.Front().Value)
	assert.Equal(t, 223, lst.Back().Value)

	lst = lists.NewList(1, 2, 4, 6)
	value = lst.Remove(lists.NewElement(1, nil))
	assert.Equal(t, 1, value)
	assert.Equal(t, 4, lst.Len())

	lst = lists.NewList(1)
	element := lst.Front()
	_ = lst.Remove(element)
	_ = lst.Remove(element)

}

func TestList_Prepend(t *testing.T) {
	lst := lists.NewList[int](1)
	assert.Equal(t, 1, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 1, lst.Back().Value)

	lst.Prepend(2)
	assert.Equal(t, 2, lst.Len())
	assert.Equal(t, 2, lst.Front().Value)
	assert.Equal(t, 1, lst.Back().Value)

	lst.Prepend(12, 33, 2, 4, 6)
	assert.Equal(t, 7, lst.Len())
	assert.Equal(t, 12, lst.Front().Value)
	assert.Equal(t, 1, lst.Back().Value)
}

func TestList_Append(t *testing.T) {
	lst := lists.NewList[int](1)
	assert.Equal(t, 1, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 1, lst.Back().Value)

	lst.Append(2)
	assert.Equal(t, 2, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 2, lst.Back().Value)

	lst.Append(12, 33, 2, 4, 6)
	assert.Equal(t, 7, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 6, lst.Back().Value)
}

func TestList_InsertBefore(t *testing.T) {
	lst := lists.NewList[int](223, 556)

	element := lst.InsertBefore(1, lst.Front())
	assert.Equal(t, 1, element.Value)
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)

	element = lst.InsertBefore(2, lst.Back())
	assert.Equal(t, 2, element.Value)
	assert.Equal(t, 4, lst.Len())
	assert.Equal(t, 2, lst.Back().Prev().Value)

	element = lst.InsertBefore(308, lst.Front().Next())
	assert.Equal(t, 308, element.Value)
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 308, lst.Front().Next().Value)

	lst = lists.NewList(1, 2, 4, 6)
	element = lst.InsertBefore(22, lists.NewElement(1, nil))
	assert.Nil(t, element)
	assert.Equal(t, 4, lst.Len())
}

func TestList_InsertAfter(t *testing.T) {
	lst := lists.NewList[int](223, 556)

	element := lst.InsertAfter(1, lst.Front())
	assert.Equal(t, 1, element.Value)
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 1, lst.Front().Next().Value)

	element = lst.InsertAfter(2, lst.Back())
	assert.Equal(t, 2, element.Value)
	assert.Equal(t, 4, lst.Len())
	assert.Equal(t, 2, lst.Back().Value)

	element = lst.InsertAfter(308, lst.Front().Next())
	assert.Equal(t, 308, element.Value)
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 308, lst.Front().Next().Next().Value)

	lst = lists.NewList(1, 2, 4, 6)
	element = lst.InsertAfter(22, lists.NewElement(1, nil))
	assert.Nil(t, element)
	assert.Equal(t, 4, lst.Len())
}

func TestList_MoveToFront(t *testing.T) {
	lst := lists.NewList[int](1, 223, 3, 4, 556)

	lst.MoveToFront(lst.Front())
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)

	lst.MoveToFront(lst.Back())
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 556, lst.Front().Value)

	lst.MoveToFront(lst.Front().Next().Next())
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)

	lst.MoveToFront(lists.NewElement(1, nil))
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
}

func TestList_MoveToBack(t *testing.T) {
	lst := lists.NewList[int](1, 223, 3, 4, 556)

	lst.MoveToBack(lst.Back())
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 556, lst.Back().Value)

	lst.MoveToBack(lst.Front())
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 1, lst.Back().Value)

	lst.MoveToBack(lst.Back().Prev())
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 556, lst.Back().Value)

	lst.MoveToBack(lists.NewElement(1, nil))
	assert.Equal(t, 5, lst.Len())
	assert.Equal(t, 556, lst.Back().Value)
}

func TestList_MoveBefore(t *testing.T) {
	lst := lists.NewList(223, 1, 556)
	lst.MoveBefore(lst.Front(), lst.Front())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
	assert.Equal(t, 556, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveBefore(lst.Front(), lst.Back())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 556, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveBefore(lst.Back(), lst.Front())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 556, lst.Front().Value)
	assert.Equal(t, 1, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveBefore(lists.NewElement(1, nil), lst.Back())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
	assert.Equal(t, 556, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveBefore(lst.Front(), lists.NewElement(1, nil))
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
	assert.Equal(t, 556, lst.Back().Value)
}

func TestList_MoveAfter(t *testing.T) {
	lst := lists.NewList(223, 1, 556)
	lst.MoveAfter(lst.Front(), lst.Front())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
	assert.Equal(t, 556, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveAfter(lst.Front(), lst.Back())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 1, lst.Front().Value)
	assert.Equal(t, 223, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveAfter(lst.Back(), lst.Front())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
	assert.Equal(t, 1, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveAfter(lists.NewElement(1, nil), lst.Back())
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
	assert.Equal(t, 556, lst.Back().Value)

	lst = lists.NewList(223, 1, 556)
	lst.MoveAfter(lst.Front(), lists.NewElement(1, nil))
	assert.Equal(t, 3, lst.Len())
	assert.Equal(t, 223, lst.Front().Value)
	assert.Equal(t, 556, lst.Back().Value)
}
