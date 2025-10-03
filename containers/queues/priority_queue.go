package queues

import (
	"container/heap"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

type Item[T any] struct {
	value    T
	index    int
	priority int
}

type PriorityQueue[T any] struct {
	items []*Item[T]
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.items[i].priority > pq.items[j].priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue[T]) Push(x any) {
	pq.Enqueue(x.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	x, ok := pq.Dequeue()
	if !ok {
		return nil
	}
	return x
}

func NewPriorityQueue[T any](comparator func(prev T, curr T) int) PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func (pq *PriorityQueue[T]) Enqueue(value T) {
	n := pq.Len()
	item := Item[T]{value: value, index: n}
	pq.items = append(pq.items, &item)
}

func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	old := pq
	n := old.Len()
	item := old.items[n-1]
	old.items[n-1] = nil
	item.index = -1
	pq.items = old.items[0 : n-1]
	return item.value, true
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.Len() == 0 {
		return *new(T), false
	}

	return pq.items[0].value, true
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.Len()
}

func (pq *PriorityQueue[T]) Clear() {
	pq = &PriorityQueue[T]{}
}

func (pq *PriorityQueue[T]) Values() []T {
	return slicex.Map(pq.items, func(item *Item[T]) T { return item.value })
}
