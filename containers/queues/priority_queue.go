package queues

import (
	"slices"
)

type PriorityQueue[T any] struct {
	items      []T
	comparator func(prev T, curr T) int
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq.items)
}

func NewPriorityQueue[T any](comparator func(prev T, curr T) int) PriorityQueue[T] {
	return PriorityQueue[T]{
		items:      make([]T, 0),
		comparator: comparator,
	}
}

func (pq *PriorityQueue[T]) Enqueue(value T) {
	pq.items = append(pq.items, value)
	slices.SortFunc(pq.Values(), pq.comparator)
}

func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	if pq.Len() > 0 {
		value := pq.items[0]
		pq.items = pq.items[1:]
		return value, true
	}

	return *new(T), false
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.Len() == 0 {
		return *new(T), false
	}

	return pq.items[0], true
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue[T]) Size() int {
	return len(pq.items)
}

func (pq *PriorityQueue[T]) Clear() {
	pq = &PriorityQueue[T]{
		items:      make([]T, 0),
		comparator: pq.comparator,
	}
}

func (pq *PriorityQueue[T]) Values() []T {
	return pq.items
}
