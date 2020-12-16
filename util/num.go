package util

import (
	"log"
)

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

// GetFrequency returns a map denoting how many times each int occur in the slice
func GetFrequency(ints []int) map[int]int {
	m := make(map[int]int)
	for _, i := range ints {
		m[i]++
	}
	return m
}

// FindNumberOfInt returns the number of times i is found in ints
func FindNumberOfInt(i int, ints []int) int {
	found := 0
	for _, v := range ints {
		if v == i {
			found++
		}
	}
	return found
}

// IsUnique returns true iff i only occurs once in ints
func IsUnique(i int, ints []int) bool {
	ni := FindNumberOfInt(i, ints)

	if ni == 1 {
		return true
	} else if ni > 1 {
		return false
	}
	log.Fatal("Found:", ni, "occurrences of:", i)
	return false
}

// RemoveIntByIndex removes the i:th int in ints
func RemoveIntByIndex(idx int, ints []int) []int {
	return append(ints[:idx], ints[idx+1:]...)
}

// RemoveIntByValue removes the first occurrence of i in ints
func RemoveIntByValue(value int, ints []int) []int {
	for i, v := range ints {
		if v == value {
			return RemoveIntByIndex(i, ints)
		}
	}
	return ints
}
