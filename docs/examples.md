go # Examples

This page collects short, practical examples that show the style of the library.

## Working with slices

Use `slicex` when you want to transform or query slices without repeatedly writing the same loop.

Common use cases include:
- filtering values
- mapping one type to another
- finding values
- grouping data
- deduplicating slices

## Working with maps

Use `mapx` when you want helpers for common map tasks such as:
- getting keys or values
- filtering entries
- converting a map into another structure
- combining multiple maps

## Working with HTTP

Use `httpx` when you want a smaller, more convenient layer over HTTP request handling.

Typical use cases include:
- sending requests
- decoding responses into typed values
- working with request options
- keeping client code concise

## Working with conditional helpers

Use `conditional` when you want a compact way to express:
- choosing between two values
- calling a function only when needed
- executing one of two callbacks based on a boolean condition

## Working with containers

Use `containers` when you want reusable collection types instead of building them from scratch.

Common use cases include:
- maintaining ordered data with a list
- tracking unique items with a set
- building a stack
- using a double-ended queue

## A few practical guidelines

- Prefer the standard library when it is already the simplest option.
- Use `go-snk` when a helper improves clarity or removes repetitive code.
- Keep examples small and focused on one concept at a time.

## Suggested structure for examples

A good example usually shows:
1. the problem
2. the helper being used
3. the result

That keeps examples readable and useful without turning them into mini tutorials.