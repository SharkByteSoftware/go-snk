package queues

type PriorityQueue[T any] struct{}

func NewPriorityQueue[T any](comparator func(prev T, curr T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (q *PriorityQueue[T]) Enqueue(value T) {
	// TODO: implement
	panic("implement me")
}

func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	// TODO: implement
	panic("implement me")
}

func (q *PriorityQueue[T]) Peek() (T, bool) {
	// TODO: implement
	panic("implement me")
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	// TODO implement me
	panic("implement me")
}

func (q *PriorityQueue[T]) Size() int {
	// TODO implement me
	panic("implement me")
}

func (q *PriorityQueue[T]) Clear() {
	// TODO implement me
	panic("implement me")
}

func (q *PriorityQueue[T]) Values() []T {
	// TODO implement me
	panic("implement me")
}
