package conditional_test

import (
	"fmt"

	"github.com/SharkByteSoftware/go-snk/conditional"
)

func ExampleIf() {
	score := 85

	label := conditional.If(score >= 60, "pass", "fail")

	fmt.Println(label)
	// Output: pass
}

func ExampleIfNotNil() {
	name := "Alice"

	conditional.IfNotNil(&name, func() {
		fmt.Println("name is set")
	})
	// Output: name is set
}

func ExampleIfCall() {
	conditional.IfCall(
		true,
		func() { fmt.Println("welcome back") },
		func() { fmt.Println("please log in") },
	)
	// Output: welcome back
}

func ExampleIfCallReturn() {
	limit := conditional.IfCallReturn(
		false,
		func() int { return 100 },
		func() int { return 10 },
	)

	fmt.Println(limit)
	// Output: 10
}

func ExampleSwitch() {
	status := 2

	label := conditional.Switch(status, map[int]string{
		1: "active",
		2: "inactive",
		3: "pending",
	}, "unknown")

	fmt.Println(label)
	// Output: inactive
}

func ExampleSwitch_fallback() {
	status := 99

	label := conditional.Switch(status, map[int]string{
		1: "active",
		2: "inactive",
		3: "pending",
	}, "unknown")

	fmt.Println(label)
	// Output: unknown
}
