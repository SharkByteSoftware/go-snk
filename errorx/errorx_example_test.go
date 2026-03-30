package errorx_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/SharkByteSoftware/go-snk/errorx"
)

func ExampleIgnore() {
	// Use Ignore to explicitly document that an error is intentionally discarded.
	errorx.Ignore(os.Remove("file-that-may-not-exist.tmp"))

	fmt.Println("done")
	// Output: done
}

func ExampleMust() {
	parse := func(s string) (int, error) {
		if s == "" {
			return 0, errors.New("empty input")
		}

		return len(s), nil
	}

	result := errorx.Must(parse("hello"))
	fmt.Println(result)
	// Output: 5
}

func ExampleIsAny() {
	var (
		errNotFound   = errors.New("not found")
		errForbidden  = errors.New("forbidden")
		errBadRequest = errors.New("bad request")
	)

	err := fmt.Errorf("request failed: %w", errForbidden)

	if errorx.IsAny(err, errNotFound, errForbidden) {
		fmt.Println("client error")
	}

	fmt.Println(errorx.IsAny(err, errBadRequest))
	// Output:
	// client error
	// false
}
