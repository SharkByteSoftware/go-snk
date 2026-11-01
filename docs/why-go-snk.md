## Why go-snk?

Go's standard library now covers a lot of ground (`slices`, `maps`, `cmp`, `iter`), and
[`samber/lo`](https://github.com/samber/lo) is the de-facto "lodash for Go." go-snk isn't
trying to replace either. It's a smaller, opinionated toolkit that grew out of production
backend services.

**Use the stdlib first.** If `slices.Contains`, `slices.SortFunc`, or `maps.Keys` does the
job, use it. go-snk doesn't wrap what the stdlib already does well, and it fills the gaps
the stdlib deliberately leaves out: `Map`, `Filter`, `Reduce`, `GroupBy`, and friends.

**How it differs from `lo`:**

- **Consistent constraints.** Every `slicex` function takes `S ~[]T` and every `mapx`
  function takes `M ~map[K]V`, so named slice and map types work everywhere with no conversions.
- **Error-aware parallelism.** `parallel.MapErr` and `parallel.MapErrWithLimit` run on a
  bounded worker pool, preserve input order, and fail fast, the same contract as
  sequential `slicex.MapErr`. Collect-all-errors behavior is a separate, explicitly named
  function, never a flag. <!-- verify: lo/parallel has no error-returning or concurrency-limited Map -->
- **Backend batteries.** Beyond collections: `jsonx` (validation, file encode/decode),
  `httpx` (client options such as auth helpers), and `httpxtest` for testing HTTP clients
  against scripted responses.
- **Focused packages over one giant namespace.** `slicex.Filter`, `mapx.Keys`,
  `stringx.ToSnakeCase`: the import tells you what you're working with.
- **Containers included.** Generic sets, queues, stacks, lists, and priority queues
  ship with the library.

**When to pick `lo` instead:** if you want the broadest possible surface area, a large
user base, and lots of community examples, `lo` is the safer bet. go-snk is narrower
and more opinionated.

**When plain Go is better:** a simple `for` loop is often clearer than a chain of
helpers. go-snk is for the cases where a helper removes real boilerplate, not for
replacing every loop.
