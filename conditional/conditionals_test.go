package conditional_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMock struct {
	mock.Mock
}

func (m *MyMock) MyFunc() {
	m.Called()
}

type MyMockR struct {
	mock.Mock
}

func (m *MyMockR) MyFunc() bool {
	args := m.Called()
	return args.Bool(0)
}

func TestIf(t *testing.T) {
	result := conditional.If(true, 1, 2)
	assert.Equal(t, 1, result)

	result = conditional.If(false, 1, 2)
	assert.Equal(t, 2, result)
}

func TestIfNotNil(t *testing.T) {
	m1 := MyMock{}
	m1.On("MyFunc")

	conditional.IfNotNil[MyMock](nil, m1.MyFunc)
	m1.AssertNotCalled(t, "MyFunc")

	m1.Calls = nil
	conditional.IfNotNil(&m1, m1.MyFunc)
	m1.AssertCalled(t, "MyFunc")
}

func TestIfCall(t *testing.T) {
	m1 := MyMock{}
	m2 := MyMock{}

	m1.On("MyFunc")
	m2.On("MyFunc")

	conditional.IfCall(true, m1.MyFunc, m2.MyFunc)
	m1.AssertCalled(t, "MyFunc")
	m2.AssertNotCalled(t, "MyFunc")

	m1.Calls = nil
	m2.Calls = nil
	conditional.IfCall(false, m1.MyFunc, m2.MyFunc)
	m1.AssertNotCalled(t, "MyFunc")
	m2.AssertCalled(t, "MyFunc")
}

func TestIfCallReturn(t *testing.T) {
	m1 := MyMockR{}
	m2 := MyMockR{}

	m1.On("MyFunc").Return(true)
	m2.On("MyFunc").Return(false)

	result := conditional.IfCallReturn(true, m1.MyFunc, m2.MyFunc)
	assert.True(t, result)
	m1.AssertCalled(t, "MyFunc")
	m2.AssertNotCalled(t, "MyFunc")

	m1.Calls = nil
	m2.Calls = nil

	m1.On("MyFunc").Return(true)
	m2.On("MyFunc").Return(false)

	result = conditional.IfCallReturn(false, m1.MyFunc, m2.MyFunc)
	assert.False(t, result)
	m1.AssertNotCalled(t, "MyFunc")
	m2.AssertCalled(t, "MyFunc")
}

func TestSwitch(t *testing.T) {
	cases := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// matching key returns associated value
	assert.Equal(t, 1, conditional.Switch("one", cases, -1))
	assert.Equal(t, 2, conditional.Switch("two", cases, -1))
	assert.Equal(t, 3, conditional.Switch("three", cases, -1))

	// missing key returns fallback
	assert.Equal(t, -1, conditional.Switch("four", cases, -1))
	assert.Equal(t, -1, conditional.Switch("", cases, -1))

	// empty cases always returns fallback
	assert.Equal(t, -1, conditional.Switch("one", map[string]int{}, -1))

	// fallback zero value
	assert.Equal(t, 0, conditional.Switch("missing", cases, 0))

	// integer keys
	intCases := map[int]string{1: "one", 2: "two"}
	assert.Equal(t, "one", conditional.Switch(1, intCases, "unknown"))
	assert.Equal(t, "unknown", conditional.Switch(99, intCases, "unknown"))

	// struct value type
	type point struct{ x, y int }

	pointCases := map[string]point{
		"origin": {0, 0},
		"unit":   {1, 1},
	}
	assert.Equal(t, point{0, 0}, conditional.Switch("origin", pointCases, point{-1, -1}))
	assert.Equal(t, point{-1, -1}, conditional.Switch("missing", pointCases, point{-1, -1}))

	// key present but mapped to zero value — should return zero value, not fallback
	zeroCases := map[string]int{"zero": 0}
	assert.Equal(t, 0, conditional.Switch("zero", zeroCases, -1))
}
