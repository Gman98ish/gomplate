package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(12), m.Add(1, 1, 2, 3, 5))
	assert.Equal(t, int64(2), m.Add(1, 1))
	assert.Equal(t, int64(1), m.Add(1))
	assert.Equal(t, int64(0), m.Add(-5, 5))
}

func TestMul(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(30), m.Mul(1, 1, 2, 3, 5))
	assert.Equal(t, int64(1), m.Mul(1, 1))
	assert.Equal(t, int64(1), m.Mul(1))
	assert.Equal(t, int64(-25), m.Mul("-5", 5))
	assert.Equal(t, int64(28), m.Mul(14, "2"))
}

func TestSub(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(0), m.Sub(1, 1))
	assert.Equal(t, int64(-10), m.Sub(-5, 5))
	assert.Equal(t, int64(-41), m.Sub(true, "42"))
}

func TestDiv(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(1), m.Div(1, 1))
	assert.Equal(t, int64(-1), m.Div(-5, 5))
	assert.Equal(t, int64(0), m.Div(true, "42"))
}

func TestMod(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(1), m.Mod(1, 1))
	assert.Equal(t, int64(-1), m.Mod(-5, 5))
	assert.Equal(t, int64(0), m.Mod(true, "42"))
}
