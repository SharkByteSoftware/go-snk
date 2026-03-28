package sets_test

import (
	"encoding/json"
	"testing"

	"github.com/SharkByteSoftware/go-snk/containers/sets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSet_ToJSON(t *testing.T) {
	set := sets.New[int]()

	result, err := set.ToJSON()
	require.NoError(t, err)
	assert.Equal(t, "[]", string(result))

	set = sets.New(1)
	result, err = set.ToJSON()
	require.NoError(t, err)
	assert.Equal(t, "[1]", string(result))

	stringSet := sets.New("one", "two", "three")
	strResult, err := stringSet.ToJSON()
	require.NoError(t, err)

	for _, item := range stringSet.Values() {
		assert.Contains(t, string(strResult), item)
	}
}

func TestSet_ToJSONFailure(t *testing.T) {
	set := sets.New[complex128]()

	result, err := set.ToJSON()
	require.NoError(t, err)
	require.NotNil(t, result)

	set = sets.New(complex(1, 2))

	result, err = set.ToJSON()
	require.Error(t, err)
	require.Nil(t, result)
}

func TestSet_FromJSON(t *testing.T) {
	set := sets.New[int]()

	err := set.FromJSON([]byte("[]"))
	require.NoError(t, err)
	assert.Empty(t, set.Values())

	stringSet := sets.New[string]()
	err = stringSet.FromJSON([]byte(`["one"]`))
	require.NoError(t, err)
	assert.Len(t, stringSet.Values(), 1)
	assert.Equal(t, "one", stringSet.Values()[0])

	err = set.FromJSON([]byte("[5,6,256]"))
	require.NoError(t, err)
	assert.Len(t, set.Values(), 3)
	assert.True(t, set.Contains(5))
	assert.True(t, set.Contains(6))
	assert.True(t, set.Contains(256))
}

func TestSet_FromJSONFailure(t *testing.T) {
	set := sets.New[complex128]()

	err := set.FromJSON([]byte(`["one"]`))
	require.Error(t, err)
	assert.Empty(t, set.Values())
}

func TestSet_MarshalJSON(t *testing.T) {
	jsonBytes, err := json.Marshal(sets.New[int]())
	require.NoError(t, err)
	assert.JSONEq(t, "[]", string(jsonBytes))

	jsonBytes, err = json.Marshal(sets.New("one"))
	require.NoError(t, err)
	assert.JSONEq(t, `["one"]`, string(jsonBytes))

	jsonBytes, err = json.Marshal(sets.New(1, 2, 256))
	require.NoError(t, err)
	assert.Len(t, string(jsonBytes), 9)

	var set sets.Set[int]

	err = json.Unmarshal(jsonBytes, &set)
	require.NoError(t, err)
	assert.Len(t, set.Values(), 3)
}

func TestSet_UnmarshalJSON(t *testing.T) {
	var set = sets.New[int]()

	err := json.Unmarshal([]byte(`[]`), &set)
	require.NoError(t, err)
	assert.Empty(t, set.Values())

	err = json.Unmarshal([]byte("[1]"), &set)
	require.NoError(t, err)
	assert.Len(t, set.Values(), 1)
	assert.Equal(t, 1, set.Values()[0])

	err = json.Unmarshal([]byte("[1, 2, 3]"), &set)
	require.NoError(t, err)
	assert.Len(t, set.Values(), 3)

	var stringSet sets.Set[string]

	err = json.Unmarshal([]byte(`["one"]`), &stringSet)
	require.NoError(t, err)
	assert.Len(t, stringSet.Values(), 1)
	assert.Equal(t, "one", stringSet.Values()[0])
}
