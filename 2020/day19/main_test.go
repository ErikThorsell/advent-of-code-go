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

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 0

	rules, messages := util.ParseSatelliteInput(string(exampleData))
	actual := part2(rules, messages)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
