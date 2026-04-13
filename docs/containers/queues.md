<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# queues

`queues` provides two queue implementations: a double-ended queue (`Queue`) and a priority queue (`PriorityQueue`).

## Queue

`Queue` supports operations at both ends and is designed to help you manage ordered collections with small functions for:

- enqueueing values at the front or back
- dequeueing values from the front or back
- peeking at either end without removing values
- inspecting and clearing the queue

### Overview

Use `Queue` when you want double-ended queue behavior that is easy to read, reuse, and test.

It is especially useful when:

- items should be processed in arrival order
- you need flexible access at both ends
- a deque is a better fit than a simple stack or slice

### When to use it

Use `Queue` when:

- items should be processed in arrival order
- you need double-ended access
- a deque is a better fit than a simple stack or slice

Prefer a slice when:

- you only need access at one end
- random access by index is required
- the collection is small and structure does not matter

### API reference

#### Create a queue

| Method     | Purpose                          |
|------------|----------------------------------|
| `NewQueue` | Creates a new double-ended queue |

#### Add values

| Method         | Purpose                     |
|----------------|-----------------------------|
| `Enqueue`      | Adds a value to the back    |
| `EnqueueFront` | Adds a value to the front   |

#### Remove values

| Method        | Purpose                             |
|---------------|-------------------------------------|
| `Dequeue`     | Removes and returns the front value |
| `DequeueBack` | Removes and returns the back value  |

#### Inspect without removing

| Method     | Purpose                                     |
|------------|---------------------------------------------|
| `Peek`     | Returns the front value without removing it |
| `PeekBack` | Returns the back value without removing it  |

#### Inspect or reset the queue

| Method    | Purpose                            |
|-----------|------------------------------------|
| `IsEmpty` | Reports whether the queue is empty |
| `Size`    | Returns the number of elements     |
| `Clear`   | Removes all elements               |
| `Values`  | Returns all values as a slice      |

### Notes

- Prefer the method that most clearly expresses your intent.
- This structure is most useful when both ends of the collection matter.
- Handle empty queue operations carefully in calling code; `Dequeue` and `Peek` return a default value and false when the queue is empty.

### Examples

Examples can be found in the [test suite](../../containers/queues/queue_test.go).

---

## PriorityQueue

`PriorityQueue` is a generic min-heap backed priority queue. Elements are ordered by a custom comparator, so it works with any type.

The comparator follows the same contract as `slices.SortFunc`: return a negative value if `a` has higher priority than `b`, positive if `b` has higher priority, and zero if they are equal.

### Overview

Use `PriorityQueue` when you want elements dequeued in a defined priority order rather than arrival order.

It is especially useful when:

- tasks or events have varying urgency
- you need efficient O(log n) insert and removal
- a sorted slice would be too costly to maintain on each insert

### When to use it

Use `PriorityQueue` when:

- elements must be processed by priority, not insertion order
- you need O(1) peek at the highest-priority element
- the ordering rule can be expressed as a comparator function

Prefer a slice when:

- you need random access by index
- the collection is small and a linear scan is acceptable
- insertion order is all that matters

### API reference

#### Create a priority queue

| Method                        | Purpose                                                      |
|-------------------------------|--------------------------------------------------------------|
| `NewPriorityQueue`            | Creates a priority queue pre-populated with the given items  |
| `NewPriorityQueueWithDefault` | Creates an empty priority queue with the given comparator    |

#### Add values

| Method    | Purpose                                              |
|-----------|------------------------------------------------------|
| `Enqueue` | Inserts a new element in O(log n) time               |

#### Remove values

| Method    | Purpose                                                        |
|-----------|----------------------------------------------------------------|
| `Dequeue` | Removes and returns the highest-priority element in O(log n)  |

#### Inspect without removing

| Method  | Purpose                                                      |
|---------|--------------------------------------------------------------|
| `Peek`  | Returns the highest-priority element without removing it in O(1) |

#### Inspect or reset the queue

| Method    | Purpose                                         |
|-----------|-------------------------------------------------|
| `IsEmpty` | Reports whether the priority queue is empty     |
| `Len`     | Returns the number of elements                  |
| `Size`    | Returns the number of elements (alias for `Len`) |
| `Clear`   | Removes all elements                            |
| `Values`  | Returns a snapshot of all elements (unordered)  |

### Notes

- The comparator determines priority: if `comparator(a, b) < 0`, then `a` is dequeued before `b`.
- `Values` returns elements in heap order, not priority order. Use repeated `Dequeue` calls to retrieve elements in priority order.
- `NewPriorityQueue` heapifies the provided slice in O(n) time; the slice does not need to be sorted.
- Handle empty queue operations carefully; `Dequeue` and `Peek` return a default value and false when the queue is empty.

### Examples

Examples can be found in the [test suite](../../containers/queues/priority_queue_test.go).