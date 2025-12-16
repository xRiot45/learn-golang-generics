package learngolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 |
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

	assert.Equal(t, Age(25), Min[Age](Age(25), Age(30)))
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, float32(9.99), Min(float32(9.99), float32(19.99)))
	assert.Equal(t, float64(49.99), Min(float64(49.99), float64(59.99)))

	assert.Equal(t, int(1000), Min(1000, 20000))
}
