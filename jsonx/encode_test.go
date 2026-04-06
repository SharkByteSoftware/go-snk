package jsonx_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/SharkByteSoftware/go-snk/jsonx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// encode_test.go

type encodeFields struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const (
	encodedJSON       = `{"name":"Alice","age":30}` + "\n"
	encodedJSONPretty = "{\n  \"name\": \"Alice\",\n  \"age\": 30\n}\n"
	htmlString        = `{"url":"https://example.com/search?a=1&b=2"}` + "\n"
	htmlStringEscaped = `{"url":"https://example.com/search?a=1\u0026b=2"}` + "\n"
)

type urlField struct {
	URL string `json:"url"`
}

func TestEncode(t *testing.T) {
	value := encodeFields{Name: "Alice", Age: 30}

	t.Run("happy path", func(t *testing.T) {
		var buf strings.Builder

		err := jsonx.Encode(&buf, value)
		require.NoError(t, err)
		assert.JSONEq(t, encodedJSON, buf.String())
	})

	t.Run("unencodable value returns error", func(t *testing.T) {
		var buf strings.Builder

		err := jsonx.Encode(&buf, complex(1, 2))
		require.Error(t, err)
	})

	t.Run("html not escaped by default", func(t *testing.T) {
		var buf strings.Builder

		err := jsonx.Encode(&buf, urlField{URL: "https://example.com/search?a=1&b=2"})
		require.NoError(t, err)
		assert.JSONEq(t, htmlString, buf.String())
	})

	t.Run("html escaped with WithEscapeHTML", func(t *testing.T) {
		var buf strings.Builder

		err := jsonx.Encode(&buf, urlField{URL: "https://example.com/search?a=1&b=2"}, jsonx.WithEscapeHTML())
		require.NoError(t, err)
		assert.JSONEq(t, htmlStringEscaped, buf.String())
	})

	t.Run("pretty printed with WithIndent", func(t *testing.T) {
		var buf strings.Builder

		err := jsonx.Encode(&buf, value, jsonx.WithIndent("  "))
		require.NoError(t, err)
		assert.JSONEq(t, encodedJSONPretty, buf.String())
	})
}

func TestEncodeBytes(t *testing.T) {
	v := encodeFields{Name: "Alice", Age: 30}

	t.Run("happy path", func(t *testing.T) {
		result, err := jsonx.EncodeBytes(v)
		require.NoError(t, err)
		//nolint:testifylint
		assert.Equal(t, []byte(encodedJSON), result)
	})

	t.Run("unencodable value returns error", func(t *testing.T) {
		result, err := jsonx.EncodeBytes(complex(1, 2))
		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("options forwarded", func(t *testing.T) {
		result, err := jsonx.EncodeBytes(v, jsonx.WithIndent("  "))
		require.NoError(t, err)
		//nolint:testifylint
		assert.Equal(t, []byte(encodedJSONPretty), result)
	})
}

func TestEncodeString(t *testing.T) {
	v := encodeFields{Name: "Alice", Age: 30}

	t.Run("happy path", func(t *testing.T) {
		result, err := jsonx.EncodeString(v)
		require.NoError(t, err)
		assert.JSONEq(t, encodedJSON, result)
	})

	t.Run("unencodable value returns error", func(t *testing.T) {
		result, err := jsonx.EncodeString(complex(1, 2))
		require.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("options forwarded", func(t *testing.T) {
		result, err := jsonx.EncodeString(v, jsonx.WithIndent("\t"))
		require.NoError(t, err)
		assert.Contains(t, result, "\t")
	})
}

func TestEncodeToFile(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "output.json")
		value := namedFields{Name: "Alice", Age: 30}

		err := jsonx.EncodeToFile(path, value)
		require.NoError(t, err)

		// Round-trip: decode what was written and verify
		result, err := jsonx.DecodeFromFile[namedFields](path)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, "Alice", result.Name)
		assert.Equal(t, 30, result.Age)
	})

	t.Run("directory does not exist", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "nonexistent", "output.json")
		value := namedFields{Name: "Alice", Age: 30}

		err := jsonx.EncodeToFile(path, value)
		require.Error(t, err)
		require.ErrorContains(t, err, "create file:")
	})

	t.Run("path is a directory", func(t *testing.T) {
		path := t.TempDir()
		value := namedFields{Name: "Alice", Age: 30}

		err := jsonx.EncodeToFile(path, value)
		require.Error(t, err)
	})

	t.Run("empty struct", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "output.json")

		err := jsonx.EncodeToFile(path, namedFields{})
		require.NoError(t, err)

		result, err := jsonx.DecodeFromFile[namedFields](path)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Empty(t, result.Name)
		assert.Equal(t, 0, result.Age)
	})
}
