package parallel_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/SharkByteSoftware/go-snk/slicex/parallel"
	"github.com/stretchr/testify/assert"
)

func TestMapWithLimit_PanicPropagates(t *testing.T) {
	assert.PanicsWithValue(t, "boom", func() {
		_ = parallel.MapWithLimit([]int{1, 2, 3, 4}, func(n int) int {
			if n == 3 {
				panic("boom")
			}

			return n
		}, 2)
	})
}

func TestApplyWithLimit_PanicPropagates(t *testing.T) {
	assert.PanicsWithValue(t, "boom", func() {
		parallel.ApplyWithLimit([]int{1, 2, 3, 4}, func(n int) {
			if n == 2 {
				panic("boom")
			}
		}, 2)
	})
}

func TestApplyWithLimit_EmptySlice(t *testing.T) {
	assert.NotPanics(t, func() {
		parallel.ApplyWithLimit([]int{}, func(int) {}, 4)
		parallel.ApplyWithLimit([]int(nil), func(int) {}, 0)
	})
}

func TestApplyWithLimit_BoundsConcurrency(t *testing.T) {
	const limit = 4

	var current, peak atomic.Int32

	parallel.ApplyWithLimit(make([]int, 100), func(int) {
		trackPeak(&current, &peak)
	}, limit)

	assert.LessOrEqual(t, peak.Load(), int32(limit))
}

func TestApply_BoundsConcurrency(t *testing.T) {
	const defaultLimit = 1000 // mirrors the package's defaultMaxConcurrency

	var current, peak atomic.Int32

	parallel.Apply(make([]int, 3*defaultLimit), func(int) {
		trackPeak(&current, &peak)
	})

	assert.LessOrEqual(t, peak.Load(), int32(defaultLimit))
}

func TestPartitionWithLimit(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	even, odd := parallel.PartitionWithLimit([]int{1, 2, 3, 4, 5, 6}, isEven, 3)
	assert.Equal(t, []int{2, 4, 6}, even)
	assert.Equal(t, []int{1, 3, 5}, odd)

	even, odd = parallel.PartitionWithLimit([]int{}, isEven, 3)
	assert.Empty(t, even)
	assert.Empty(t, odd)
}

func trackPeak(current, peak *atomic.Int32) {
	n := current.Add(1)

	for {
		p := peak.Load()
		if n <= p || peak.CompareAndSwap(p, n) {
			break
		}
	}

	time.Sleep(time.Millisecond)
	current.Add(-1)
}
