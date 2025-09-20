// Package graphs provides a set data structure.
package sets

import (
	"github.com/SharkByteSoftware/go-snk/conditional"
)

// Set is a set data structure.
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet creates a new set with the given items.
func NewSet[T comparable](items ...T) Set[T] {
	set := Set[T]{items: make(map[T]struct{})}
	set.Add(items...)

	return set
}

// Add adds the given items to the set.
func (s *Set[T]) Add(item ...T) {
	for _, i := range item {
		s.items[i] = struct{}{}
	}
}

// IsEmpty returns true of the set contains zero items.
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Equals returns true of the two sets contain the same items.
func (s *Set[T]) Equals(other Set[T]) bool {
	if len(s.items) != len(other.items) {
		return false
	}

	for item := range s.items {
		if !other.Contains(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the set contains the given item.
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

// Remove removes the given item from the set.
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Size returns the len of the set.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Clear clears the set.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Values returns a slice of the set values.
func (s *Set[T]) Values() []T {
	items := make([]T, len(s.items))

	idx := 0
	for item := range s.items {
		items[idx] = item
		idx++
	}

	return items
}

// Intersect returns the intersection of the set with the given set.
func (s *Set[T]) Intersect(other Set[T]) Set[T] {
	size1 := s.Size()
	size2 := other.Size()
	smallSet := conditional.If(size1 < size2, *s, other)
	largerSet := conditional.If(size1 < size2, other, *s)

	result := NewSet[T]()

	for item := range smallSet.items {
		if largerSet.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// Union returns the union of the set with the given set.
func (s *Set[T]) Union(other Set[T]) Set[T] {
	size1 := s.Size()
	size2 := other.Size()
	smallSet := conditional.If(size1 < size2, *s, other)
	largerSet := conditional.If(size1 < size2, other, *s)

	result := largerSet.Clone()
	for item := range smallSet.items {
		result.Add(item)
	}

	return result
}

// Difference returns the difference of the set with the given set.
func (s *Set[T]) Difference(other Set[T]) Set[T] {
	result := NewSet[T]()

	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// SymmetricDifference returns a set with elements from either set but not both.
func (s *Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := NewSet[T]()

	s.Apply(func(item T) {
		if !other.Contains(item) {
			result.Add(item)
		}
	})

	other.Apply(func(item T) {
		if !s.Contains(item) {
			result.Add(item)
		}
	})

	return result
}

// Subset returns true of the set is a subset of a given set.
func (s *Set[T]) Subset(other Set[T]) bool {
	for item := range s.items {
		if !other.Contains(item) {
			return false
		}
	}

	return true
}

// Apply applies a function to each item in the set.
func (s *Set[T]) Apply(apply func(item T)) {
	for item := range s.items {
		apply(item)
	}
}

// Clone creates a clone of the set.
func (s *Set[T]) Clone() Set[T] {
	return NewSet[T](s.Values()...)
}
