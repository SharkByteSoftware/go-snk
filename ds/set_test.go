package ds_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	set := ds.NewSet[int]()
	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())

	set = ds.NewSet(1, 2, 3, 4, 5)
	assert.NotNil(t, set)
	assert.Equal(t, 5, set.Size())
}

func TestSet_Add(t *testing.T) {
	set := ds.NewSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())
}

func TestSet_Contains(t *testing.T) {
	set := ds.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(5))
}

func TestSet_Remove(t *testing.T) {
	set := ds.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())
	assert.True(t, set.Contains(5))

	set.Remove(5)
	assert.Equal(t, 4, set.Size())
	assert.False(t, set.Contains(5))
}

func TestSet_Size(t *testing.T) {
	set := ds.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())

	set = ds.NewSet[int]()
	assert.Equal(t, 0, set.Size())
}

func TestSet_Clear(t *testing.T) {
	set := ds.NewSet[int](1, 2, 3, 4, 5)
	assert.Equal(t, 5, set.Size())

	set.Clear()
	assert.Equal(t, 0, set.Size())
}

func TestSet_Values(t *testing.T) {
	set := ds.NewSet[int](1, 2, 3, 4, 5)
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

func TestSet_Intersect(t *testing.T) {
	set1 := ds.NewSet[int](1, 2, 3, 4, 5)
	set2 := ds.NewSet[int](2, 3, 4, 5, 6)

	result := set1.Intersect(set1)
	assert.Equal(t, set1.Size(), result.Size())
	assert.Equal(t, set1, result)

	result = set1.Intersect(set2)
	assert.Equal(t, 4, result.Size())
	assert.NotContains(t, result.Values(), 1)
	assert.NotContains(t, result.Values(), 6)

	result = set2.Intersect(set1)
	assert.Equal(t, 4, result.Size())
	assert.NotContains(t, result.Values(), 1)
	assert.NotContains(t, result.Values(), 6)
}

func TestSet_Union(t *testing.T) {
	set1 := ds.NewSet[int](1, 2, 3, 4, 5)
	set2 := ds.NewSet[int](4, 5, 6, 7, 256)

	result := set1.Union(set1)
	assert.Equal(t, 5, result.Size())
	assert.Equal(t, set1, result)

	result = set1.Union(set2)
	assert.Equal(t, 8, result.Size())
	assert.Subset(t, result.Values(), set1.Values())
	assert.Subset(t, result.Values(), set2.Values())

	result = set2.Union(set1)
	assert.Equal(t, 8, result.Size())
	assert.Subset(t, result.Values(), set1.Values())
	assert.Subset(t, result.Values(), set2.Values())
}

func TestSet_Difference(t *testing.T) {
	set1 := ds.NewSet[int](1, 2, 3, 4, 5)
	set2 := ds.NewSet[int](4, 5, 6, 7, 256)
	nullSet := ds.NewSet[int]()

	result := set1.Difference(set1)
	assert.Equal(t, 0, result.Size())

	result = set1.Difference(set2)
	assert.Equal(t, 3, result.Size())
	assert.Subset(t, result.Values(), []int{1, 2, 3})

	result = set2.Difference(set1)
	assert.Equal(t, 3, result.Size())
	assert.Subset(t, result.Values(), []int{6, 7, 256})

	result = nullSet.Difference(set1)
	assert.Equal(t, 0, result.Size())
	assert.Equal(t, nullSet, result)
}
