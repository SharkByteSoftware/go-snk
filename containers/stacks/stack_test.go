package stacks_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/containers/stacks"
	"github.com/stretchr/testify/assert"
)

func TestStack_New(t *testing.T) {
	st := stacks.New[int]()
	assert.NotNil(t, st)
	assert.Equal(t, 0, st.Size())

}

func TestStack_Push(t *testing.T) {
	st := stacks.New[int]()
	assert.Equal(t, 0, st.Size())

	st.Push(1)
	assert.Equal(t, 1, st.Size())

	st.Push(1)
	st.Push(2)
	st.Push(3)
	assert.Equal(t, 4, st.Size())

	assert.Equal(t, []int{3, 2, 1, 1}, st.Values())
}

func TestStack_Pop(t *testing.T) {
	st := stacks.New[int]()

	value, ok := st.Pop()
	assert.False(t, ok)

	st.Push(10)
	st.Push(20)

	value, ok = st.Pop()
	assert.True(t, ok)
	assert.Equal(t, 20, value)
	assert.Equal(t, 1, st.Size())

	value, ok = st.Pop()
	assert.True(t, ok)
	assert.Equal(t, 10, value)
	assert.Equal(t, 0, st.Size())

	value, ok = st.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, st.Size())
}

func TestStack_Peek(t *testing.T) {
	st := stacks.New[int]()

	value, ok := st.Peek()
	assert.False(t, ok)

	st.Push(10)
	st.Push(20)

	value, ok = st.Peek()
	assert.True(t, ok)
	assert.Equal(t, 20, value)
	assert.Equal(t, 2, st.Size())

	_, ok = st.Pop()
	assert.True(t, ok)

	value, ok = st.Peek()
	assert.True(t, ok)
	assert.Equal(t, 10, value)
	assert.Equal(t, 1, st.Size())
}

func TestStack_Values(t *testing.T) {
	st := stacks.New[int]()

	assert.Equal(t, []int{}, st.Values())

	st.Push(10)
	st.Push(20)

	assert.Equal(t, []int{20, 10}, st.Values())
}
