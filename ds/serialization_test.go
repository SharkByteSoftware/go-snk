package ds_test

import (
	"encoding/json"
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

func TestSet_FromJSON(t *testing.T) {
	set := ds.NewSet[int]()

	err := set.FromJSON([]byte("[]"))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(set.Values()))

	stringSet := ds.NewSet[string]()
	err = stringSet.FromJSON([]byte(`["one"]`))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(stringSet.Values()))
	assert.Equal(t, "one", stringSet.Values()[0])

	err = set.FromJSON([]byte("[5,6,256]"))
	assert.NoError(t, err)
	assert.Len(t, set.Values(), 3)
	assert.True(t, set.Contains(5))
	assert.True(t, set.Contains(6))
	assert.True(t, set.Contains(256))
}

func TestSet_MarshalJSON(t *testing.T) {
	jsonBytes, err := json.Marshal(ds.NewSet[int]())
	assert.NoError(t, err)
	assert.Equal(t, "[]", string(jsonBytes))

	jsonBytes, err = json.Marshal(ds.NewSet("one"))
	assert.NoError(t, err)
	assert.Equal(t, `["one"]`, string(jsonBytes))

	jsonBytes, err = json.Marshal(ds.NewSet(1, 2, 256))
	assert.NoError(t, err)
	assert.Len(t, string(jsonBytes), 9)

	var set ds.Set[int]
	err = json.Unmarshal(jsonBytes, &set)
	assert.NoError(t, err)
	assert.Len(t, set.Values(), 3)
}

func TestSet_UnmarshalJSON(t *testing.T) {
	var set = ds.NewSet[int]()

	err := json.Unmarshal([]byte(`[]`), &set)
	assert.NoError(t, err)
	assert.Len(t, set.Values(), 0)

	err = json.Unmarshal([]byte("[1]"), &set)
	assert.NoError(t, err)
	assert.Len(t, set.Values(), 1)
	assert.Equal(t, 1, set.Values()[0])

	err = json.Unmarshal([]byte("[1, 2, 3]"), &set)
	assert.NoError(t, err)
	assert.Len(t, set.Values(), 3)

	var stringSet ds.Set[string]
	err = json.Unmarshal([]byte(`["one"]`), &stringSet)
	assert.NoError(t, err)
	assert.Len(t, stringSet.Values(), 1)
	assert.Equal(t, "one", stringSet.Values()[0])
}
