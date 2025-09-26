package containers

type Container[T any] interface {
	IsEmpty() bool
	Size() int
	Clear()
	Values() []T
}
