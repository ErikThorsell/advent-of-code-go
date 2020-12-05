package util

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
