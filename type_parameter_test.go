package learngolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var resultString string = Length("Thomas")
	assert.Equal(t, "Thomas", resultString)

	var resultInt int = Length(123)
	assert.Equal(t, 123, resultInt)

	var resultFloat float64 = Length(45.67)
	assert.Equal(t, 45.67, resultFloat)

	var resultBool bool = Length(true)
	assert.Equal(t, true, resultBool)
}
