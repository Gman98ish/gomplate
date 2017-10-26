package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert.Equal(t, int64(3), AddInt(1, 2))
	assert.Equal(t, int64(-1), SubInt(2, 3))
	assert.Equal(t, int64(12), MultInt(3, 4))
	assert.Equal(t, int64(2), DivInt(4, 2))
	assert.Equal(t, float64(1.5), AddFloat(1, 0.5))
	assert.Equal(t, float64(-6), SubFloat(2.5, 8.5))
	assert.Equal(t, float64(12.0), MultFloat(3, 4))
	assert.Equal(t, float64(2.5), DivFloat(5, 2))
}

// func TestParseInt(t *testing.T) {
// 	assert.Equal(t, int64(1), ParseInt(1))
// 	assert.Equal(t, int64(1), ParseInt(int32(1)))
// 	assert.Equal(t, int64(1), ParseInt(int64(1)))
// 	assert.Equal(t, int64(1), ParseInt(float32(1)))
// 	assert.Equal(t, int64(1), ParseInt(float64(1)))
// 	assert.Equal(t, int64(42), ParseInt(42))
// }

func TestParseInts(t *testing.T) {
	// in := []string{}
	expected := []int64{}
	assert.Equal(t, expected, ParseInts())

	assert.Equal(t, []int64{0}, ParseInts(""))
	assert.Equal(t, []int64{0}, ParseInts("0"))
	assert.Equal(t, []int64{42, 15}, ParseInts("42", "15"))
	assert.Equal(t, []int64{0, 0, 0, 1, 1, 2, 3, 5, 8, 13},
		ParseInts(nil, false, "", true, 1, 2.0, uint8(3), int64(5), float32(8), "13"))
}
