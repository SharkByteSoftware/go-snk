package containers

type Container[T any] interface {
	IsEmpty() bool
	Size() int
	Clear()
	Values() []T
}

type Queue[T any] interface {
	Container[T]

	Enqueue(value T)
	Dequeue() (T, bool)
	Peek() (T, bool)
}

type Stack[T any] interface {
	Container[T]

	Push(value T)
	Pop() (T, bool)
	Peek() (T, bool)
}
