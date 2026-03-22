<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# queues

`queues` provides a double-ended queue with support for operations at both ends.

It is designed to help you manage ordered collections with small functions for:

- enqueueing values at the front or back
- dequeueing values from the front or back
- peeking at either end without removing values
- inspecting and clearing the queue

## Overview

Use `queues` when you want double-ended queue behavior that is easy to read, reuse, and test.

It is especially useful when:

- items should be processed in arrival order
- you need flexible access at both ends
- a deque is a better fit than a simple stack or slice

## When to use it

Use `queues` when:

- items should be processed in arrival order
- you need double-ended access
- a deque is a better fit than a simple stack or slice

Prefer a slice when:

- you only need access at one end
- random access by index is required
- the collection is small and structure does not matter

## API reference

### Create a queue

| Method     | Purpose                          |
|------------|----------------------------------|
| `NewQueue` | Creates a new double-ended queue |

### Add values

| Method         | Purpose                     |
|----------------|-----------------------------|
| `Enqueue`      | Adds a value to the back    |
| `EnqueueFront` | Adds a value to the front   |

### Remove values

| Method        | Purpose                             |
|---------------|-------------------------------------|
| `Dequeue`     | Removes and returns the front value |
| `DequeueBack` | Removes and returns the back value  |

### Inspect without removing

| Method     | Purpose                                     |
|------------|---------------------------------------------|
| `Peek`     | Returns the front value without removing it |
| `PeekBack` | Returns the back value without removing it  |

### Inspect or reset the queue

| Method    | Purpose                            |
|-----------|------------------------------------|
| `IsEmpty` | Reports whether the queue is empty |
| `Size`    | Returns the number of elements     |
| `Clear`   | Removes all elements               |
| `Values`  | Returns all values as a slice      |

## Notes

- Prefer the method that most clearly expresses your intent.
- This structure is most useful when both ends of the collection matter.
- Handle empty queue operations carefully in calling code; `Dequeue` and `Peek` return a default value and false when the queue is empty.

## Examples

Examples can be found in the [test suite](../../containers/queues/dqueue_test.go).