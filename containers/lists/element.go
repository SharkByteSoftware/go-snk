package lists

// Element represents an element of a linked list.
type Element[T any] struct {
	next, prev *Element[T]
	parent     *List[T]
	Value      T
}

// NewElement creates a new element with a given value and linked list parent.
func NewElement[T any](value T, parent *List[T]) *Element[T] {
	return &Element[T]{
		next:   nil,
		prev:   nil,
		parent: parent,
		Value:  value,
	}
}

// Next returns the next element in the list. It will return
// nil if element is the end of the list.
func (e *Element[T]) Next() *Element[T] {
	if !e.isMemberOfList() || e.isBack() {
		return nil
	}

	return e.next
}

// Prev returns the previous element in the list.  It will return
// nil if element is the front of the list.
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
