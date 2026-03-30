// Package stringx provides small helpers for common string operations.
package stringx

import (
	"strings"

	"github.com/SharkByteSoftware/go-snk/slicex"
)

// IsBlank returns true if the string is empty or contains only whitespace.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Coalesce returns the first non-empty string from the provided values.
// If all values are empty, an empty string is returned.
func Coalesce(values ...string) string {
	return slicex.FindOrBy(values, func(value string) bool { return value != "" }, "")
}

// Truncate returns the string trimmed to a maximum length.
// If the string is shorter than or equal to max, it is returned unchanged.
// If max is less than or equal to 0, an empty string is returned.
func Truncate(str string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}

	runes := []rune(str)
	if len(runes) <= maxLen {
		return str
	}

	return string(runes[:maxLen])
}

// Wrap surrounds the string with the given prefix and suffix.
func Wrap(s, prefix, suffix string) string {
	return prefix + s + suffix
}

// PadLeft returns the string left-padded with char to the specified length.
// If the string is already at or longer than length, it is returned unchanged.
func PadLeft(s string, length int, char rune) string {
	runes := []rune(s)
	for len(runes) < length {
		runes = append([]rune{char}, runes...)
	}

	return string(runes)
}

// PadRight returns the string right-padded with char to the specified length.
// If the string is already at or longer than length, it is returned unchanged.
func PadRight(s string, length int, char rune) string {
	runes := []rune(s)
	for len(runes) < length {
		runes = append(runes, char)
	}

	return string(runes)
}
