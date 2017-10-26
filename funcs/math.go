package funcs

import (
	"sync"

	"github.com/hairyhenderson/gomplate/conv"

	"github.com/hairyhenderson/gomplate/math"
)

var (
	mathNS     *MathFuncs
	mathNSInit sync.Once
)

// MathNS - the math namespace
func MathNS() *MathFuncs {
	mathNSInit.Do(func() { mathNS = &MathFuncs{} })
	return mathNS
}

// AddMathFuncs -
func AddMathFuncs(f map[string]interface{}) {
	f["math"] = MathNS

	f["add"] = MathNS().Add
}

// MathFuncs -
type MathFuncs struct{}

// Add -
func (f *MathFuncs) Add(n ...interface{}) int64 {
	return math.AddInt(conv.ToInt64s(n...)...)
}

// Mul -
func (f *MathFuncs) Mul(n ...interface{}) int64 {
	return math.MulInt(conv.ToInt64s(n...)...)
}

// Sub -
func (f *MathFuncs) Sub(a, b interface{}) int64 {
	return conv.ToInt64(a) - conv.ToInt64(b)
}

// Div -
func (f *MathFuncs) Div(a, b interface{}) int64 {
	return conv.ToInt64(a) / conv.ToInt64(b)
}

// Mod -
func (f *MathFuncs) Mod(a, b interface{}) int64 {
	return conv.ToInt64(a) % conv.ToInt64(b)
}
