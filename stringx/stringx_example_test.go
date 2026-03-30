package stringx_test

import "fmt"
import "github.com/SharkByteSoftware/go-snk/stringx"

func ExampleIsBlank() {
	fmt.Println(stringx.IsBlank(""))
	fmt.Println(stringx.IsBlank("   "))
	fmt.Println(stringx.IsBlank("hello"))
	// Output:
	// true
	// true
	// false
}

func ExampleCoalesce() {
	fmt.Println(stringx.Coalesce("", "", "fallback"))
	fmt.Println(stringx.Coalesce("first", "second"))
	// Output:
	// fallback
	// first
}

func ExampleTruncate() {
	fmt.Println(stringx.Truncate("hello world", 5))
	fmt.Println(stringx.Truncate("hi", 10))
	// Output:
	// hello
	// hi
}

func ExampleWrap() {
	fmt.Println(stringx.Wrap("hello", "(", ")"))
	fmt.Println(stringx.Wrap("world", "<b>", "</b>"))
	// Output:
	// (hello)
	// <b>world</b>
}
