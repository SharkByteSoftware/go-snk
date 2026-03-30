// Package errorx provides small helpers for common error handling patterns.
package errorx

import "errors"

// Ignore explicitly discards an error. It is intended to document
// intentional error suppression rather than silently assigning to _.
func Ignore(_ error) {}

// Must returns the value if err is nil, and panics otherwise.
// It is intended for use at program initialization time when
// an error represents a non-recoverable misconfiguration.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// IsAny reports whether err matches any of the provided targets,
// using errors.Is semantics for each comparison.
func IsAny(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}
