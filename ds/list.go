package ds

import "github.com/SharkByteSoftware/go-snk/conditionals"

type ListElement[T comparable] struct {
	next, prev *ListElement[T]
	parent     *List[T]
	Value      T
}

func (e *ListElement[T]) Next() *ListElement[T] {
	return e.next
}

func (e *ListElement[T]) Prev() *ListElement[T] {
	return e.prev
}

func (e *ListElement[T]) Parent() *List[T] {
	return e.parent
}

type List[T comparable] struct {
	first *ListElement[T]
	last  *ListElement[T]
	size  int
}

// NewList creates a new linked list from all the values.
func NewList[T comparable](values ...T) *List[T] {
	result := &List[T]{}
	result.Add(values...)

	return result
}

func (l *List[T]) First() *ListElement[T] {
	return l.first
}

func (l *List[T]) Last() *ListElement[T] {
	return l.last
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) IsEmpty() bool {
	return l.Size() == 0
}

// Add adds the values to the end of the list
func (l *List[T]) Add(values ...T) {
	for _, value := range values {
		_ = l.insertValue(value, l.last)
	}
}

func (l *List[T]) Append(values ...T) {
	l.Add(values...)
}

func (l *List[T]) Prepend(values ...T) {
}

func (l *List[T]) insertAt(element *ListElement[T], at *ListElement[T]) *ListElement[T] {
	if l.IsEmpty() {
		l.first = element
		l.last = element
		l.size++

		return element
	}

	if at == nil {
		return nil
	}

	element.prev = at
	element.next = at.next
	element.prev.next = element

	conditionals.IfNotNil(element.next, func() { element.next.prev = element })

	//e.prev = at
	//e.next = at.next
	//e.prev.next = e
	//e.next.prev = e

	l.last = conditionals.If(l.last == at, element, at)
	l.size++

	return element
}

func (l *List[T]) insertValue(value T, at *ListElement[T]) *ListElement[T] {
	element := &ListElement[T]{Value: value, parent: l}
	return l.insertAt(element, at)
}
