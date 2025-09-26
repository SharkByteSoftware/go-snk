package containers

type Container[T any] interface {
	IsEmpty() bool
	Size() int
	Clear()
	Values() []T
}

type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, bool)
	Peek() (T, bool)
}
