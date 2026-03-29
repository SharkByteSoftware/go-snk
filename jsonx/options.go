package jsonx

import "github.com/SharkByteSoftware/go-snk/slicex"

// Option is a function that configures DecodeOptions.
type Option func(options *DecodeOptions)

// DecodeOptions contains the configuration options for JSON decoding.
type DecodeOptions struct {
	strictDecoding bool
	useNumber      bool
}

func newDecodeOptions(options []Option) *DecodeOptions {
	cfg := &DecodeOptions{
		strictDecoding: false,
		useNumber:      false,
	}

	slicex.Apply(options, func(option Option) { option(cfg) })

	return cfg
}

// WithStrictDecoding disallows unknown fields when decoding JSON.
func WithStrictDecoding() Option {
	return func(options *DecodeOptions) {
		options.strictDecoding = true
	}
}

// WithUseNumber causes the decoder to unmarshal JSON numbers as [json.Number]
// instead of float64. This is useful when precision matters or when the caller
// wants to determine the numeric type themselves.
func WithUseNumber() Option {
	return func(options *DecodeOptions) {
		options.useNumber = true
	}
}
