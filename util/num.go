package util

// IntInSlice returns true iff a is found in list.
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// MakeRange returns a list with nummbers between min and max (inclusive)
func MakeRange(min int, max int) []int {
	rangeList := make([]int, max-min+1)
	for i := range rangeList {
		rangeList[i] = min + i
	}
	return rangeList
}

// SumSlice sums all ints in a slice
func SumSlice(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}
