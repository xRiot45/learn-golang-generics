package learngolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestMin(t *testing.T) {
	assert.Equal(t, int(10), Min[int](10, 20))
	assert.Equal(t, int8(5), Min[int8](5, 15))
	assert.Equal(t, int16(100), Min[int16](100, 200))
	assert.Equal(t, int32(1000), Min[int32](1000, 2000))
	assert.Equal(t, int64(10000), Min[int64](10000, 20000))
}
