<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# stacks

`stacks` provides a last-in, first-out collection backed by a linked structure.

It is designed to help you manage ordered temporary state with small functions for:

- pushing and popping values
- peeking at the top value without removing it
- inspecting the stack's contents and size

## Overview

Use `stacks` when you want LIFO collection behavior that is easy to read, reuse, and test.

It is especially useful when:

- the most recently added item should be processed first
- the code naturally follows a LIFO pattern
- you want a simple abstraction over a linked collection

## When to use it

Use `stacks` when:

- the most recently added item should be processed first
- the code naturally follows a LIFO pattern
- you want a simple abstraction over a linked collection

Prefer a slice when:

- you need access at both ends
- random access by index is required
- the collection is small and structure does not matter

## API reference

### Create a stack

| Method | Purpose             |
|--------|---------------------|
| `New`  | Creates a new stack |

### Add or remove values

| Method | Purpose                           |
|--------|-----------------------------------|
| `Push` | Pushes a value onto the stack     |
| `Pop`  | Removes and returns the top value |

### Inspect without removing

| Method | Purpose                                   |
|--------|-------------------------------------------|
| `Peek` | Returns the top value without removing it |

### Inspect the stack

| Method   | Purpose                        |
|----------|--------------------------------|
| `Size`   | Returns the number of elements |
| `Values` | Returns all values as a slice  |

## Notes

- Prefer the method that most clearly expresses your intent.
- A stack is a natural fit for LIFO workflows.
- `Pop` and `Peek` return a default value and false when the stack is empty; handle these cases carefully in calling code.

## Examples

Examples can be found in the  [test suite](../../containers/stacks/stack_test.go).