package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := "67384529"

	parsedExampleData := util.GetStringsAsInts(strings.Split(string(exampleData), ""))
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", exampleResult, actual)
	}
}

func TestCreateRing(t *testing.T) {
	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := string(exampleData) + "101112"

	parsedExampleData := util.GetStringsAsInts(strings.Split(string(exampleData), ""))
	ring, _ := createRing(parsedExampleData, 12)
	actual := stringRing(ring)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 149245887792

	parsedExampleData := util.GetStringsAsInts(strings.Split(string(exampleData), ""))
	actual := part2(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", exampleResult, actual)
	}
}
