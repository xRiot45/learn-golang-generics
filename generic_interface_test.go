package learngolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	value T
}

func (m *MyData[T]) GetValue() T {
	return m.value
}

func (m *MyData[T]) SetValue(value T) {
	m.value = value
}

func TestGeneritInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "Thomas")

	assert.Equal(t, "Thomas", result)
	assert.Equal(t, "Thomas", myData.value)
}
