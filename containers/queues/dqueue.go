// Package queues provides queue implementations.
package queues

import "github.com/SharkByteSoftware/go-snk/containers/lists"

// DQueue provides an implementation of a double-ended queue based on a linked list.
type DQueue[T any] struct {
	members *lists.List[T]
}

func NewQueue[T any](values ...T) *DQueue[T] {
	return &DQueue[T]{
		members: lists.New[T](values...),
	}
}

func (q *DQueue[T]) Enqueue(value T) {
	q.members.PushBack(value)
}

func (q *DQueue[T]) EnqueueFront(value T) {
	q.members.PushFront(value)
}

func (q *DQueue[T]) Dequeue() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Remove(q.members.Front()), true
}

func (q *DQueue[T]) DequeueBack() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Remove(q.members.Back()), true
}

func (q *DQueue[T]) Peek() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Front().Value, true
}

func (q *DQueue[T]) PeekBack() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Back().Value, true
}

func (q *DQueue[T]) IsEmpty() bool {
	return q.members.IsEmpty()
}

func (q *DQueue[T]) Size() int {
	return q.members.Len()
}

func (q *DQueue[T]) Clear() {
	q.members = lists.New[T]()
}

func (q *DQueue[T]) Values() []T {
	return q.members.Values()
}
