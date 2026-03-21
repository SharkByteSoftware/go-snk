<div style="text-align: center;">
    <a href="https://github.com/SharkByteSoftware/go-snk">
        <img width="" height="202" alt="sink-logo" src="../../img/logo.png" />
    </a>
</div>

# Stacks

This document covers the stack implementation in `containers`.

## Overview

The stack provides a last-in, first-out collection backed by a linked structure.

It is useful when you need:
- push/pop behavior
- a simple LIFO collection
- a compact way to manage ordered temporary state

## Common operations

- create a new stack
- push values
- pop the top value
- peek at the top value
- inspect the stack size
- retrieve all values

## API reference

### `Stack`

| Method   | Purpose                                   |
|----------|-------------------------------------------|
| `New`    | Creates a new stack                       |
| `Push`   | Pushes a value onto the stack             |
| `Pop`    | Removes and returns the top value         |
| `Peek`   | Returns the top value without removing it |
| `Size`   | Returns the number of elements            |
| `Values` | Returns all values as a slice             |

## Notes

- A stack is a natural fit for LIFO workflows.
- Empty stack operations should be handled carefully by callers.

## When to use it

Use the stack when:
- the most recently added item should be processed first
- the code naturally follows a LIFO pattern
- you want a simple abstraction over a linked collection

## Examples

See [`docs/examples.md`](../examples.md) for general examples.