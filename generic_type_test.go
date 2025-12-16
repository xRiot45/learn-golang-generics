package learngolanggenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, item := range bag {
		fmt.Println(item)
	}
}

func TestPrintBagString(t *testing.T) {
	bag := Bag[string]{"Apple", "Banana", "Cherry"}
	PrintBag[string](bag)
}

func TestPrintBagInt(t *testing.T) {
	bag := Bag[int]{10, 20, 30}
	PrintBag[int](bag)
}

func TestPrintBagFloat(t *testing.T) {
	bag := Bag[float64]{9.99, 19.99, 29.99}
	PrintBag[float64](bag)
}
