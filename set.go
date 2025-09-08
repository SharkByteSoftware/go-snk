package sink

// Set is a set data structure.
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet creates a new set with the given items.
func NewSet[T comparable](items ...T) *Set[T] {
	set := &Set[T]{
		items: make(map[T]struct{}, len(items)),
	}

	set.Add(items...)

	return set
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

// Size returns the size of the set.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Clear clears the set.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Values returns the values of the set.
func (s *Set[T]) Values() []T {
	items := make([]T, len(s.items))

	idx := 0
	for item := range s.items {
		items[idx] = item
		idx++
	}

	return items
}
