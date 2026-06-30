package httpxtest

import (
	"net/http"
	"sync/atomic"
)

// ExhaustBehavior controls what happens when a sequence runs out of responses.
type ExhaustBehavior int

const (
	ExhaustRepeatLast  ExhaustBehavior = iota // repeat the final entry indefinitely
	ExhaustCycle                              // wrap around to the first entry
	ExhaustServerError                        // return 500 (default)
)

// routeEntry holds a sequence of handlers for a single method/route key.
// A plain OnRoute registers a single-entry sequence.
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
		i = i % n
	default: // ExhaustServerError
		if i >= n {
			return defaultHandler
		}
	}

	return e.handlers[i]
}