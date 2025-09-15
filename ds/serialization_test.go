package ds_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/ds"
	"github.com/stretchr/testify/assert"
)

func TestSet_ToJSON(t *testing.T) {
	set := ds.NewSet[int]()

	result, err := set.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, "[]", string(result))

	set = ds.NewSet(1)
	result, err = set.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, "[1]", string(result))

	stringSet := ds.NewSet("one", "two", "three")
	strResult, err := stringSet.ToJSON()
	assert.NoError(t, err)
	for _, item := range stringSet.Values() {
		assert.Contains(t, string(strResult), item)
	}

}
