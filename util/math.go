package util

import "math"

// Max returns the bigger int of x and y
func Max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

// Min returns the smaller int of x and y
func Min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

// Abs returns the absolute value of x
func Abs(x int) int {
	if x <= 0 {
		return -x
	}
	return x
}

// GetIntRem returns the reminder of two integers
func GetIntRem(a, b int) float64 {
	return math.Remainder(float64(a), float64(b))
}

// RemIsZero checks if the remainder of two ints is zero
func RemIsZero(a, b int) bool {
	return GetIntRem(a, b) == 0
}

// RealMod computes the mod like Python do
func RealMod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
