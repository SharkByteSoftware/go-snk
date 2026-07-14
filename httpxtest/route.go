package httpxtest

import (
	"net/http"
	"sync/atomic"
)

// ExhaustBehavior controls how a sequence of responses is exhausted.
type ExhaustBehavior int

const (
	// ExhaustCycle wraps around to the first entry.
	ExhaustCycle ExhaustBehavior = iota

	// ExhaustRepeatLast repeats the final entry indefinitely.
	ExhaustRepeatLast ExhaustBehavior = iota

	// ExhaustServerError returns 500 (default).
	ExhaustServerError
)

type routeEntry struct {
	handlers []http.HandlerFunc
	exhaust  ExhaustBehavior
	index    atomic.Int64
}

func (e *routeEntry) next() http.HandlerFunc {
	i := int(e.index.Add(1) - 1)
	n := len(e.handlers)

	switch e.exhaust {
	case ExhaustRepeatLast:
		i = min(i, n-1)
	case ExhaustCycle:
		i %= n
	case ExhaustServerError:
		if i >= n {
			return defaultHandler
		}
	}

	return e.handlers[i]
}