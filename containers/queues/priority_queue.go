package queues

type PriorityQueue[T any] struct{}

func NewPriorityQueue[T any](comparator func(prev T, curr T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (q *PriorityQueue[T]) Enqueue(value T) {}

func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	return *new(T), false
}

func (q *PriorityQueue[T]) Peek() (T, bool) {
	return *new(T), false
}

func (q *PriorityQueue[T]) Values() []T {
	return []T{}
}
