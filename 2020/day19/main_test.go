package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example0")
	exampleResult := 0

	rules, messages := util.ParseSatelliteInput(string(exampleData))
	actual := part1(rules, messages)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example2")
	exampleResultPart1 := 3
	exampleResultPart2 := 12

	rules, messages := util.ParseSatelliteInput(string(exampleData))

	actual1 := part1(rules, messages)
	actual2 := part2(rules, messages)

	if exampleResultPart1 != actual1 {
		t.Errorf("Test failed for Part1, expected: '%d', got:  '%d'", exampleResultPart1, actual1)
	}
	if exampleResultPart2 != actual2 {
		t.Errorf("Test failed for Part2, expected: '%d', got:  '%d'", exampleResultPart2, actual2)
	}
}
