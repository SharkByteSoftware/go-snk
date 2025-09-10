package sets_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-sink/sets"
	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	set := sets.NewSet[int]()
	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())

	set = sets.NewSet(1, 2, 3, 4, 5)
	assert.NotNil(t, set)
	assert.Equal(t, 5, set.Size())
}

func TestSet_Add(t *testing.T) {
	set := sets.NewSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())
}

func TestSet_Contains(t *testing.T) {
	set := sets.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(5))
}

func TestSet_Remove(t *testing.T) {
	set := sets.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())
	assert.True(t, set.Contains(5))

	set.Remove(5)
	assert.Equal(t, 4, set.Size())
	assert.False(t, set.Contains(5))
}

func TestSet_Size(t *testing.T) {
	set := sets.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())

	set = sets.NewSet[int]()
	assert.Equal(t, 0, set.Size())
}

func TestSet_Clear(t *testing.T) {
	set := sets.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())

	set.Clear()
	assert.Equal(t, 0, set.Size())
}

func TestSet_Values(t *testing.T) {
	set := sets.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())

	values := set.Values()
	assert.Equal(t, 5, len(values))
	for _, v := range values {
		assert.True(t, set.Contains(v))
	}

	set.Clear()
	assert.Equal(t, 0, set.Size())
	assert.Equal(t, []int{}, set.Values())
}
