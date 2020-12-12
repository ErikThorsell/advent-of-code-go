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

func SliceFromSlice(s []int, is []int) []int {
	newSlice := []int{}
	for _, i := range is {
		newSlice = append(newSlice, s[i])
	}
	return newSlice
}

func RealMod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
