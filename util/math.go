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

// MakeRange returns a list with nummbers between min and max (inclusive)
func MakeRange(min int, max int) []int {
	rangeList := make([]int, max-min+1)
	for i := range rangeList {
		rangeList[i] = min + i
	}
	return rangeList
}

// Abs returns the absolute value of x
func Abs(x int) int {
	if x <= 0 {
		return -x
	}
	return x
}

// SumSlice sums all ints in a slice
func SumSlice(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

// GetIntRem returns the reminder of two integers
func GetIntRem(a, b int) float64 {
	return math.Remainder(float64(a), float64(b))
}

// RemIsZero checks if the remainder of two ints is zero
func RemIsZero(a, b int) bool {
	return GetIntRem(a, b) == 0
}
