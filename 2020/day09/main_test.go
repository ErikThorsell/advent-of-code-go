package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 127

	parsedExampleData := util.ParseInputBySepToInts(string(exampleData), '\n')
	actual := part1(parsedExampleData, 5)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 62

	parsedExampleData := util.ParseInputBySepToInts(string(exampleData), '\n')
	actual := part2(parsedExampleData, 5)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
