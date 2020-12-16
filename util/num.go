package util

import (
	"fmt"
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

// GetMostFrequent return the most frequently occurring int in a slice
func GetMostFrequent(ints []int) int {

	var mostFreq int
	var freq int

	frequency := GetFrequency(ints)
	fmt.Println("FL:", frequency)

	for n, f := range frequency {
		if f > freq {
			freq = f
			mostFreq = n
		}
	}

	return mostFreq

}

func FindNumberOfInt(i int, ints []int) int {
	found := 0
	for _, v := range ints {
		if v == i {
			found++
		}
	}
	return found
}

// GetMostFrequentForIndex looks at a number of maps and determines which key is the uniquely
// highest for each key.
//
// Example:
//
// Given these two maps as input, the identifies that 14 is uniquely highest for the first map.
// map[0:189 1:189 2:189 3:189 4:189 5:189 6:189 7:189 8:189 9:189 10:189 11:189 12:189 13:189 14:190 15:189 16:189 17:189 18:189 19:189]
// map[0:189 1:190 2:189 3:189 4:189 5:189 6:189 7:189 8:189 9:189 10:189 11:189 12:189 13:189 14:190 15:189 16:189 17:189 18:189 19:189]
//
// The function notes that the then removes key 14 from all other maps:
// map[0:189 1:190 2:189 3:189 4:189 5:189 6:189 7:189 8:189 9:189 10:189 11:189 12:189 13:189 15:189 16:189 17:189 18:189 19:189]
// Now, 2 is uniquely highest.
//
// The function returns, for each key, the index that identifies that key.
func GetMostFrequentForIndex(keyToRange map[string][]int) map[int]int {

	keyToIndex := make(map[int]int)

	keysLeft := []int{}
	for k := range keyToRange {
		keysLeft = append(keysLeft, k)
	}

	for {

		if len(keysLeft) == 0 {
			return keyToIndex
		}

		for _, r := range keysLeft {
			// Is there a uniqely highest idx in the list?
			mostFreqCand := GetMostFrequent(keyToRange[r])
			if isUnique(mostFreqCand, keyToRange[r]) {
				keyToIndex[r] = mostFreqCand
				keysLeft = remove(r, keysLeft)
			}
		}
	}

}

func remove(i int, s []int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func isUnique(i int, ints []int) bool {
	ni := FindNumberOfInt(i, ints)

	if ni == 1 {
		return true
	} else if ni > 1 {
		return false
	}
	log.Fatal("Found:", ni, "occurrences of:", i)
	return false
}
