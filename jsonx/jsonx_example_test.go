package jsonx_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/SharkByteSoftware/go-snk/jsonx"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Encode

func ExampleEncode() {
	var buf strings.Builder

	_ = jsonx.Encode(&buf, person{Name: "Alice", Age: 30})

	fmt.Print(buf.String())
	// Output: {"name":"Alice","age":30}
}

func ExampleEncode_withIndent() {
	var buf strings.Builder

	_ = jsonx.Encode(&buf, person{Name: "Alice", Age: 30}, jsonx.WithIndent("  "))

	fmt.Print(buf.String())
	// Output:
	// {
	//   "name": "Alice",
	//   "age": 30
	// }
}

func ExampleEncode_withEscapeHTML() {
	type link struct {
		URL string `json:"url"`
	}

	var buf strings.Builder

	_ = jsonx.Encode(&buf, link{URL: "https://example.com?a=1&b=2"}, jsonx.WithEscapeHTML())

	fmt.Print(buf.String())
	// Output: {"url":"https://example.com?a=1\u0026b=2"}
}

func ExampleEncodeBytes() {
	b, _ := jsonx.EncodeBytes(person{Name: "Alice", Age: 30})

	fmt.Print(string(b))
	// Output: {"name":"Alice","age":30}
}

func ExampleEncodeString() {
	s, _ := jsonx.EncodeString(person{Name: "Alice", Age: 30})

	fmt.Print(s)
	// Output: {"name":"Alice","age":30}
}

// Decode

func ExampleDecode() {
	r := strings.NewReader(`{"name":"Alice","age":30}`)

	p, _ := jsonx.Decode[person](r)

	fmt.Println(p.Name, p.Age)
	// Output: Alice 30
}

func ExampleDecodeBytes() {
	p, _ := jsonx.DecodeBytes[person]([]byte(`{"name":"Alice","age":30}`))

	fmt.Println(p.Name, p.Age)
	// Output: Alice 30
}

func ExampleDecodeString() {
	p, _ := jsonx.DecodeString[person](`{"name":"Alice","age":30}`)

	fmt.Println(p.Name, p.Age)
	// Output: Alice 30
}

// Options

func ExampleWithStrictDecoding() {
	_, err := jsonx.DecodeString[person](
		`{"name":"Alice","age":30,"role":"admin"}`,
		jsonx.WithStrictDecoding(),
	)

	fmt.Println(err != nil)
	// Output: true
}

func ExampleWithUseNumber() {
	type record struct {
		Value any `json:"value"`
	}

	r, _ := jsonx.DecodeString[record](`{"value":12345}`, jsonx.WithUseNumber())

	num, ok := r.Value.(json.Number)

	fmt.Println(ok, num)
	// Output: true 12345
}

func ExampleEncodeToFile() {
	f, _ := os.CreateTemp("", "snk-*.json")

	defer func() { _ = os.Remove(f.Name()) }()

	_ = f.Close()

	_ = jsonx.EncodeToFile(f.Name(), person{Name: "Alice", Age: 30})

	p, _ := jsonx.DecodeFromFile[person](f.Name())

	fmt.Println(p.Name, p.Age)
	// Output: Alice 30
}

func ExampleDecodeFromFile() {
	f, _ := os.CreateTemp("", "snk-*.json")

	defer func() { _ = os.Remove(f.Name()) }()

	_ = os.WriteFile(f.Name(), []byte(`{"name":"Alice","age":30}`), 0o600)

	p, _ := jsonx.DecodeFromFile[person](f.Name())

	fmt.Println(p.Name, p.Age)
	// Output: Alice 30
}
