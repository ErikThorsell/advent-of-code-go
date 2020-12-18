package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	for i, f := range []int{1, 2, 3, 4} {

		exampleFile := fmt.Sprintf("./example%d", f)
		exampleData, _ := ioutil.ReadFile(exampleFile)
		parsedExampleData := util.ParseInputByLine(string(exampleData))

		exampleResults := []int{71, 51, 26, 13632}

		actual := part1(parsedExampleData)

		if exampleResults[i] != actual {
			t.Errorf("Test failed for %v, expected: '%d', got:  '%d'", exampleFile, exampleResults[i], actual)
		} else {
			t.Logf("Test successful for %v", exampleFile)
		}
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
