package ds

import "github.com/SharkByteSoftware/go-snk/conditionals"

// Element an element of a linked list.
type Element[T comparable] struct {
	next, prev *Element[T]
	parent     *List[T]
	Value      T
}

func NewListElement[T comparable](value T, parent *List[T]) *Element[T] {
	return &Element[T]{
		next:   nil,
		prev:   nil,
		parent: parent,
		Value:  value,
	}
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}

// List represents a doubly linked list.  Api compatible with the Go
// containers List implementation.
type List[T comparable] struct {
	root Element[T]
	len  int
}

// NewList creates a new linked list from all the values.
func NewList[T comparable](values ...T) *List[T] {
	result := &List[T]{
		root: Element[T]{
			Value:  *new(T),
			next:   nil,
			prev:   nil,
			parent: nil,
		},
		len: 0,
	}

	result.root.next = &result.root
	result.root.prev = &result.root
	result.root.parent = result

	result.PushBack(values...)

	return result
}

// Len returns the number of elements in the list.
func (l *List[T]) Len() int {
	return l.len
}

// Front returns the first element in the list. If the list is empty,
// it will return nil.
func (l *List[T]) Front() *Element[T] {
	return conditionals.If(l.Len() == 0, nil, l.root.next)
}

// Back returns the last element in the list. If the list is empty,
// it will return nil.
func (l *List[T]) Back() *Element[T] {
	return conditionals.If(l.Len() == 0, nil, l.root.prev)
}

// IsEmpty checks to see if the list is empty.
func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List[T]) Remove(element *Element[T]) T {
	if l.isElementMemberOfList(element) {
		element.prev.next = element.next
		element.next.prev = element.prev
		element.next = nil
		element.prev = nil
		l.len--
	}

	return element.Value
}

// PushFront inserts values to the front of the list.
func (l *List[T]) PushFront(values ...T) {
	for idx := len(values) - 1; idx >= 0; idx-- {
		_ = l.insertValue(values[idx], &l.root)
	}
}

// PushBack adds the values to the end of the list.
func (l *List[T]) PushBack(values ...T) {
	for _, value := range values {
		_ = l.insertValue(value, l.root.prev)
	}
}

func (l *List[T]) InsertBefore(value T, mark *Element[T]) *Element[T] {
	return conditionals.If(l.isElementMemberOfList(mark),
		l.insertValue(value, mark.prev), nil)
}

func (l *List[T]) InsertAfter(value T, mark *Element[T]) *Element[T] {
	return conditionals.If(l.isElementMemberOfList(mark),
		l.insertValue(value, mark), nil)
}

func (l *List[T]) MoveToFront(element *Element[T]) {
	if !l.isElementMemberOfList(element) || l.Front() == element {
		return
	}
	
	l.insertValue(l.Remove(element), &l.root)
}

func (l *List[T]) MoveToBack(element *Element[T]) {
	if !l.isElementMemberOfList(element) || l.Back() == element {
		return
	}

	l.insertValue(l.Remove(element), l.root.prev)
}

func (l *List[T]) MoveBefore(element *Element[T], mark *Element[T]) {
}

func (l *List[T]) MoveAfter(element *Element[T], mark *Element[T]) {
}

func (l *List[T]) PushBackList(other *List[T]) {}

func (l *List[T]) PushFrontList(other *List[T]) {}

func (l *List[T]) isElementMemberOfList(element *Element[T]) bool {
	return l == element.parent
}

func (l *List[T]) insertAt(element *Element[T], atLocation *Element[T]) *Element[T] {
	element.prev = atLocation
	element.next = atLocation.next
	element.prev.next = element
	element.next.prev = element
	l.len++

	return element
}

func (l *List[T]) insertValue(value T, at *Element[T]) *Element[T] {
	element := NewListElement(value, l)
	return l.insertAt(element, at)
}
