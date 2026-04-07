package jsonx_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/SharkByteSoftware/go-snk/jsonx"
)

type benchFields struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	benchJSON  = `{"name":"Alice","age":30}`
	benchBytes = []byte(benchJSON)
)

func BenchmarkDecode(b *testing.B) {
	b.Run("Decode", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_, _ = jsonx.Decode[benchFields](bytes.NewReader(benchBytes))
		}
	})

	b.Run("DecodeBytes", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_, _ = jsonx.DecodeBytes[benchFields](benchBytes)
		}
	})

	b.Run("DecodeString", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_, _ = jsonx.DecodeString[benchFields](benchJSON)
		}
	})
}

func BenchmarkDecode_WithStrictDecoding(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = jsonx.Decode[benchFields](strings.NewReader(benchJSON), jsonx.WithStrictDecoding())
	}
}
