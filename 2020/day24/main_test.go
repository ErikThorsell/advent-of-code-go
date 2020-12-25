package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func TestRowParser(t *testing.T) {

	exampleData := "esenee"
	exampleResult := []string{"e", "se", "ne", "e"}

	actual := parseInstruction(exampleData)

	for i := 0; i < len(exampleResult); i++ {
		if exampleResult[i] != actual[i] {
			t.Errorf("Test failed, expected: '%v', got:  '%v'", exampleResult, actual)
		}
	}
}
func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 10

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	actual, _ := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 2208

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	_, tiles := part1(parsedExampleData)
	actual := part2(tiles)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
