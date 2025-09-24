package queues

type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Values() []T
}
