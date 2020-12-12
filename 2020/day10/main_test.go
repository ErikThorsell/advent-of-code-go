package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 35

	parsedExampleData := util.ParseInputBySepToInts(string(exampleData), '\n')
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test12(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example2")
	exampleResult := 220

	parsedExampleData := util.ParseInputBySepToInts(string(exampleData), '\n')
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 8

	parsedExampleData := util.ParseInputBySepToInts(string(exampleData), '\n')
	actual := part2(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
