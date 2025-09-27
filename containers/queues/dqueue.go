// Package queues provides queue implementations.
package queues

type DQueue[T any] struct{}

func NewQueue[T any]() *DQueue[T] {
	return &DQueue[T]{}
}

func (q *DQueue[T]) Enqueue(value T) {
	// TODO: implement
	panic("implement me")
}

func (q *DQueue[T]) Dequeue() (T, bool) {
	// TODO: implement
	panic("implement me")
}

func (q *DQueue[T]) Peek() (T, bool) {
	// TODO: implement
	panic("implement me")
}

func (q *DQueue[T]) IsEmpty() bool {
	// TODO implement me
	panic("implement me")
}

func (q *DQueue[T]) Size() int {
	// TODO implement me
	panic("implement me")
}

func (q *DQueue[T]) Clear() {
	// TODO implement me
	panic("implement me")
}

func (q *DQueue[T]) Values() []T {
	// TODO implement me
	panic("implement me")
}
