package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 5

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	actual, _, _ := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := "mxmxvkd,sqjhc,fvjkl"

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	_, allergenToIngredients, _ := part1(parsedExampleData)
	actual := part2(allergenToIngredients)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", exampleResult, actual)
	}
}
