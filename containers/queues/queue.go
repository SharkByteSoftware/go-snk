package queues

type Queue[T any] struct{}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {}

func (q *Queue[T]) Dequeue() (value T, ok bool) {
	return *new(T), false
}

func (q *Queue[T]) Peek() (value T, ok bool) {
	return *new(T), false
}

func (q *Queue[T]) Values() []T {
	return []T{}
}
