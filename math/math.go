package math

// AddInt -
func AddInt(n ...int64) int64 {
	var x int64
	for _, i := range n {
		x += i
	}
	return x
}

// SubInt -
func SubInt(x, y int64) int64 {
	return x - y
}

// MultInt -
func MultInt(n ...int64) int64 {
	var x int64 = 1
	for _, i := range n {
		x *= i
	}
	return x
}

// DivInt -
func DivInt(x, y int64) int64 {
	return x / y
}

// AddFloat -
func AddFloat(x, y float64) float64 {
	return x + y
}

// SubFloat -
func SubFloat(x, y float64) float64 {
	return x - y
}

// MultFloat -
func MultFloat(x, y float64) float64 {
	return x * y
}

// DivFloat -
func DivFloat(x, y float64) float64 {
	return x / y
}
