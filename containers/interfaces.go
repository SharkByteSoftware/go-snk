// Package containers provides a variety of container types.
package containers

// Container is the base interface for all data structures to implement.
type Container[T any] interface {
	IsEmpty() bool
	Size() int
	Clear()
	Values() []T
}

// Queue is a base interface for all queue implementations to implement.
type Queue[T any] interface {
	Container[T]

	Enqueue(value T)
	Dequeue() (T, bool)
	Peek() (T, bool)
}

// Stack is a base interface for all stack implementations to implement.
type Stack[T any] interface {
	Container[T]

	Push(value T)
	Pop() (T, bool)
	Peek() (T, bool)
}
