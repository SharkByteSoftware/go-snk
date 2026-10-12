package parallel

import (
	"sync"
	"sync/atomic"
)

// runIndexed calls apply once for every element of items using at most concurrency goroutines.
//
// If apply panics, the first panic value is captured, remaining unstarted work is skipped, and the
// panic is re-raised on the caller's goroutine once all workers have stopped. This lets callers
// recover from a panicking callback exactly as they could with the sequential slicex functions;
// a panic on a bare worker goroutine would otherwise terminate the whole process.
func runIndexed[S ~[]T, T any](items S, concurrency int, apply func(idx int, item T)) {
	if len(items) == 0 {
		return
	}

	concurrency = max(min(concurrency, len(items)), 1)

	tasks := make(chan int, len(items))
	for idx := range items {
		tasks <- idx
	}

	close(tasks)

	var (
		waitGroup sync.WaitGroup
		aborted   atomic.Bool
		once      sync.Once
		recovered any
	)

	for range concurrency {
		waitGroup.Go(func() {
			defer func() {
				if r := recover(); r != nil {
					once.Do(func() { recovered = r })
					aborted.Store(true)
				}
			}()

			for idx := range tasks {
				if aborted.Load() {
					return
				}

				apply(idx, items[idx])
			}
		})
	}

	waitGroup.Wait()

	if aborted.Load() {
		panic(recovered)
	}
}
