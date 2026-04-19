package stacks_test

import (
	"fmt"

	"github.com/SharkByteSoftware/go-snk/containers/stacks"
)

func ExampleNew() {
	s := stacks.New(1, 2, 3)

	fmt.Println(s.Values())
	// Output: [1 2 3]
}

func ExampleStack_Push() {
	s := stacks.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	val, _ := s.Peek()

	fmt.Println(val)
	// Output: 3
}

func ExampleStack_Pop() {
	s := stacks.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	val, ok := s.Pop()

	fmt.Println(val, ok)
	fmt.Println(s.Size())
	// Output:
	// 3 true
	// 2
}

func ExampleStack_Peek() {
	s := stacks.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	val, ok := s.Peek()

	fmt.Println(val, ok)
	fmt.Println(s.Size()) // still 3 after peek
	// Output:
	// 3 true
	// 3
}

func ExampleStack_IsEmpty() {
	s := stacks.New[int]()

	fmt.Println(s.IsEmpty())
	s.Push(1)
	fmt.Println(s.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleStack_Size() {
	s := stacks.New(10, 20, 30)

	fmt.Println(s.Size())
	// Output: 3
}

func ExampleStack_Clear() {
	s := stacks.New(1, 2, 3)

	s.Clear()

	fmt.Println(s.IsEmpty())
	// Output: true
}

func ExampleStack_Values() {
	s := stacks.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Values())
	// Output: [3 2 1]
}
