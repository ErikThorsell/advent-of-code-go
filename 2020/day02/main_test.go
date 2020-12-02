package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	inputs := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	expectedOutputs := 2

	actual := part1(inputs)

	if expectedOutputs != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expectedOutputs, actual)
	}
}

func TestPart2(t *testing.T) {

	inputs := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	expectedOutputs := 1

	actual := part2(inputs)

	if expectedOutputs != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expectedOutputs, actual)
	}
}
