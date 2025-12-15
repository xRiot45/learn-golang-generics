package learngolanggenerics

import (
	"fmt"
	"testing"
)

func MultipleTypeParameters[T any, U any](param1 T, param2 U) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleTypeParameter(t *testing.T) {
	MultipleTypeParameters[string, int]("Thomas", 123)
	MultipleTypeParameters[float64, bool](45.67, true)
}
