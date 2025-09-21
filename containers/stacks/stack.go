// Package stacks contains a stack implementation.
package stacks

import (
	"github.com/SharkByteSoftware/go-snk/containers/lists"
)

type Stack[T comparable] struct {
	members *lists.List[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		members: lists.New[T](),
	}
}

func (s *Stack[T]) Push(value T) {
	s.members.PushFront(value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.members.IsEmpty() {
		return *new(T), false
	}

	return s.members.Remove(s.members.Front()), true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.members.IsEmpty() {
		return *new(T), false
	}

	return s.members.Front().Value, true
}

func (s *Stack[T]) Size() int {
	return s.members.Len()
}

func (s *Stack[T]) Values() []T {
	return s.members.Values()
}
