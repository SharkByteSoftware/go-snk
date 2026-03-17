<div align="center">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="130" alt="sink-logo" src="img/logo.png" />
    </a>
</div>

# go-snk - Slices, Maps, Data Structures, Channels, and more...

[![GitHub Tag](https://img.shields.io/github/v/tag/SharkByteSoftware/go-snk)](https://github.com/SharkByteSoftware/go-snk/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/SharkByteSoftware/go-snk.svg)](https://pkg.go.dev/github.com/SharkByteSoftware/go-snk)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/SharkByteSoftware/go-snk)](./go.mod)
[![Build Status](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/go.yml)
[![Lint Status](https://github.com/SharkByteSoftware/go-snk/actions/workflows/lint.yml/badge.svg)](https://github.com/SharkByteSoftware/go-snk/actions/workflows/lint.yml)
[![Go report](https://goreportcard.com/badge/github.com/SharkByteSoftware/go-snk)](https://goreportcard.com/report/SharkByteSoftware/go-snk)
[![GitHub License](https://img.shields.io/github/license/SharkByteSoftware/go-snk)](./LICENSE)
[![Contributors](https://img.shields.io/github/contributors/SharkByteSoftware/go-snk)](https://github.com/SharkByteSoftware/go-snk/graphs/contributors)

Everything and the kitchen sink for Go.

A utility library that provides a variety of functions for working with slices, maps, channels, and more.

## About

**go-snk** is a Go utility library that brings together practical, type-safe helpers for slices, maps, sets, containers, conditionals, 
and HTTP requests and responses. Built with Go generics, it helps you write cleaner, more expressive code with fewer repetitive utility 
functions.

## Getting Started

### Prerequisites

- **Go version**: go-snk uses [Go](https://go.dev/) version [1.24](https://go.dev/doc/devel/release#go1.24.0) or above
- **Basic Go knowledge**: Familiarity with Go syntax and package management is helpful

### Installation

With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), you can simply import `go-snk`, and Go will automatically fetch it during the build:

```go
import "github.com/SharkByteSoftware/go-snk"
```

Or

you can use `go get` command to get the latest version of `go-snk`:

```sh
go get github.com/SharkByteSoftware/go-snk@latest
```

## Features

### slicex - Helpers for slices

| Helpers      | Description                                                                |
|--------------|----------------------------------------------------------------------------|
| FirstOr      | Returns the first item of a slice or a fallback value of slice is empty    |
| FirstOrEmpty | Returns the first item of a slice or the empty value if the slice is empty |
| Filter       | Filters a slice using a predicate                                          |
| Map          | Transforms a slice to a slice of another type using a mapper               |
| FilterMap    | Filters and transforms a slice to a slice of another type using a mapper   |
| UniqueMap    | Similar to Map but removes duplicates from the result                      |
| Reduce       | Transforms and flattens a slice to another type                            |
| Find         | Returns the first element matching a condition                             |
| FindBy       | Returns the first element matching a predicate function                    |
| FindOr       | Returns the first matching element or a default value                      |
| Any          | Checks if any element satisfies a condition                                |
| AnyBy        | Checks if any element satisfies a predicate function                       |
| All          | Checks if all elements are equal to a given candidate                      |
| AllBy        | Checks if all elements satisfy a predicate                                 |
| Unique       | Removes duplicates from a slice                                            |
| UniqueBy     | Removes duplicates from a slice as determined by a predicate function      |
| Apply        | Applies a function to each element in a slice                              |
| Reverse      | Reverses the order of elements in a slice                                  |
| Compact      | Compact returns a slice with all the non-zero items.                       |
| ToMap        | Converts a slice to a map using a key selector function                    |
| GroupBy      | Groups elements of a slice by a key selector function                      |
| Partition    | Splits a slice into two based on a predicate function                      |
| Intersect    | Returns a slice with the intersection of two slices                        |
| Union        | Returns a slice with the union of two slices                               |
| Difference   | Returns a slice with the difference of two slices                          |                             

| Math Helpers | Description                                                           |
|--------------|-----------------------------------------------------------------------|
| Sum          | Calculates the sum of a slice of numeric values                       |
| SumBy        | Calculates the sum of a slice using a custom value function           |
| Product      | Calculates the product of a slice of numeric values                   |
| ProductBy    | Calculates the product of a slice using a custom value function       |
| Mean         | Calculates the arithmetic mean of a slice of numeric values           |
| MeanBy       | Calculates the mean of a slice using a custom value function          |
| Max          | Finds the maximum value in a slice of comparable values               |
| MaxBy        | Finds the maximum value in a slice using a custom comparison function |
| Min          | Finds the minimum value in a slice of comparable values               |
| MinBy        | Finds the minimum value in a slice using a custom comparison function |

### slicex/parallel

| Function           | Description                                                                                      |
|--------------------|--------------------------------------------------------------------------------------------------|
| Map                | Transforms a slice to a slice of another type using a mapper function in parallel, preserving order |
| MapWithLimit       | Same as Map but limits the concurrency                                                           |
| Apply              | Applies a function to each item in a slice in parallel                                           |
| ApplyWithLimit     | Same as Apply but limits the concurrency                                                         |
| GroupBy            | Groups a slice into a map of slices based on a predicate function in parallel                    |
| GroupByWithLimit   | Same as GroupBy but limits the concurrencyl                                                      |
| Partition          | Splits a slice into two slices based on a predicate function in parallel, preserving order       |
| PartitionWithLimit | Same as Partition but limits the concurrency                                                     |

### mapx - Helpers for maps

| Helpers           | Description                                                         |
|-------------------|---------------------------------------------------------------------|
| Keys              | Returns a slice of the map's keys                                   |
| Values            | Returns a slice of the map's values                                 |
| UniqueValues      | Returns a slice of unique values from the map                       |
| Contains          | Checks if the map contains all specified keys                       |
| ValueOr           | Returns the value for a key or a fallback value if not found        |
| Invert            | Inverts the map, swapping keys and values                           |
| Combine           | Combines multiple maps into a single map                            |
| ToSlice           | Converts a map to a slice using a mapper function                   |
| Filter            | Filters a map based on a predicate function                         |
| Apply             | Applies a function to each key-value pair in the map                |

### httpx - Helpers for HTTP requests and responses

| Helper               | Description                                                                  |
|----------------------|------------------------------------------------------------------------------|
| Get                  | Sends a GET request and returns a processed response                         |
| GetRawResponse       | Sends a GET request and returns the raw HTTP response                        |
| Post                 | Sends a POST request and returns a processed response                        |
| PostRawResponse      | Sends a POST request and returns the raw HTTP response                       |
| Put                  | Sends a PUT request and returns a processed response                         |
| PutRawResponse       | Sends a PUT request and returns the raw HTTP response                        |
| Delete               | Sends a DELETE request and returns a processed response                      |
| DeleteRawResponse    | Sends a DELETE request and returns the raw HTTP response                     |
| ConfigOptions        | Configures request behavior and response handling options                    |
| AlwaysIncludeRawBody | Configures responses to always retain the raw response body                  |
| DecodeRawBody        | Decodes a raw response body into a target value                              |

### Conditionals

| Conditional  | Description                                                            |
|--------------|------------------------------------------------------------------------|
| If           | Returns one of two values based on a condition                         |
| IfNotNil     | Calls a function if the input pointer is not nil                       |
| IfCall       | Calls one of two functions based on a condition                        |
| IfCallReturn | Calls one of two functions based on a condition and returns the result |

### Containers

#### Lists

| List          | Description                                             |
|---------------|---------------------------------------------------------|
| New           | Creates a new doubly linked list from the given values |
| Init          | Initializes or resets the linked list to an empty state |
| Len           | Returns the number of elements in the list              |
| Front         | Returns the first element in the list or nil if empty   |
| Back          | Returns the last element in the list or nil if empty    |
| IsEmpty       | Checks if the list is empty                             |
| Remove        | Removes an element from the list and returns its value  |
| PushFront     | Inserts a value at the front of the list                |
| Prepend       | Adds multiple values to the front of the list           |
| PushBack      | Adds a value to the end of the list                     |
| Append        | Adds multiple values to the end of the list             |
| InsertBefore  | Inserts a value before a specified element              |
| InsertAfter   | Inserts a value after a specified element               |
| MoveToFront   | Moves an element to the front of the list               |
| MoveToBack    | Moves an element to the back of the list                |
| MoveBefore    | Moves an element before a specified mark                |
| MoveAfter     | Moves an element after a specified mark                 |
| PushBackList  | Appends all values from another list to the end         |
| PushFrontList | Prepends all values from another list to the front      |
| Values        | Returns a slice of all values in the list               |


| Element    | Description                                                     |
|------------|-----------------------------------------------------------------|
| NewElement | Creates a new element with the given value and parent list      |
| Next       | Returns the next element in the list or nil if at the end       |
| Prev       | Returns the previous element in the list or nil if at the start |
| Init          | Initializes or resets the linked list to an empty state |
| Len           | Returns the number of elements in the list              |
| Front         | Returns the first element in the list or nil if empty   |
| Back          | Returns the last element in the list or nil if empty    |
| IsEmpty       | Checks if the list is empty                             |
| Remove        | Removes an element from the list and returns its value  |
| PushFront     | Inserts a value at the front of the list                |
| Prepend       | Adds multiple values to the front of the list           |
| PushBack      | Adds a value to the end of the list                     |
| Append        | Adds multiple values to the end of the list             |
| InsertBefore  | Inserts a value before a specified element              |
| InsertAfter   | Inserts a value after a specified element               |
| MoveToFront   | Moves an element to the front of the list               |
| MoveToBack    | Moves an element to the back of the list                |
| MoveBefore    | Moves an element before a specified mark                |
| MoveAfter     | Moves an element after a specified mark                 |
| PushBackList  | Appends all values from another list to the end         |
| PushFrontList | Prepends all values from another list to the front      |
| Values        | Returns a slice of all values in the list               |


| Element    | Description                                                     |
|------------|-----------------------------------------------------------------|
| NewElement | Creates a new element with the given value and parent list      |
| Next       | Returns the next element in the list or nil if at the end       |
| Prev       | Returns the previous element in the list or nil if at the start |

#### Sets

| Set                 | Description                                                          |
|---------------------|----------------------------------------------------------------------|
| New                 | Creates a new set with the given items                               |
| Add                 | Adds one or more items to the set                                    |
| IsEmpty             | Returns true if the set contains zero items                          |
| Equals              | Returns true if two sets contain the same items                      |
| Contains            | Returns true if the set contains the specified item                  |
| Remove              | Removes the specified item from the set                              |
| Size                | Returns the number of items in the set                               |
| Clear               | Removes all items from the set                                       |
| Values              | Returns a slice of all values in the set                             |
| Intersect           | Returns a new set with items common to both sets                     |
| Union               | Returns a new set with all items from both sets                      |
| Difference          | Returns a new set with items in the current set but not in the other |
| SymmetricDifference | Returns a new set with items in either set but not both              |
| Subset              | Returns true if the set is a subset of the given set                 |
| Apply               | Applies a function to each item in the set                           |

### Stacks

| Stack  | Description                                                                     |
|--------|---------------------------------------------------------------------------------|
| New    | Creates a new stack using a linked list                                         |
| Push   | Adds a value to the top of the stack                                            |
| Pop    | Removes and returns the top element; returns default and false if empty         |
| Peek   | Returns the top element without removing it; returns default and false if empty |
| Size   | Returns the number of elements in the stack                                     |
| Values | Returns a slice of all elements in the stack                                    |

### Queues

| DQueue       | Description                                                                     |
|--------------|---------------------------------------------------------------------------------|
| NewQueue     | Creates a new double-ended queue with optional initial values                   |
| Enqueue      | Adds a value to the back of the queue                                           |
| EnqueueFront | Adds a value to the front of the queue                                          |
| Dequeue      | Removes and returns the front value; returns default and false if empty         |
| DequeueBack  | Removes and returns the back value; returns default and false if empty          |
| Peek         | Returns the front value without removing it; returns default and false if empty |
| PeekBack     | Returns the back value without removing it; returns default and false if empty  |
| IsEmpty      | Checks if the queue is empty                                                    |
| Size         | Returns the number of elements in the queue                                     |
| Clear        | Removes all elements from the queue                                             |
| Values       | Returns a slice of all values in the queue                                      |

## Roadmap

- [ ] Queues
- [ ] Pipeline
- [ ] Graphs

## License
MIT License, see [LICENSE](https://github.com/SharkByteSoftware/go-snk/blob/master/LICENSE) for details

// Package lists provides various linked list implementations
package lists

import (
	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/helpers"
)

// List represents a doubly linked list. API-compatible with Go's
// container/list implementation.
type List[T any] struct {
	root Element[T]
	len  int
}

// New creates a new list with the given values.
func New[T any](values ...T) *List[T] {
	l := &List[T]{}
	l.Init()
	l.Append(values...)

	return l
}

// Init initializes or clears list l.
func (l *List[T]) Init() *List[T] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List[T]) Len() int {
	return l.len
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// Remove removes e from l if e is an element of list l.
// It returns the element value.
// The element must not be nil.
func (l *List[T]) Remove(e *Element[T]) T {
	if e == nil {
		var zero T
		return zero
	}

	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// (current contract is that uninitialized list are not usable)
		e.prev.next = e.next
		e.next.prev = e.prev
		e.list = nil
		l.len--
	}
	return e.value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) PushFront(v T) *Element[T] {
	e := &Element[T]{
		value: v,
		list:  l,
	}

	e.prev = &l.root
	e.next = l.root.next
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

// Prepend inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) Prepend(values ...T) {
	for _, v := range values {
		l.PushFront(v)
	}
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List[T]) PushBack(v T) *Element[T] {
	e := &Element[T]{
		value: v,
		list:  l,
	}

	e.next = &l.root
	e.prev = l.root.prev
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

// Append inserts a new element e with value v at the back of list l and returns e.
func (l *List[T]) Append(values ...T) {
	for _, v := range values {
		l.PushBack(v)
	}
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}

	// see comment in List.Remove about initialization of l
	e := &Element[T]{
		value: v,
		list:  l,
	}

	e.prev = mark.prev
	e.next = mark
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}

	// see comment in List.Remove about initialization of l
	e := &Element[T]{
		value: v,
		list:  l,
	}
	e.next = mark.next
	e.prev = mark
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = &l.root
	e.next = l.root.next
	e.prev.next = e
	e.next.prev = e
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToBack(e *Element[T]) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	e.prev.next = e.next
	e.next.prev = e.prev

	e.next = &l.root
	e.prev = l.root.prev
	e.prev.next = e
	e.next.prev = e
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if e.list != l || mark.list != l || e == mark {
		return
	}
	// see comment in List.Remove about initialization of l
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = mark.prev
	e.next = mark
	e.prev.next = e
	e.next.prev = e
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e.list != l || mark.list != l || e == mark {
		return
	}
	// see comment in List.Remove about initialization of l
	e.prev.next = e.next
	e.next.prev = e.prev

	e.next = mark.next
	e.prev = mark
	e.prev.next = e
	e.next.prev = e
}

// PushFrontList inserts a copy of another list at the front of list l.
// After the operation, the elements of other will not be associated with l.
func (l *List[T]) PushFrontList(other *List[T]) {
	l.Prepend(other.Values()...)
}

// PushBackList pushes all the elements of the other list to the back of the list.
func (l *List[T]) PushBackList(other *List[T]) {
	l.Append(other.Values()...)
}

// IsEmpty returns true if the list is empty.
func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

// Size returns the size of the list.
func (l *List[T]) Size() int {
	return l.len
}

// Values returns a slice of all values in the list.
func (l *List[T]) Values() []T {
	if l.IsEmpty() {
		return nil
	}

	values := make([]T, 0, l.Len())

	for e := l.Front(); e != nil; e = e.Next() {
		values = append(values, e.Value())
	}

	return values
}

// Element is an element of a linked list.
type Element[T any] struct {
	// Next and previous pointers in the doubly linked list of elements.
	next, prev *Element[T]

	// The list to which this element belongs.
	list *List[T]

	// The value stored with this element.
	value T
}

// NewElement returns an element with the given value.
func NewElement[T any](value T) *Element[T] {
	return &Element[T]{value: value}
}

// Value returns the value of the list element.
func (e *Element[T]) Value() T {
	if e == nil {
		var zero T
		return zero
	}

	return e.value
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	conditional.If(e == nil, func() {
		return nil
	})

	if p := e.next; p != nil && p.list == e.list {
		return p
	}

	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	conditional.If(e == nil, func() {
		return nil
	})

	if p := e.prev; p != nil && p.list == e.list {
		return p
	}

	return nil
}

// Any returns true if any element in the list satisfies the condition.
func (l *List[T]) Any(condition func(T) bool) bool {
	if l.IsEmpty() {
		return false
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if condition(e.Value()) {
			return true
		}
	}

	return false
}

// All returns true if all elements in the list satisfy the condition.
func (l *List[T]) All(condition func(T) bool) bool {
	if l.IsEmpty() {
		return false
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if !condition(e.Value()) {
			return false
		}
	}

	return true
}

// Find returns the first element in the list that satisfies the condition.
func (l *List[T]) Find(condition func(T) bool) T {
	if l.IsEmpty() {
		var zero T
		return zero
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if condition(e.Value()) {
			return e.Value()
		}
	}

	var zero T
	return zero
}

// Apply applies the given function to each element in the list.
func (l *List[T]) Apply(f func(T)) {
	if l.IsEmpty() {
		return
	}

	for e := l.Front(); e != nil; e = e.Next() {
		f(e.Value())
	}
}

// Compact returns a new list with all non-zero elements from the original list.
func (l *List[T]) Compact() *List[T] {
	newList := New[T]()

	if l.IsEmpty() {
		return newList
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if !helpers.IsZeroValue(e.Value()) {
			newList.PushBack(e.Value())
		}
	}

	return newList
}

// Filter returns a new list with all elements that satisfy the condition.
func (l *List[T]) Filter(condition func(T) bool) *List[T] {
	newList := New[T]()

	if l.IsEmpty() {
		return newList
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if condition(e.Value()) {
			newList.PushBack(e.Value())
		}
	}

	return newList
}

## License
MIT License, see [LICENSE](https://github.com/SharkByteSoftware/go-snk/blob/master/LICENSE) for details



