//nolint:mnd,varnamelen,exhaustruct
package queues

import (
	"slices"

	"github.com/SharkByteSoftware/go-snk/helpers"
)

// heapAdapter is an unexported type that manages the binary min-heap invariant
// entirely in terms of T, with no use of any or type assertions.
type heapAdapter[T any] struct {
	items      []T
	comparator func(a, b T) int
}

func (h *heapAdapter[T]) len() int { return len(h.items) }

func (h *heapAdapter[T]) less(i, j int) bool {
	return h.comparator(h.items[i], h.items[j]) < 0
}

func (h *heapAdapter[T]) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// push appends a value and sifts it up to restore the heap invariant.
func (h *heapAdapter[T]) push(value T) {
	h.items = append(h.items, value)
	h.siftUp(h.len() - 1)
}

// pop removes and returns the root (highest-priority) element and sifts down
// the replacement to restore the heap invariant.
func (h *heapAdapter[T]) pop() T {
	n := h.len() - 1
	h.swap(0, n)
	h.siftDown(0, n)
	val := h.items[n]
	h.items = h.items[:n]

	return val
}

// init establishes the heap invariant over an arbitrary slice in O(n) time.
func (h *heapAdapter[T]) init() {
	for i := h.len()/2 - 1; i >= 0; i-- {
		h.siftDown(i, h.len())
	}
}

// siftDown restores the heap invariant below index i within [0, n).
func (h *heapAdapter[T]) siftDown(i, n int) {
	for {
		left := 2*i + 1
		if left >= n {
			break
		}

		j := left
		if right := left + 1; right < n && h.less(right, left) {
			j = right
		}

		if !h.less(j, i) {
			break
		}

		h.swap(i, j)
		i = j
	}
}

// siftUp restores the heap invariant above index i.
func (h *heapAdapter[T]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !h.less(i, parent) {
			break
		}

		h.swap(i, parent)
		i = parent
	}
}

// PriorityQueue represents a generic priority queue where elements are ordered
// based on a custom comparator. The comparator follows the same contract as
// slices.SortFunc: return negative if a has higher priority than b, positive
// if b has higher priority than a, and zero if they are equal.
type PriorityQueue[T any] struct {
	h *heapAdapter[T]
}

// NewPriorityQueue creates a PriorityQueue pre-populated with a copy of the
// provided slice. The initial elements are heapified immediately, so the slice
// does not need to be sorted beforehand. The original slice is not modified.
func NewPriorityQueue[T any](items []T, comparator func(a, b T) int) *PriorityQueue[T] {
	h := &heapAdapter[T]{
		items:      slices.Clone(items),
		comparator: comparator,
	}
	h.init()

	return &PriorityQueue[T]{h: h}
}

// NewPriorityQueueWithDefault initializes a new empty PriorityQueue with the given comparator.
func NewPriorityQueueWithDefault[T any](comparator func(a, b T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		h: &heapAdapter[T]{comparator: comparator},
	}
}

// Enqueue inserts a new element into the priority queue in O(log n) time.
func (pq *PriorityQueue[T]) Enqueue(value T) {
	pq.h.push(value)
}

// Dequeue removes and returns the highest-priority element in O(log n) time.
// The boolean indicates whether the queue was non-empty.
func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	if pq.h.len() == 0 {
		return helpers.Empty[T](), false
	}

	return pq.h.pop(), true
}

// Peek returns the highest-priority element without removing it in O(1) time.
// The boolean indicates whether the queue was non-empty.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.h.len() == 0 {
		return helpers.Empty[T](), false
	}

	return pq.h.items[0], true
}

// IsEmpty reports whether the priority queue contains no elements.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.h.len() == 0
}

// Len returns the number of elements currently in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return pq.h.len()
}

// Size returns the number of elements currently in the priority queue.
func (pq *PriorityQueue[T]) Size() int {
	return pq.h.len()
}

// Clear removes all elements from the priority queue.
func (pq *PriorityQueue[T]) Clear() {
	pq.h.items = []T{}
}

// Values returns a clone of the elements currently in the priority queue.
// The order of elements is not guaranteed; use repeated Dequeue calls
// to retrieve elements in priority order.
func (pq *PriorityQueue[T]) Values() []T {
	if len(pq.h.items) == 0 {
		return []T{}
	}

	return slices.Clone(pq.h.items)
}
