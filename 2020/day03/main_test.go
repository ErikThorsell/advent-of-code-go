package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func TestPart1(t *testing.T) {

	data, _ := ioutil.ReadFile("./example")
	input, numberOfRows := util.ParseInputByLineAndRune(string(data))

	expectedOutput := 7

	actual := part1(input, numberOfRows)

	if expectedOutput != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expectedOutput, actual)
	}
}

func TestPart2(t *testing.T) {

	data, _ := ioutil.ReadFile("./example")
	input, numberOfRows := util.ParseInputByLineAndRune(string(data))

	expectedOutput := 336
	actual := part2(input, numberOfRows)

	if expectedOutput != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expectedOutput, actual)
	}
}
