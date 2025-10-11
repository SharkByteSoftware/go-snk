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

**go-snk**: A versatile Go utility library to streamline and clean up your projects with a single import. Powered by Go generics, 
it offers type-safe, flexible tools for data manipulation, including functional slice operations (`Filter`, `Map`, `Reduce`, `UniqueMap`), 
aggregation functions (`Sum`, `Mean`, `Max`), map management (`Keys`, `Values`, `Invert`, `Combine`), set operations (`Union`, `Intersect`, 
`SymmetricDifference`), and a generic doubly linked list (`PushFront`, `InsertAfter`, `MoveToBack`). Simplify your codebase and boost 
productivity with go-snk's all-in-one toolkit.

## Getting Started

### Prerequisites

- **Go version**: go-snk uses [Go](https://go.dev/) version [1.24](https://go.dev/doc/devel/release#go1.24.0) or above
- **Basic Go knowledge**: Familiarity with Go syntax and package management is helpful

### Installation

With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), you just import `go-snk` and Go will automatiically fetch it during build:

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

| Helpers   | Description                                                  |
|-----------|--------------------------------------------------------------|
| Filter    | Filters a slice using a predicate                            |
| Map       | Transforms a slice to a slice of another type using a mapper |
| UniqueMap | Similar to Map but removes duplicates from the result        |
| Reduce    | Transforms and flattens a slice to another type              |
| Find      | Returns the first element matching a condition               |
| FindBy    | Returns the first element matching a predicate function      |
| FindOr    | Returns the first matching element or a default value        |
| Any       | Checks if any element satisfies a condition                  |
| AnyBy     | Checks if any element satisfies a predicate function         |
| All       | Checks if all elements satisfy a condition                   |
| Unique    | Removes duplicates from a slice                              |
| Apply     | Applies a function to each element in a slice                |
| Reverse   | Reverses the order of elements in a slice                    |
| Compact   | Compact returns a slice with all the non-zero items.         |
| ToMap     | Converts a slice to a map using a key selector function      |
| GroupBy   | Groups elements of a slice by a key selector function        |
| Partition | Splits a slice into two based on a predicate function        |

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

| Function  | Description                                                                                         |
|-----------|-----------------------------------------------------------------------------------------------------|
| Map       | Transforms a slice to a slice of another type using a mapper function in parallel, preserving order |
| Apply     | Applies a function to each item in a slice in parallel                                              |
| GroupBy   | Groups a slice into a map of slices based on a predicate function in parallel                       |
| Partition | Splits a slice into two slices based on a predicate function in parallel, preserving order          |

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

## Conditionals

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
| New           | Creates a new doubly linked list from the given values  |
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



