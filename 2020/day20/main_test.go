package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test0(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example0")
	exampleResult := 0

	parsedExampleData := util.ParseInputByBlankLine(string(exampleData))
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 0

	parsedExampleData := util.ParseInputByBlankLine(string(exampleData))
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 0

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	actual := part2(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
