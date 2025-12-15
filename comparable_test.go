package learngolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSameType[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.True(t, IsSameType[int](10, 10))
	assert.False(t, IsSameType[int](10, 20))

	assert.True(t, IsSameType[string]("Golang", "Golang"))
	assert.False(t, IsSameType[string]("Golang", "Java"))

	assert.True(t, IsSameType[float64](99.99, 99.99))
	assert.False(t, IsSameType[float64](99.99, 88.88))
}
