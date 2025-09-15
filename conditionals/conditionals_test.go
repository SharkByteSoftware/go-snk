package conditionals_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/conditionals"
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
	result := conditionals.If(true, 1, 2)
	assert.Equal(t, 1, result)

	result = conditionals.If(false, 1, 2)
	assert.Equal(t, 2, result)
}

func TestIfCall(t *testing.T) {
	m1 := MyMock{}
	m2 := MyMock{}
	m1.On("MyFunc")
	m2.On("MyFunc")

	conditionals.IfCall(true, m1.MyFunc, m2.MyFunc)
	m1.AssertCalled(t, "MyFunc")
	m2.AssertNotCalled(t, "MyFunc")

	m1.Calls = nil
	m2.Calls = nil
	conditionals.IfCall(false, m1.MyFunc, m2.MyFunc)
	m1.AssertNotCalled(t, "MyFunc")
	m2.AssertCalled(t, "MyFunc")
}

func TestIfCallReturn(t *testing.T) {
	m1 := MyMockR{}
	m2 := MyMockR{}
	m1.On("MyFunc").Return(true)
	m2.On("MyFunc").Return(false)

	result := conditionals.IfCallReturn(true, m1.MyFunc, m2.MyFunc)
	assert.True(t, result)
	m1.AssertCalled(t, "MyFunc")
	m2.AssertNotCalled(t, "MyFunc")

	m1.Calls = nil
	m2.Calls = nil
	m1.On("MyFunc").Return(true)
	m2.On("MyFunc").Return(false)

	result = conditionals.IfCallReturn(false, m1.MyFunc, m2.MyFunc)
	assert.False(t, result)
	m1.AssertNotCalled(t, "MyFunc")
	m2.AssertCalled(t, "MyFunc")
}
