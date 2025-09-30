// Package queues provides queue implementations.
package queues

import "github.com/SharkByteSoftware/go-snk/containers/lists"

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

func (q *DQueue[T]) Dequeue() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Remove(q.members.Front()), true
}

func (q *DQueue[T]) Peek() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Front().Value, true
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
