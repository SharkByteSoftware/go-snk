package queues_test

import (
	"cmp"
	"fmt"

	"github.com/SharkByteSoftware/go-snk/containers/queues"
)

// Queue examples

func ExampleNewQueue() {
	q := queues.NewQueue(1, 2, 3)

	fmt.Println(q.Values())
	// Output: [1 2 3]
}

func ExampleQueue_Enqueue() {
	q := queues.NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Values())
	// Output: [1 2 3]
}

func ExampleQueue_EnqueueFront() {
	q := queues.NewQueue(2, 3)

	q.EnqueueFront(1)

	fmt.Println(q.Values())
	// Output: [1 2 3]
}

func ExampleQueue_Dequeue() {
	q := queues.NewQueue(1, 2, 3)

	val, ok := q.Dequeue()

	fmt.Println(val, ok)
	fmt.Println(q.Values())
	// Output:
	// 1 true
	// [2 3]
}

func ExampleQueue_DequeueBack() {
	q := queues.NewQueue(1, 2, 3)

	val, ok := q.DequeueBack()

	fmt.Println(val, ok)
	fmt.Println(q.Values())
	// Output:
	// 3 true
	// [1 2]
}

func ExampleQueue_Peek() {
	q := queues.NewQueue(1, 2, 3)

	val, ok := q.Peek()

	fmt.Println(val, ok)
	fmt.Println(q.Size()) // still 3 after peek
	// Output:
	// 1 true
	// 3
}

func ExampleQueue_PeekBack() {
	q := queues.NewQueue(1, 2, 3)

	val, ok := q.PeekBack()

	fmt.Println(val, ok)
	fmt.Println(q.Size()) // still 3 after peek
	// Output:
	// 3 true
	// 3
}

func ExampleQueue_IsEmpty() {
	q := queues.NewQueue[int]()

	fmt.Println(q.IsEmpty())
	q.Enqueue(1)
	fmt.Println(q.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleQueue_Size() {
	q := queues.NewQueue(10, 20, 30)

	fmt.Println(q.Size())
	// Output: 3
}

func ExampleQueue_Clear() {
	q := queues.NewQueue(1, 2, 3)

	q.Clear()

	fmt.Println(q.IsEmpty())
	// Output: true
}

func ExampleQueue_Values() {
	q := queues.NewQueue(10, 20, 30)

	fmt.Println(q.Values())
	// Output: [10 20 30]
}

// PriorityQueue examples

func ExampleNewPriorityQueue() {
	items := []int{5, 1, 3, 2, 4}
	pq := queues.NewPriorityQueue(items, cmp.Compare[int])

	val, _ := pq.Dequeue()

	fmt.Println(val)
	// Output: 1
}

func ExampleNewPriorityQueueWithDefault() {
	pq := queues.NewPriorityQueueWithDefault(cmp.Compare[int])

	pq.Enqueue(3)
	pq.Enqueue(1)
	pq.Enqueue(2)

	val, _ := pq.Dequeue()

	fmt.Println(val)
	// Output: 1
}

func ExamplePriorityQueue_Enqueue() {
	pq := queues.NewPriorityQueueWithDefault(cmp.Compare[int])

	pq.Enqueue(5)
	pq.Enqueue(1)
	pq.Enqueue(3)

	val, _ := pq.Peek()

	fmt.Println(val)
	// Output: 1
}

func ExamplePriorityQueue_Dequeue() {
	pq := queues.NewPriorityQueue([]int{3, 1, 2}, cmp.Compare[int])

	for !pq.IsEmpty() {
		val, _ := pq.Dequeue()
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExamplePriorityQueue_Peek() {
	pq := queues.NewPriorityQueue([]int{3, 1, 2}, cmp.Compare[int])

	val, ok := pq.Peek()

	fmt.Println(val, ok)
	fmt.Println(pq.Size()) // still 3 after peek
	// Output:
	// 1 true
	// 3
}

func ExamplePriorityQueue_IsEmpty() {
	pq := queues.NewPriorityQueueWithDefault(cmp.Compare[int])

	fmt.Println(pq.IsEmpty())
	pq.Enqueue(1)
	fmt.Println(pq.IsEmpty())
	// Output:
	// true
	// false
}

func ExamplePriorityQueue_Size() {
	pq := queues.NewPriorityQueue([]int{1, 2, 3}, cmp.Compare[int])

	fmt.Println(pq.Size())
	// Output: 3
}

func ExamplePriorityQueue_Clear() {
	pq := queues.NewPriorityQueue([]int{1, 2, 3}, cmp.Compare[int])

	pq.Clear()

	fmt.Println(pq.IsEmpty())
	// Output: true
}
