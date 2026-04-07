package jsonx_test

import (
	"bytes"
	"testing"

	"github.com/SharkByteSoftware/go-snk/jsonx"
)

type benchFields struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Tags    []string `json:"tags"`
	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
	Meta map[string]any `json:"meta"`
}

var (
	benchJSON  = `{"name":"Alice","age":30,"tags":["go","json","benchmark"],"address":{"city":"Berlin","country":"Germany"},"meta":{"id":123,"active":true}}`
	benchBytes = []byte(benchJSON)
)

func BenchmarkDecode(b *testing.B) {
	b.Run("Decode T", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_, _ = jsonx.Decode[benchFields](bytes.NewReader(benchBytes))
		}
	})

	b.Run("Decode *T", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_, _ = jsonx.DecodePtr[benchFields](bytes.NewReader(benchBytes))
		}
	})
}
