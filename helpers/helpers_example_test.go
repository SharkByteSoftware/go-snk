package helpers_test

import (
	"fmt"

	"github.com/SharkByteSoftware/go-snk/helpers"
)

// Pointer helpers

func ExampleNil() {
	p := helpers.Nil[int]()

	fmt.Println(p)
	// Output: <nil>
}

func ExampleIsNil() {
	n := helpers.Nil[string]()
	s := "hello"

	fmt.Println(helpers.IsNil(n))
	fmt.Println(helpers.IsNil(&s))
	// Output:
	// true
	// false
}

func ExampleAsPtr() {
	p := helpers.AsPtr(42)

	fmt.Println(*p)
	// Output: 42
}

func ExampleAsValue() {
	n := 7
	p := &n

	fmt.Println(helpers.AsValue(p))
	fmt.Println(helpers.AsValue[int](nil))
	// Output:
	// 7
	// 0
}

func ExampleAsValueOr() {
	n := 7
	p := &n

	fmt.Println(helpers.AsValueOr(p, -1))
	fmt.Println(helpers.AsValueOr[int](nil, -1))
	// Output:
	// 7
	// -1
}

// Value helpers

func ExampleEmpty() {
	fmt.Println(helpers.Empty[int]())
	fmt.Println(helpers.Empty[string]())
	fmt.Println(helpers.Empty[bool]())
	// Output:
	// 0
	//
	// false
}

func ExampleIsEmpty() {
	fmt.Println(helpers.IsEmpty(0))
	fmt.Println(helpers.IsEmpty(""))
	fmt.Println(helpers.IsEmpty(42))
	fmt.Println(helpers.IsEmpty("hello"))
	// Output:
	// true
	// true
	// false
	// false
}
