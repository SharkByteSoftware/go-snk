package stringx_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/stringx"
	"github.com/stretchr/testify/assert"
)

func Test_IsBlank(t *testing.T) {
	assert.True(t, stringx.IsBlank(""))
	assert.True(t, stringx.IsBlank("   "))
	assert.True(t, stringx.IsBlank("\t\n"))

	assert.False(t, stringx.IsBlank("a"))
	assert.False(t, stringx.IsBlank("  a  "))
}

func Test_Coalesce(t *testing.T) {
	// returns first non-empty value
	assert.Equal(t, "first", stringx.Coalesce("first", "second"))
	assert.Equal(t, "second", stringx.Coalesce("", "second"))
	assert.Equal(t, "third", stringx.Coalesce("", "", "third"))

	// all empty returns empty
	assert.Empty(t, stringx.Coalesce("", "", ""))

	// no args returns empty
	assert.Empty(t, stringx.Coalesce())

	// single non-empty value
	assert.Equal(t, "only", stringx.Coalesce("only"))
}

func Test_Truncate(t *testing.T) {
	// shorter than max — returned unchanged
	assert.Equal(t, "hello", stringx.Truncate("hello", 10))

	// equal to max — returned unchanged
	assert.Equal(t, "hello", stringx.Truncate("hello", 5))

	// longer than max — trimmed
	assert.Equal(t, "hel", stringx.Truncate("hello", 3))

	// max of 0 — returns empty
	assert.Empty(t, stringx.Truncate("hello", 0))

	// negative max — returns empty
	assert.Empty(t, stringx.Truncate("hello", -1))

	// empty input
	assert.Empty(t, stringx.Truncate("", 5))

	// multibyte characters — truncates by rune, not byte
	assert.Equal(t, "héll", stringx.Truncate("héllo", 4))
}

func Test_Wrap(t *testing.T) {
	assert.Equal(t, "(hello)", stringx.Wrap("hello", "(", ")"))
	assert.Equal(t, `"hello"`, stringx.Wrap("hello", `"`, `"`))
	assert.Equal(t, "<div>hello</div>", stringx.Wrap("hello", "<div>", "</div>"))

	// empty prefix and suffix — unchanged
	assert.Equal(t, "hello", stringx.Wrap("hello", "", ""))

	// empty input
	assert.Equal(t, "()", stringx.Wrap("", "(", ")"))
}

func Test_PadLeft(t *testing.T) {
	assert.Equal(t, "  hello", stringx.PadLeft("hello", 7, ' '))
	assert.Equal(t, "00042", stringx.PadLeft("42", 5, '0'))

	// already at length — unchanged
	assert.Equal(t, "hello", stringx.PadLeft("hello", 5, ' '))

	// longer than length — unchanged
	assert.Equal(t, "hello", stringx.PadLeft("hello", 3, ' '))

	// empty input
	assert.Equal(t, "   ", stringx.PadLeft("", 3, ' '))
}

func Test_PadRight(t *testing.T) {
	assert.Equal(t, "hello  ", stringx.PadRight("hello", 7, ' '))
	assert.Equal(t, "42000", stringx.PadRight("42", 5, '0'))

	// already at length — unchanged
	assert.Equal(t, "hello", stringx.PadRight("hello", 5, ' '))

	// longer than length — unchanged
	assert.Equal(t, "hello", stringx.PadRight("hello", 3, ' '))

	// empty input
	assert.Equal(t, "   ", stringx.PadRight("", 3, ' '))
}
