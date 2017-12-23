package testutil

import "math"

var eps float64 = 1e-10

func FloatEqual(a, b, epsilon float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

func FloatEqE10(a, b float64) bool {
	return FloatEqual(a, b, eps)
}
