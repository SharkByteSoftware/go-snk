package queues_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/containers/queues"
	"github.com/stretchr/testify/assert"
)

func TestDQueue_NewQueue(t *testing.T) {
	q := queues.NewQueue[int]()
	assert.NotNil(t, q)
	assert.True(t, q.IsEmpty())
	assert.Equal(t, 0, q.Size())

	q = queues.NewQueue[int](1, 2, 3)
	assert.NotNil(t, q)
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 3, q.Size())
}

func TestDQueue_Enqueue(t *testing.T) {
	q := queues.NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 3, q.Size())
	assert.False(t, q.IsEmpty())

	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestDQueue_EnqueueFront(t *testing.T) {
	q := queues.NewQueue[int]()

	q.EnqueueFront(1)
	q.EnqueueFront(2)
	q.EnqueueFront(3)

	assert.Equal(t, 3, q.Size())
	assert.False(t, q.IsEmpty())

	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
}

func TestDQueue_Dequeue(t *testing.T) {
	q := queues.NewQueue(1, 2, 4)

	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 1, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 2, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 0, q.Size())
	assert.Equal(t, 4, v)

	v, ok = q.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, 0, q.Size())
}

func TestDQueue_DequeueBack(t *testing.T) {
	q := queues.NewQueue(1, 2, 4)

	v, ok := q.DequeueBack()
	assert.True(t, ok)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 4, v)

	v, ok = q.DequeueBack()
	assert.True(t, ok)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 2, v)

	v, ok = q.DequeueBack()
	assert.True(t, ok)
	assert.Equal(t, 0, q.Size())
	assert.Equal(t, 1, v)

	v, ok = q.DequeueBack()
	assert.False(t, ok)
	assert.Equal(t, 0, q.Size())
}

func TestDQueue_Peek(t *testing.T) {
	q := queues.NewQueue[int]()

	v, ok := q.Peek()
	assert.False(t, ok)
	assert.Equal(t, v, 0)

	q.Enqueue(1)
	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 1, v)

	q.Enqueue(2)
	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 1, v)
}

func TestDQueue_PeekBack(t *testing.T) {
	q := queues.NewQueue[int]()

	v, ok := q.PeekBack()
	assert.False(t, ok)
	assert.Equal(t, v, 0)

	q.Enqueue(1)
	v, ok = q.PeekBack()
	assert.True(t, ok)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 1, v)

	q.Enqueue(2)
	v, ok = q.PeekBack()
	assert.True(t, ok)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 2, v)
}

func TestDQueue_Clear(t *testing.T) {
	q := queues.NewQueue[int]()
	assert.Equal(t, 0, q.Size())

	q.Clear()
	assert.Equal(t, 0, q.Size())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 3, q.Size())

	q.Clear()
	assert.Equal(t, 0, q.Size())
}

func TestDQueue_Values(t *testing.T) {
	q := queues.NewQueue[int]()

	values := q.Values()
	assert.Len(t, values, 0)

	q = queues.NewQueue(1, 2, 5)
	values = q.Values()
	assert.Len(t, values, 3)
	assert.Equal(t, []int{1, 2, 5}, values)
}
