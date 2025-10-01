// Package queues provides queue implementations.
package queues

import "github.com/SharkByteSoftware/go-snk/containers/lists"

// DQueue provides an implementation of a double-ended queue based on a linked list.
type DQueue[T any] struct {
	members *lists.List[T]
}

// NewQueue creates a new double-ended queue with optional initial values.
func NewQueue[T any](values ...T) *DQueue[T] {
	return &DQueue[T]{
		members: lists.New[T](values...),
	}
}

// Enqueue adds a value to the end of the queue.
func (q *DQueue[T]) Enqueue(value T) {
	q.members.PushBack(value)
}

// EnqueueFront Enqueue adds a value to the front of the queue.
func (q *DQueue[T]) EnqueueFront(value T) {
	q.members.PushFront(value)
}

// Dequeue removes and returns the value from the front of the queue;
// returns default and false if empty.
func (q *DQueue[T]) Dequeue() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Remove(q.members.Front()), true
}

// DequeueBack removes and returns the value from the back of the queue;
// returns default and false if empty.
func (q *DQueue[T]) DequeueBack() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Remove(q.members.Back()), true
}

// Peek returns the value from the front of the queue;
// returns default and false if empty.
func (q *DQueue[T]) Peek() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Front().Value, true
}

// PeekBack returns the value from the back of the queue;
// returns default and false if empty.
func (q *DQueue[T]) PeekBack() (T, bool) {
	if q.members.IsEmpty() {
		return *new(T), false
	}

	return q.members.Back().Value, true
}

// IsEmpty checks if the queue is empty.
func (q *DQueue[T]) IsEmpty() bool {
	return q.members.IsEmpty()
}

// Size returns the number of elements in the queue.
func (q *DQueue[T]) Size() int {
	return q.members.Len()
}

// Clear removes all elements from the queue.
func (q *DQueue[T]) Clear() {
	q.members = lists.New[T]()
}

// Values returns a slice of all the values in the queue.
func (q *DQueue[T]) Values() []T {
	return q.members.Values()
}
