package jsonx

import "github.com/SharkByteSoftware/go-snk/slicex"

// EncodeOption is a function that configures EncodeOptions.
type EncodeOption func(options *EncodeOptions)

// EncodeOptions contains the configuration options for JSON encoding.
type EncodeOptions struct {
	escapeHTML bool
	indent     string
}

func newEncodeOptions(options []EncodeOption) *EncodeOptions {
	cfg := &EncodeOptions{
		escapeHTML: false,
		indent:     "",
	}

	slicex.Apply(options, func(option EncodeOption) { option(cfg) })

	return cfg
}

// WithEscapeHTML enables escaping of HTML characters (<, >, &) in the output.
// By default HTML escaping is disabled, unlike the standard library which
// enables it. Disable only when you are certain the output will not be
// embedded in HTML.
func WithEscapeHTML() EncodeOption {
	return func(options *EncodeOptions) {
		options.escapeHTML = true
	}
}

// WithIndent enables pretty-printing with the given indent string per level.
// Common values are "\t" or "  " (two spaces).
func WithIndent(indent string) EncodeOption {
	return func(options *EncodeOptions) {
		options.indent = indent
	}
}
