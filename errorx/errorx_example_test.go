//nolint:err113
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

func ExampleFirstErr() {
	validate := func(name string, age int) error {
		return errorx.FirstErr(
			validateName(name),
			validateAge(age),
		)
	}

	fmt.Println(validate("alice", 30))
	fmt.Println(validate("", 30))
	// Output:
	// <nil>
	// name is required
}

func ExampleFirstErr_allNil() {
	err := errorx.FirstErr(nil, nil, nil)

	fmt.Println(err)
	// Output: <nil>
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	return nil
}

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age must be non-negative")
	}

	return nil
}
