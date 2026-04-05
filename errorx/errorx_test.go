package errorx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/SharkByteSoftware/go-snk/errorx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errOne   = errors.New("error one")
	errTwo   = errors.New("error two")
	errThree = errors.New("error three")
)

func Test_Ignore(t *testing.T) {
	// confirms Ignore accepts an error and does not panic
	assert.NotPanics(t, func() { errorx.Ignore(errOne) })
	assert.NotPanics(t, func() { errorx.Ignore(nil) })
}

func Test_Must(t *testing.T) {
	// returns value when err is nil
	result := errorx.Must(42, nil)
	assert.Equal(t, 42, result)

	result = errorx.Must(0, nil)
	assert.Equal(t, 0, result)

	strResult := errorx.Must("hello", nil)
	assert.Equal(t, "hello", strResult)

	// panics when err is non-nil
	assert.Panics(t, func() { errorx.Must(0, errOne) })
	assert.PanicsWithValue(t, errOne, func() { errorx.Must("", errOne) })
}

func Test_IsAny(t *testing.T) {
	wrapped := fmt.Errorf("wrapped: %w", errOne)

	// matches a single target
	assert.True(t, errorx.IsAny(errOne, errOne))

	// matches one of multiple targets
	assert.True(t, errorx.IsAny(errOne, errTwo, errOne, errThree))

	// matches via errors.Is unwrapping
	assert.True(t, errorx.IsAny(wrapped, errOne))

	// no match
	assert.False(t, errorx.IsAny(errOne, errTwo, errThree))

	// nil error does not match non-nil targets
	assert.False(t, errorx.IsAny(nil, errOne, errTwo))

	// nil error matches nil target
	assert.True(t, errorx.IsAny(nil, nil))

	// no targets always returns false
	assert.False(t, errorx.IsAny(errOne))
}

func Test_FirstErr(t *testing.T) {
	// all nil returns nil
	assert.NoError(t, errorx.FirstErr())
	assert.NoError(t, errorx.FirstErr(nil))
	assert.NoError(t, errorx.FirstErr(nil, nil, nil))

	// single non-nil error
	require.ErrorIs(t, errorx.FirstErr(errOne), errOne)

	// returns first non-nil, ignores the rest
	require.ErrorIs(t, errorx.FirstErr(errOne, errTwo, errThree), errOne)
	require.ErrorIs(t, errorx.FirstErr(nil, errTwo, errThree), errTwo)
	require.ErrorIs(t, errorx.FirstErr(nil, nil, errThree), errThree)

	// nil after non-nil — still returns the first non-nil
	require.ErrorIs(t, errorx.FirstErr(errOne, nil), errOne)

	// wrapped errors are returned as-is, not unwrapped
	wrapped := fmt.Errorf("wrapped: %w", errOne)
	result := errorx.FirstErr(nil, wrapped, errTwo)
	require.ErrorIs(t, result, errOne) // unwraps correctly via errors.Is
	assert.NotErrorIs(t, result, errTwo)
}
