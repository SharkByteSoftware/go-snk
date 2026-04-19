// Package errorx provides small helpers for common error handling patterns.
package errorx

import (
	"errors"

	"github.com/SharkByteSoftware/go-snk/slicex"
)

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
	return slicex.AnyBy(targets, func(target error) bool {
		return errors.Is(err, target)
	})
}

// FirstErr returns the first non-nil error from the provided errors,
// or nil if all errors are nil.
// It is useful for reducing a set of validation or initialization errors
// to a single result without chaining multiple if statements.
func FirstErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
