package util

import "strconv"

// StringInSlice returns true iff a is found in list.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IsInt returns true if s looks like an int
func IsInt(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

// RemoveStringByValue removes the i:th string
func RemoveStringByValue(s string, xs []string) []string {
	for i, v := range xs {
		if v == s {
			return append(xs[:i], xs[i+1:]...)
		}
	}
	return xs
}

// RemoveStringByIndex removes the i:th string
func RemoveStringByIndex(i int, s []string) []string {
	return append(s[:i], s[i+1:]...)
}
