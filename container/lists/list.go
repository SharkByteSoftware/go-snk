package lists

import (
	"github.com/SharkByteSoftware/go-snk/conditionals"
)

// Element an element of a linked lists.
type Element[T comparable] struct {
	next, prev *Element[T]
	parent     *List[T]
	Value      T
}

func NewElement[T comparable](value T, parent *List[T]) *Element[T] {
	return &Element[T]{
		next:   nil,
		prev:   nil,
		parent: parent,
		Value:  value,
	}
}

func (e *Element[T]) Next() *Element[T] {
	if !e.isMemberOfList() || e.isBack() {
		return nil
	}

	return e.next
}

func (e *Element[T]) Prev() *Element[T] {
	if !e.isMemberOfList() || e.isFront() {
		return nil
	}

	return e.prev
}

func (e *Element[T]) isMemberOfList() bool {
	return e.parent != nil
}

func (e *Element[T]) isFront() bool {
	return e.prev == &e.parent.root
}

func (e *Element[T]) isBack() bool {
	return e.next == &e.parent.root
}

// List represents a doubly linked lists.  Api compatible with the Go
// containers List implementation.
type List[T comparable] struct {
	root Element[T]
	len  int
}

// NewList creates a new linked lists from all the values.
func NewList[T comparable](values ...T) *List[T] {
	result := &List[T]{
		root: Element[T]{
			next:   nil,
			prev:   nil,
			parent: nil,
			Value:  *new(T),
		},
		len: 0,
	}

	result.Init()

	result.Append(values...)

	return result
}

func (l *List[T]) Init() *List[T] {
	l.root.next = &l.root
	l.root.prev = &l.root

	l.len = 0

	return l
}

// Len returns the number of elements in the lists.
func (l *List[T]) Len() int {
	return l.len
}

// Front returns the first element in the lists. If the lists is empty,
// it will return nil.
func (l *List[T]) Front() *Element[T] {
	return conditionals.If(l.IsEmpty(), nil, l.root.next)
}

// Back returns the last element in the lists. If the lists is empty,
// it will return nil.
func (l *List[T]) Back() *Element[T] {
	return conditionals.If(l.IsEmpty(), nil, l.root.prev)
}

// IsEmpty checks to see if the lists is empty.
func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List[T]) Remove(element *Element[T]) T {
	if l.isNotMember(element) {
		return element.Value
	}

	return l.remove(element).Value
}

// PushFront inserts values to the front of the lists.
func (l *List[T]) PushFront(value T) *Element[T] {
	l.checkInit()

	return l.insertAt(NewElement(value, l), &l.root)
}

// Prepend adds values to the front of the lists.
func (l *List[T]) Prepend(values ...T) {
	l.checkInit()

	for idx := len(values) - 1; idx >= 0; idx-- {
		_ = l.insertValue(values[idx], &l.root)
	}
}

// PushBack adds the value to the end of the lists.
func (l *List[T]) PushBack(value T) *Element[T] {
	l.checkInit()

	return l.insertValue(value, l.root.prev)
}

// Append adds the values to the end of the lists.
func (l *List[T]) Append(values ...T) {
	l.checkInit()

	for _, value := range values {
		_ = l.insertValue(value, l.root.prev)
	}
}

// InsertBefore insert a new value before the mark and returns the element.
func (l *List[T]) InsertBefore(value T, mark *Element[T]) *Element[T] {
	if l.isNotMember(mark) {
		return nil
	}

	return l.insertValue(value, mark.prev)
}

// InsertAfter inserts a new value after the mark and returns the element.
func (l *List[T]) InsertAfter(value T, mark *Element[T]) *Element[T] {
	if l.isNotMember(mark) {
		return nil
	}

	return l.insertValue(value, mark)
}

// MoveToFront moves the element to the front of the lists.
func (l *List[T]) MoveToFront(element *Element[T]) {
	if l.isNotMember(element) || l.Front() == element {
		return
	}

	l.insertAt(l.remove(element), &l.root)
}

// MoveToBack moves the element to the back of the lists.
func (l *List[T]) MoveToBack(element *Element[T]) {
	if l.isNotMember(element) || l.Back() == element {
		return
	}

	l.insertAt(l.remove(element), l.root.prev)
}

// MoveBefore moves the element before the mark.
func (l *List[T]) MoveBefore(element *Element[T], mark *Element[T]) {
	if l.isNotMember(element) || l.isNotMember(mark) || element == mark {
		return
	}

	l.move(element, mark.prev)
}

// MoveAfter moves the element after the mark.
func (l *List[T]) MoveAfter(element *Element[T], mark *Element[T]) {
	if l.isNotMember(element) || l.isNotMember(mark) || element == mark {
		return
	}

	l.move(element, mark)
}

func (l *List[T]) PushBackList(other *List[T]) {
	l.Append(other.Values()...)
}

func (l *List[T]) PushFrontList(other *List[T]) {
	l.Prepend(other.Values()...)
}

func (l *List[T]) Values() []T {
	values := make([]T, 0, l.len)

	for e := l.Front(); e != nil; e = e.Next() {
		values = append(values, e.Value)
	}

	return values
}

func (l *List[T]) isNotMember(element *Element[T]) bool {
	return l != element.parent
}

func (l *List[T]) insertAt(element *Element[T], mark *Element[T]) *Element[T] {
	element.prev = mark
	element.next = mark.next
	element.prev.next = element
	element.next.prev = element
	element.parent = l

	l.len++

	return element
}

func (l *List[T]) insertValue(value T, mark *Element[T]) *Element[T] {
	element := NewElement(value, l)
	return l.insertAt(element, mark)
}

func (l *List[T]) remove(element *Element[T]) *Element[T] {
	element.prev.next = element.next
	element.next.prev = element.prev
	element.next = nil
	element.prev = nil
	element.parent = nil

	l.len--

	return element
}

func (l *List[T]) move(element *Element[T], mark *Element[T]) {
	if element == mark {
		return
	}

	element.prev.next = element.next
	element.next.prev = element.prev

	element.prev = mark
	element.next = mark.next
	element.prev.next = element
	element.next.prev = element
}

func (l *List[T]) checkInit() {
	if l.root.next == nil {
		l.Init()
	}
}
