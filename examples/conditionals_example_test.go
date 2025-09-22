package examples

import (
	"fmt"
	"github.com/SharkByteSoftware/go-snk/conditional"
)

func ExampleIf() {
	var data int

	data = 5
	resultTrue := conditional.If(data > 1, "true", "false")
	resultFalse := conditional.If(data < 2, "false", "true")

	fmt.Println(resultTrue)
	fmt.Println(resultFalse)
	// Output:
	// true
}

func ExampleIfNotNil() {
	f := func() {
		fmt.Println("function is called")
	}

	conditional.IfNotNil(new(int), f)

	// Output:
	// function is called
}

func ExampleIfCall() {
	ft := func() {
		fmt.Println("function is called when condition is true")
	}

	ff := func() {
		fmt.Println("function is called when condition is false")
	}

	var data int

	data = 5
	conditional.IfCall(data > 1, ft, ff)
	conditional.IfCall(data < 1, ft, ff)

	// Output:
	// function is called when condition is true
	// function is called when condition is false
}

func ExampleIfCallReturn() {
	addNum := func(a, b int) int {
		return a + b
	}

	mulNum := func(a, b int) int {
		return a * b
	}

	result1 := conditional.IfCallReturn(true, func() int {
		return addNum(2, 3)
	}, func() int {
		return mulNum(2, 3)
	})
	fmt.Println(result1)

	result2 := conditional.IfCallReturn(false, func() int {
		return addNum(2, 3)
	}, func() int {
		return mulNum(2, 3)
	})
	fmt.Println(result2)

	// Output:
	// 5
	// 6
}
