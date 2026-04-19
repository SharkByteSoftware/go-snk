package lists_test

import (
	"fmt"

	"github.com/SharkByteSoftware/go-snk/containers/lists"
)

func ExampleNew() {
	l := lists.New(1, 2, 3)

	fmt.Println(l.Values())
	// Output: [1 2 3]
}

func ExampleList_Len() {
	l := lists.New(1, 2, 3)

	fmt.Println(l.Len())
	// Output: 3
}

func ExampleList_IsEmpty() {
	l := lists.New[int]()

	fmt.Println(l.IsEmpty())
	l.PushBack(1)
	fmt.Println(l.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleList_Size() {
	l := lists.New(10, 20, 30)

	fmt.Println(l.Size())
	// Output: 3
}

func ExampleList_Front() {
	l := lists.New(1, 2, 3)

	fmt.Println(l.Front().Value)
	// Output: 1
}

func ExampleList_Back() {
	l := lists.New(1, 2, 3)

	fmt.Println(l.Back().Value)
	// Output: 3
}

func ExampleList_PushBack() {
	l := lists.New(1, 2)

	l.PushBack(3)

	fmt.Println(l.Values())
	// Output: [1 2 3]
}

func ExampleList_PushFront() {
	l := lists.New(2, 3)

	l.PushFront(1)

	fmt.Println(l.Values())
	// Output: [1 2 3]
}

func ExampleList_Append() {
	l := lists.New(1, 2)

	l.Append(3, 4, 5)

	fmt.Println(l.Values())
	// Output: [1 2 3 4 5]
}

func ExampleList_Prepend() {
	l := lists.New(3, 4)

	l.Prepend(1, 2)

	fmt.Println(l.Values())
	// Output: [1 2 3 4]
}

func ExampleList_Remove() {
	l := lists.New(1, 2, 3)

	front := l.Front()
	val := l.Remove(front)

	fmt.Println(val)
	fmt.Println(l.Values())
	// Output:
	// 1
	// [2 3]
}

func ExampleList_InsertBefore() {
	l := lists.New(1, 3)

	mark := l.Back()
	l.InsertBefore(2, mark)

	fmt.Println(l.Values())
	// Output: [1 2 3]
}

func ExampleList_InsertAfter() {
	l := lists.New(1, 3)

	mark := l.Front()
	l.InsertAfter(2, mark)

	fmt.Println(l.Values())
	// Output: [1 2 3]
}

func ExampleList_MoveToFront() {
	l := lists.New(1, 2, 3)

	l.MoveToFront(l.Back())

	fmt.Println(l.Values())
	// Output: [3 1 2]
}

func ExampleList_MoveToBack() {
	l := lists.New(1, 2, 3)

	l.MoveToBack(l.Front())

	fmt.Println(l.Values())
	// Output: [2 3 1]
}

func ExampleList_MoveBefore() {
	l := lists.New(1, 3, 2)

	l.MoveBefore(l.Back(), l.Back().Prev())

	fmt.Println(l.Values())
	// Output: [1 2 3]
}

func ExampleList_MoveAfter() {
	l := lists.New(2, 1, 3)

	l.MoveAfter(l.Back().Prev(), l.Front())

	fmt.Println(l.Values())
	// Output: [2 1 3]
}

func ExampleList_PushBackList() {
	l1 := lists.New(1, 2)
	l2 := lists.New(3, 4)

	l1.PushBackList(l2)

	fmt.Println(l1.Values())
	// Output: [1 2 3 4]
}

func ExampleList_PushFrontList() {
	l1 := lists.New(3, 4)
	l2 := lists.New(1, 2)

	l1.PushFrontList(l2)

	fmt.Println(l1.Values())
	// Output: [1 2 3 4]
}

func ExampleList_Values() {
	l := lists.New(10, 20, 30)

	fmt.Println(l.Values())
	// Output: [10 20 30]
}

func ExampleList_Clear() {
	l := lists.New(1, 2, 3)

	l.Clear()

	fmt.Println(l.IsEmpty())
	// Output: true
}

func ExampleList_ForEach() {
	l := lists.New(1, 2, 3, 4, 5)

	var sum int
	l.ForEach(func(v int) { sum += v })

	fmt.Println(sum)
	// Output: 15
}

func ExampleList_Init() {
	l := lists.New(1, 2, 3)

	l.Init()

	fmt.Println(l.IsEmpty())
	// Output: true
}

func ExampleElement_Next() {
	l := lists.New(1, 2, 3)

	fmt.Println(l.Front().Next().Value)
	// Output: 2
}

func ExampleElement_Prev() {
	l := lists.New(1, 2, 3)

	fmt.Println(l.Back().Prev().Value)
	// Output: 2
}
