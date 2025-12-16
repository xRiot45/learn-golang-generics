package learngolanggenerics

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Hello",
		Second: "World",
	}

	fmt.Println(data)
}
