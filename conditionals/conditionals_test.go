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

func (m *MyMock) MyFunc() bool {
	m.Called()
	return true
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
	m1.On("MyFunc").Return(true)
	m2.On("MyFunc").Return(false)

	_ = conditionals.IfCall(true, m1.MyFunc, m2.MyFunc)
	m1.AssertCalled(t, "MyFunc")
	m2.AssertNotCalled(t, "MyFunc")

	m1.Calls = nil
	m2.Calls = nil
	_ = conditionals.IfCall(false, m1.MyFunc, m2.MyFunc)
	m1.AssertNotCalled(t, "MyFunc")
	m2.AssertCalled(t, "MyFunc")

}
