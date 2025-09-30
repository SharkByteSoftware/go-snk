// Package stacks contains a stack implementation.
package stacks

import (
	"github.com/SharkByteSoftware/go-snk/containers/lists"
)

// Stack provides a stack implementation based on a link list.
type Stack[T comparable] struct {
	members *lists.List[T]
}

// New creates a new stack of a given type.
func New[T comparable](values ...T) *Stack[T] {
	return &Stack[T]{
		members: lists.New[T](values...),
	}
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.members.PushFront(value)
}

// Pop removes the top element of the stack and returns it.  If stack is empty,
// it returns a default value and false.
func (s *Stack[T]) Pop() (T, bool) {
	if s.members.IsEmpty() {
		return *new(T), false
	}

	return s.members.Remove(s.members.Front()), true
}

// Peek returns the top element on the stack without removing it.   If stack is empty,
// it returns a default value and false.
func (s *Stack[T]) Peek() (T, bool) {
	if s.members.IsEmpty() {
		return *new(T), false
	}

	return s.members.Front().Value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.members.IsEmpty()
}

// Size returns the number of elements on the stack.
func (s *Stack[T]) Size() int {
	return s.members.Len()
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.members = lists.New[T]()
}

// Values returns a slice with all the elements fromt he stack.
func (s *Stack[T]) Values() []T {
	return s.members.Values()
}
