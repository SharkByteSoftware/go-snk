package queues

import (
	"slices"
)

// PriorityQueue represents a generic priority queue where elements are ordered based on a custom comparator.
// items is a slice that holds the elements in the priority queue.
// comparator is a function that determines the priority order of elements in the queue.
type PriorityQueue[T any] struct {
	items      *[]T
	comparator func(prev T, curr T) int
}

// Len returns the number of elements currently present in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return len(*pq.items)
}

// NewPriorityQueue creates and returns a new PriorityQueue using the provided comparator to determine element order.
func NewPriorityQueue[T any](value *[]T, comparator func(prev T, curr T) int) PriorityQueue[T] {
	return PriorityQueue[T]{
		items:      value,
		comparator: comparator,
	}
}

// NewPriorityQueueWithDefault initializes a new PriorityQueue with a given comparator for ordering elements.
func NewPriorityQueueWithDefault[T any](comparator func(prev T, curr T) int) PriorityQueue[T] {
	return PriorityQueue[T]{
		items:      nil,
		comparator: comparator,
	}
}

// Enqueue inserts a new element into the priority queue and rearranges the elements based on the comparator function.
func (pq *PriorityQueue[T]) Enqueue(value T) {
	if pq.items == nil {
		pq.items = &[]T{value}
		return
	}

	newItems := append(*pq.items, value)
	pq.items = &newItems
	slices.SortFunc(*pq.items, pq.comparator)
}

// Dequeue removes and returns the highest-priority element from the
// priority queue. The boolean indicates success or failure.
func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	if pq.Len() > 0 {
		value := (*pq.items)[0]
		*pq.items = (*pq.items)[1:]

		return value, true
	}

	return *new(T), false
}

// Peek returns the highest-priority element without removing it. The boolean indicates if the queue is not empty.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.Len() == 0 {
		return *new(T), false
	}

	return (*pq.items)[0], true
}

// IsEmpty checks if the priority queue contains no elements and returns true if it is empty, otherwise returns false.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

// Size returns the total number of elements currently present in the priority queue.
func (pq *PriorityQueue[T]) Size() int {
	return len(*pq.items)
}

// Clear removes all elements from the priority queue, resetting it to an empty state.
func (pq *PriorityQueue[T]) Clear() {
	pq.items = nil
}

// Values return a slice of all the elements currently present in the priority queue without altering their order.
func (pq *PriorityQueue[T]) Values() []T {
	return slices.Clone(*pq.items)
}
