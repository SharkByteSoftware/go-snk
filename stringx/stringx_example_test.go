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

func ExampleCoalesceFunc() {
	notBlank := func(s string) bool { return !stringx.IsBlank(s) }

	// Unlike Coalesce, CoalesceFunc skips whitespace-only strings
	fmt.Println(stringx.CoalesceFunc(notBlank, "   ", "\t", "fallback"))
	fmt.Println(stringx.CoalesceFunc(notBlank, "hello", "world"))
	// Output:
	// fallback
	// hello
}

func ExampleCoalesceFunc_customPredicate() {
	// Only accept strings longer than 3 characters
	longEnough := func(s string) bool { return len(s) > 3 }

	fmt.Println(stringx.CoalesceFunc(longEnough, "hi", "hey", "hello"))
	fmt.Println(stringx.CoalesceFunc(longEnough, "hi", "hey"))
	// Output:
	// hello
	//
}
