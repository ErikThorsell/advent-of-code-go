package main

import (
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {

	inputs := [][]int{{1721, 979, 366, 299, 675, 1456}}
	expectedOutputs := []int{514579}

	for i := 0; i < len(inputs); i++ {
		actual := part1(inputs[i])
		if !reflect.DeepEqual(expectedOutputs[i], actual) {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", expectedOutputs[i], actual)
		}
	}
}

func TestPart2(t *testing.T) {

	inputs := [][]int{{1721, 979, 366, 299, 675, 1456}}
	expectedOutputs := []int{241861950}

	for i := 0; i < len(inputs); i++ {
		actual := part2(inputs[i])
		if !reflect.DeepEqual(expectedOutputs[i], actual) {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", expectedOutputs[i], actual)
		}
	}
}
