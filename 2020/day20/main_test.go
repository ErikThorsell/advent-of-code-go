package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 20899048083289

	parsedExampleData := util.ParseInputByBlankLine(string(exampleData))
	actual, _, _ := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 273

	parsedExampleData := util.ParseInputByBlankLine(string(exampleData))
	_, tiles, idGrid := part1(parsedExampleData)
	actual := part2(tiles, idGrid)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
