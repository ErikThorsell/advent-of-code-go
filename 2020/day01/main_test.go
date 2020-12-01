package main

import "testing"

func TestMain(t *testing.T) {

	inputs := []int{1, 2}
	expectedOutputs := []int{0, 0}

	for i := 0; i < len(inputs); i++ {
		actual := part1(inputs)
		if expectedOutputs[i] != actual {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", expectedOutputs, actual)
		}
	}
}
