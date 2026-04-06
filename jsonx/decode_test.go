package jsonx_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/SharkByteSoftware/go-snk/jsonx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type typedNumber struct {
	Value int `json:"value"`
}

type anyNumber struct {
	Value any `json:"value"`
}

type namedFields struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const (
	validJSON        = `{"name":"Alice","age":30}`
	unknownFieldJSON = `{"name":"Alice","age":30,"unknown":"field"}`
	numberJSON       = `{"value":12345}`
	invalidJSON      = `{invalid}`
)

func TestDecode(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		result, err := jsonx.Decode[namedFields](strings.NewReader(validJSON))
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Alice", result.Name)
		assert.Equal(t, 30, result.Age)
	})

	t.Run("invalid json", func(t *testing.T) {
		result, err := jsonx.Decode[namedFields](strings.NewReader(invalidJSON))
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("unknown fields ignored by default", func(t *testing.T) {
		result, err := jsonx.Decode[namedFields](strings.NewReader(unknownFieldJSON))
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Alice", result.Name)
	})
}

func TestDecode_WithStrictDecoding(t *testing.T) {
	t.Run("unknown field returns error", func(t *testing.T) {
		result, err := jsonx.Decode[namedFields](strings.NewReader(unknownFieldJSON), jsonx.WithStrictDecoding())
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("known fields only succeeds", func(t *testing.T) {
		result, err := jsonx.Decode[namedFields](strings.NewReader(validJSON), jsonx.WithStrictDecoding())
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Alice", result.Name)
	})
}

func TestDecode_WithUseNumber(t *testing.T) {
	t.Run("any field is json.Number not float64", func(t *testing.T) {
		result, err := jsonx.Decode[anyNumber](strings.NewReader(numberJSON), jsonx.WithUseNumber())
		require.NoError(t, err)
		require.NotNil(t, result)

		num, ok := result.Value.(json.Number)
		require.True(t, ok, "expected json.Number, got %T", result.Value)
		assert.Equal(t, "12345", num.String())
	})

	t.Run("without UseNumber any field is float64", func(t *testing.T) {
		result, err := jsonx.Decode[anyNumber](strings.NewReader(numberJSON))
		require.NoError(t, err)
		require.NotNil(t, result)

		_, ok := result.Value.(float64)
		require.True(t, ok, "expected float64, got %T", result.Value)
	})

	t.Run("typed field unaffected by UseNumber", func(t *testing.T) {
		result, err := jsonx.Decode[typedNumber](strings.NewReader(numberJSON), jsonx.WithUseNumber())
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 12345, result.Value)
	})
}

func TestDecodeBytes(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		result, err := jsonx.DecodeBytes[namedFields]([]byte(validJSON))
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Alice", result.Name)
		assert.Equal(t, 30, result.Age)
	})

	t.Run("invalid json", func(t *testing.T) {
		result, err := jsonx.DecodeBytes[namedFields]([]byte(invalidJSON))
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("options forwarded", func(t *testing.T) {
		result, err := jsonx.DecodeBytes[anyNumber]([]byte(numberJSON), jsonx.WithUseNumber())
		require.NoError(t, err)
		require.NotNil(t, result)

		_, ok := result.Value.(json.Number)
		require.True(t, ok, "expected json.Number, got %T", result.Value)
	})
}

func TestDecodeString(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		result, err := jsonx.DecodeString[namedFields](validJSON)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Alice", result.Name)
		assert.Equal(t, 30, result.Age)
	})

	t.Run("invalid json", func(t *testing.T) {
		result, err := jsonx.DecodeString[namedFields](invalidJSON)
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("options forwarded", func(t *testing.T) {
		result, err := jsonx.DecodeString[anyNumber](numberJSON, jsonx.WithUseNumber())
		require.NoError(t, err)
		require.NotNil(t, result)

		_, ok := result.Value.(json.Number)
		require.True(t, ok, "expected json.Number, got %T", result.Value)
	})
}

func TestDecodeFromFile(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		result, err := jsonx.DecodeFromFile[namedFields]("../testdata/jsonx/happy_path.json")
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Alice", result.Name)
		assert.Equal(t, 30, result.Age)
	})

	t.Run("file does not exist", func(t *testing.T) {
		result, err := jsonx.DecodeFromFile[namedFields]("does_not_exist.json")
		assert.Nil(t, result)
		require.Error(t, err)
		require.ErrorContains(t, err, "open file: open does_not_exist.json: no such file or directory")
	})

	t.Run("invalid json", func(t *testing.T) {
		result, err := jsonx.DecodeFromFile[namedFields]("../testdata/jsonx/invalid_json.json")
		assert.Nil(t, result)
		require.Error(t, err)
	})
}
