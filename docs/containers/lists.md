<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# lists

`lists` provides a doubly linked list with a familiar API and container-style behavior.

It is designed to help you manage ordered collections with small functions for:

- inserting and removing values at either end or around a given element
- moving elements within the list
- inspecting the list's contents and structure
- copying values between lists

## Overview

Use `lists` when you want linked list behavior that is easier to read, reuse, and test.

It is especially useful when:

- you need stable insertion and removal at either end without shifting elements
- you want to move elements within the collection without rebuilding it
- a linked structure is a better fit than a slice

## Core concepts

### List
The `List` type represents the collection itself and manages the linked structure.

### Element
An `Element` represents a node in the list and can be used to navigate forward and backward.

## When to use it

Use `lists` when:

- you need stable insertion and removal behavior at either end
- you want to move elements without rebuilding the whole collection
- a linked structure is a better fit than a slice

Prefer a slice when:

- random access by index is required
- the collection is small and insertion order does not matter
- allocation simplicity is more important than insertion performance

## API reference

### Create or reset a list

| Method  | Purpose                                |
|---------|----------------------------------------|
| `New`   | Creates a new list from initial values |
| `Init`  | Initializes or resets the list         |
| `Clear` | Removes all elements from the list     |

### Inspect the list

| Method    | Purpose                               |
|-----------|---------------------------------------|
| `Len`     | Returns the number of elements        |
| `Size`    | Alias for `Len`                       |
| `Front`   | Returns the first element             |
| `Back`    | Returns the last element              |
| `IsEmpty` | Reports whether the list is empty     |
| `Values`  | Returns all values as a slice         |

### Insert values

| Method          | Purpose                                   |
|-----------------|-------------------------------------------|
| `PushFront`     | Inserts a value at the front              |
| `Prepend`       | Inserts multiple values at the front      |
| `PushBack`      | Inserts a value at the back               |
| `Append`        | Inserts multiple values at the back       |
| `InsertBefore`  | Inserts a value before a mark element     |
| `InsertAfter`   | Inserts a value after a mark element      |
| `PushBackList`  | Appends all values from another list      |
| `PushFrontList` | Prepends all values from another list     |

### Move or remove elements

| Method        | Purpose                                  |
|---------------|------------------------------------------|
| `Remove`      | Removes an element and returns its value |
| `MoveToFront` | Moves an element to the front            |
| `MoveToBack`  | Moves an element to the back             |
| `MoveBefore`  | Moves an element before another element  |
| `MoveAfter`   | Moves an element after another element   |

### Visit elements

| Method     | Purpose                                    |
|------------|--------------------------------------------|
| `ForEach`  | Calls a function on each element in order  |

### Navigate elements

| Method | Purpose                                                          |
|--------|------------------------------------------------------------------|
| `Next` | Returns the next element in the list, or nil if at the end       |
| `Prev` | Returns the previous element in the list, or nil if at the start |

## Notes

- Prefer the method that most clearly expresses your intent.
- Elements are only valid within the list they belong to.
- Mutating a list may affect element references.
- Use the list's own methods to manage membership and movement.

## Examples

Examples can be found in the [test suite](../../containers/lists/list_test.go).