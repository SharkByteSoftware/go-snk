<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# Queues

This document covers the queue / deque implementation in `containers`.

## Overview

The queue is a double-ended collection that supports operations at both ends.

It is useful when you need:
- enqueue/dequeue behavior
- access to both front and back
- a flexible queue-like structure
- a deque-style workflow

## Common operations

- create a new queue
- enqueue values at the back
- enqueue values at the front
- dequeue from the front
- dequeue from the back
- peek at the front
- peek at the back
- check whether the queue is empty
- inspect the size
- clear the queue
- retrieve all values

## API reference

### `Queue` / `DQueue`

| Method         | Purpose                                     |
|----------------|---------------------------------------------|
| `NewQueue`     | Creates a new double-ended queue            |
| `Enqueue`      | Adds a value to the back                    |
| `EnqueueFront` | Adds a value to the front                   |
| `Dequeue`      | Removes and returns the front value         |
| `DequeueBack`  | Removes and returns the back value          |
| `Peek`         | Returns the front value without removing it |
| `PeekBack`     | Returns the back value without removing it  |
| `IsEmpty`      | Reports whether the queue is empty          |
| `Size`         | Returns the number of elements              |
| `Clear`        | Removes all elements                        |
| `Values`       | Returns all values as a slice               |

## Notes

- This structure is useful when both ends matter.
- Handle empty queue operations carefully in calling code.

## When to use it

Use the queue when:
- items should be processed in arrival order
- you need double-ended access
- a deque is a better fit than a simple stack or slice

## Examples

