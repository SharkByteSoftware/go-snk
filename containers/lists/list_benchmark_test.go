package lists_test

import (
	"fmt"
	"testing"

	"github.com/SharkByteSoftware/go-snk/containers/lists"
)

const (
	sliceSize  = 1000
	sliceCount = 5
	maxRandInt = 1000
)

var startingSize = []int{0, 1, 10, 100, 1000, 10000}

func BenchmarkList_PushFront(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("len: %d", size), func(b *testing.B) {
			for b.Loop() {
				l := lists.New[int]()
				for v := range size {
					l.PushFront(v)
				}
			}
		})
	}
}

func BenchmarkList_PushBack(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("len: %d", size), func(b *testing.B) {
			for b.Loop() {
				l := lists.New[int]()
				for v := range size {
					l.PushBack(v)
				}
			}
		})
	}
}

func BenchmarkList_Remove(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("len: %d", size), func(b *testing.B) {
			l := lists.New[int]()
			for v := range size {
				l.PushBack(v)
			}
			for b.Loop() {
				if !l.IsEmpty() {
					l.Remove(l.Front())
				}
			}
		})
	}
}

func BenchmarkList_MoveToFront(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("len: %d", size), func(b *testing.B) {
			l := lists.New[int]()
			for v := range size {
				l.PushBack(v)
			}
			for b.Loop() {
				if !l.IsEmpty() {
					l.MoveToFront(l.Back())
				}
			}
		})
	}
}

func BenchmarkList_MoveToBack(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("len: %d", size), func(b *testing.B) {
			l := lists.New[int]()
			for v := range size {
				l.PushBack(v)
			}
			for b.Loop() {
				if !l.IsEmpty() {
					l.MoveToBack(l.Front())
				}
			}
		})
	}
}

func BenchmarkList_PushFrontList(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("len: %d", size), func(b *testing.B) {
			other := lists.New[int]()
			for v := range size {
				other.PushBack(v)
			}
			for b.Loop() {
				l := lists.New[int]()
				l.PushFrontList(other)
			}
		})
	}
}

func BenchmarkList_PushBackList(b *testing.B) {
	for _, size := range startingSize {
		b.Run(fmt.Sprintf("len: %d", size), func(b *testing.B) {
			other := lists.New[int]()
			for v := range size {
				other.PushBack(v)
			}
			for b.Loop() {
				l := lists.New[int]()
				l.PushBackList(other)
			}
		})
	}
}
