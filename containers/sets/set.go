// Package sets provides a set data structure.
package sets

// Set implements a set data structure.
type Set[T comparable] struct {
	items map[T]struct{}
}

// New creates a new set with the given items.
func New[T comparable](items ...T) *Set[T] {
	set := Set[T]{items: make(map[T]struct{}, len(items))}
	set.Add(items...)

	return &set
}

// Add adds the given items to the set.
func (s *Set[T]) Add(item ...T) {
	for _, i := range item {
		s.items[i] = struct{}{}
	}
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

// Intersect returns the intersection of the set with the given set.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	size1 := s.Size()
	size2 := other.Size()

	smallSet := s
	largerSet := other

	if size1 > size2 {
		smallSet = other
		largerSet = s
	}

	result := &Set[T]{items: make(map[T]struct{}, max(size1, size2))}

	for item := range smallSet.items {
		if largerSet.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// Union returns the union of the set with the given set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	smallSet := s
	largerSet := other

	if s.Size() > other.Size() {
		smallSet = other
		largerSet = s
	}

	result := largerSet.Clone()
	for item := range smallSet.items {
		result.Add(item)
	}

	return result
}

// Difference returns a new set with elements that are in s but not in the other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()

	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// SymmetricDifference returns a set with elements from either set but not both.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	result := New[T]()

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

// Subset returns true if every element of s is contained in the other.
func (s *Set[T]) Subset(other *Set[T]) bool {
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

// Equals returns true if the two sets contain the same items.
func (s *Set[T]) Equals(other *Set[T]) bool {
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

// Clone creates a clone of the set.
func (s *Set[T]) Clone() *Set[T] {
	result := &Set[T]{items: make(map[T]struct{}, len(s.items))}

	for k := range s.items {
		result.items[k] = struct{}{}
	}

	return result
}

// IsEmpty returns true if the set contains zero items.
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the set.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Clear removes all items from the set.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Values returns a slice of all items in the set in no particular order.
func (s *Set[T]) Values() []T {
	items := make([]T, len(s.items))

	idx := 0
	for item := range s.items {
		items[idx] = item
		idx++
	}

	return items
}
