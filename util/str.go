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
