package queues

import (
	"github.com/stretchr/testify/assert"
	"sort"
	_ "sort"
	"testing"
)

var s sort.Interface

type Package struct {
	Weight int
}

func TestIntPriorityQueueAscending(t *testing.T) {
	intQueue := NewPriorityQueueWithDefault[int](func(prev int, curr int) int {
		if prev < curr {
			return -1
		}
		if prev > curr {
			return 1
		}

		return 0
	})

	intQueue.Enqueue(7)
	intQueue.Enqueue(3)
	intQueue.Enqueue(5)

	assert.Equal(t, 3, intQueue.Len()) // 3,5,7

	peek, ok := intQueue.Peek()
	assert.Equal(t, 3, peek) // 3

	top, ok := intQueue.Dequeue()
	assert.Equal(t, 3, top) // 3
	assert.True(t, ok)
	assert.Equal(t, 2, intQueue.Len()) //5,7

	intQueue.Enqueue(2) // 2,5,7
	peekAgain, okAgain := intQueue.Peek()
	assert.Equal(t, 2, peekAgain) // 2
	assert.True(t, okAgain)
	topAgain, okAgainAgain := intQueue.Dequeue() // 2
	assert.Equal(t, 2, topAgain)
	assert.True(t, okAgainAgain)
	assert.Equal(t, 2, intQueue.Len()) //5,7

	intQueue.Enqueue(6)                                    // 5, 6, 7
	topAgainAgain, okAgainAgainAgain := intQueue.Dequeue() // 5
	assert.Equal(t, 5, topAgainAgain)
	assert.True(t, okAgainAgainAgain)
	assert.Equal(t, 2, intQueue.Len()) // 6,7
}

func TestIntPriorityQueueDescending(t *testing.T) {
	intQueue := NewPriorityQueueWithDefault[int](func(prev int, curr int) int {
		if prev > curr {
			return -1
		}
		if prev < curr {
			return 1
		}

		return 0
	})

	intQueue.Enqueue(7)
	intQueue.Enqueue(3)
	intQueue.Enqueue(5)

	assert.Equal(t, 3, intQueue.Len()) // 7,5,3

	peek, ok := intQueue.Peek()
	assert.Equal(t, 7, peek) // 7

	top, ok := intQueue.Dequeue()
	assert.Equal(t, 7, top) // 7
	assert.True(t, ok)
	assert.Equal(t, 2, intQueue.Len()) //5,3

	intQueue.Enqueue(2) // 5,3,2
	peekAgain, okAgain := intQueue.Peek()
	assert.Equal(t, 5, peekAgain) // 5
	assert.True(t, okAgain)
	topAgain, okAgainAgain := intQueue.Dequeue() // 5
	assert.Equal(t, 5, topAgain)
	assert.True(t, okAgainAgain)
	assert.Equal(t, 2, intQueue.Len()) //3,2

	intQueue.Enqueue(6)                                    // 6,3,2
	topAgainAgain, okAgainAgainAgain := intQueue.Dequeue() // 6
	assert.Equal(t, 6, topAgainAgain)
	assert.True(t, okAgainAgainAgain)
	assert.Equal(t, 2, intQueue.Len()) // 3,2
}

func TestPackagePriorityQueue(t *testing.T) {
	packageQueue := NewPriorityQueueWithDefault[Package](func(prev Package, curr Package) int {
		if prev.Weight < curr.Weight {
			return -1
		}
		if prev.Weight > curr.Weight {
			return 1
		}

		return 0
	})

	packageQueue.Enqueue(Package{Weight: 7})
	packageQueue.Enqueue(Package{Weight: 3})
	packageQueue.Enqueue(Package{Weight: 5})

	assert.Equal(t, 3, packageQueue.Len()) // 3,5,7

	peek, ok := packageQueue.Peek()
	assert.Equal(t, 3, peek.Weight) // 3

	top, ok := packageQueue.Dequeue()
	assert.Equal(t, 3, top.Weight) // 3
	assert.True(t, ok)
	assert.Equal(t, 2, packageQueue.Len()) //5,7

	packageQueue.Enqueue(Package{Weight: 2}) // 2,5,7
	peekAgain, okAgain := packageQueue.Peek()
	assert.Equal(t, 2, peekAgain.Weight) // 2
	assert.True(t, okAgain)
	topAgain, okAgainAgain := packageQueue.Dequeue() // 2
	assert.Equal(t, 2, topAgain.Weight)
	assert.True(t, okAgainAgain)
	assert.Equal(t, 2, packageQueue.Len()) //5,7

	packageQueue.Enqueue(Package{Weight: 6})                   // 5, 6, 7
	topAgainAgain, okAgainAgainAgain := packageQueue.Dequeue() // 5
	assert.Equal(t, 5, topAgainAgain.Weight)
	assert.True(t, okAgainAgainAgain)
	assert.Equal(t, 2, packageQueue.Len()) // 6,7

	v := packageQueue.Values()
	assert.Equal(t, 2, len(v))
	v = append(v, Package{Weight: 1})

}
