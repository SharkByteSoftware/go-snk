package pipeline

import "context"

func NewInputStream[T any](ctx context.Context, fn func(ctx context.Context) T) <-chan T {
	// TODO: implement
	panic("implement me")
}

func NewSliceStream[T any](in []T) <-chan T {
	// TODO: implement
	panic("implement me")
}

func NewMapStream[K comparable, V any](in map[K]V) <-chan K {
	// TODO: implement
	panic("implement me")
}

func FanOut[T any, OutT any](ctx context.Context, input <-chan T,
	process func(context.Context, T) OutT, concurrency int) []<-chan OutT {
	// TODO: implement
	panic("implement me")
}

func FanIn[T any](ctx context.Context, input <-chan T) <-chan T {
	// TODO: implement
	panic("implement me")
}

func Filter[T any, OutT any](ctx context.Context, process func(context.Context, T) OutT) <-chan OutT {
	// TODO: implement
	panic("implement me")
}
