package funcs

import (
	"sync"

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
	var x int64
	for _, i := range n {
		x += math.ParseInt(i)
	}
	return math.AddInt(n)
}
