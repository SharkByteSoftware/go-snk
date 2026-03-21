<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# Lists

This document covers the linked list implementation in `containers`.

## Overview

The list implementation provides a doubly linked list with a familiar API and container-style behavior.

It is useful when you need:
- ordered insertion and removal
- direct access to the front or back
- element movement within the list
- list-to-list copying behavior

## Core concepts

### List
The list type represents the collection itself and manages the linked structure.

### Element
An element represents a node in the list and can be used to navigate forward and backward.

## Common operations

- create a new list
- initialize or reset a list
- inspect the length
- get the front or back element
- check whether the list is empty
- insert values at the front or back
- insert values before or after a given element
- move elements within the list
- remove elements
- clear the list
- retrieve all values as a slice

### `List`

| Method          | Purpose                                  |
|-----------------|------------------------------------------|
| `New`           | Creates a new list from initial values   |
| `Init`          | Initializes or resets the list           |
| `Len`           | Returns the number of elements           |
| `Size`          | Returns the number of elements           |
| `Front`         | Returns the first element                |
| `Back`          | Returns the last element                 |
| `IsEmpty`       | Reports whether the list is empty        |
| `Remove`        | Removes an element and returns its value |
| `PushFront`     | Inserts a value at the front             |
| `Prepend`       | Inserts multiple values at the front     |
| `PushBack`      | Inserts a value at the back              |
| `Append`        | Inserts multiple values at the back      |
| `InsertBefore`  | Inserts a value before a mark element    |
| `InsertAfter`   | Inserts a value after a mark element     |
| `MoveToFront`   | Moves an element to the front            |
| `MoveToBack`    | Moves an element to the back             |
| `MoveBefore`    | Moves an element before another element  |
| `MoveAfter`     | Moves an element after another element   |
| `PushBackList`  | Appends all values from another list     |
| `PushFrontList` | Prepends all values from another list    |
| `Values`        | Returns all values as a slice            |
| `Clear`         | Removes all elements from the list       |

## Notes

- Elements are only valid within the list they belong to.
- Mutating a list may affect element references.
- Use the list’s own methods to manage membership and movement.

## When to use it

Use the list when:
- you need stable insertion and removal behavior at either end
- you want to move elements without rebuilding the whole collection
- a linked structure is a better fit than a slice

## API reference


## Examples

See [`docs/examples.md`](../examples.md) for general examples. Add list-specific examples here if needed later.