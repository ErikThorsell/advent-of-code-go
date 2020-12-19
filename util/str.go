package util

import (
	"regexp"
	"strconv"
)

// StringInSlice returns true iff a is found in list.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// StringsInSlice returns true iff all strings are present in the slice
func StringsInSlice(strings []string, slice []string) bool {
	for _, s := range strings {
		if !StringInSlice(s, slice) {
			return false
		}
	}
	return true
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

// GetIntsAsStrings return a list of all ints in the string
func GetIntsAsStrings(s string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(s, -1)
}

// GetIntsAsInts return a list of all ints in the string, converted to int
func GetIntsAsInts(s string) []int {
	re := regexp.MustCompile("[0-9]+")
	intsAsStrings := re.FindAllString(s, -1)
	ints := []int{}
	for _, i := range intsAsStrings {
		ints = append(ints, ToInt(i))
	}
	return ints
}

// RemoveDuplicateString returns the set of the slice
func RemoveDuplicateString(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
