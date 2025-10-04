package queues

// PriorityQueue implements a priority queue.
type PriorityQueue[T any] struct{}

// NewPriorityQueue creates a new priority queue with the given comparator function.
func NewPriorityQueue[T any](_ func(prev T, curr T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

// Enqueue adds a value to the priority queue.
func (q *PriorityQueue[T]) Enqueue(_ T) {
	// TODO: implement
	panic("implement me")
}

// Dequeue removes the next value in the priority queue and removes it.
func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	// TODO: implement
	panic("implement me")
}

// Peek returns the next value in the priority queue without removing i.
func (q *PriorityQueue[T]) Peek() (T, bool) {
	// TODO: implement
	panic("implement me")
}

// IsEmpty returns true of the priority queue is empty.
func (q *PriorityQueue[T]) IsEmpty() bool {
	// TODO implement me
	panic("implement me")
}

// Size returns the size of the priority queue.
func (q *PriorityQueue[T]) Size() int {
	// TODO implement me
	panic("implement me")
}

// Clear removes all the values from the priority queue.
func (q *PriorityQueue[T]) Clear() {
	// TODO implement me
	panic("implement me")
}

// Values returns a slice with all the values in the priority queue.
func (q *PriorityQueue[T]) Values() []T {
	// TODO implement me
	panic("implement me")
}
