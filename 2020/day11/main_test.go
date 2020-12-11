package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 37

	parsedExampleData, _ := util.ParseInputByLineAndRune(string(exampleData))
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 26

	parsedExampleData, _ := util.ParseInputByLineAndRune(string(exampleData))
	actual := part2(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test21(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	parsedExampleData, _ := util.ParseInputByLineAndRune(string(exampleData))

	exampleResult, _ := ioutil.ReadFile("./example2-1rot")
	parsedExampleResult, _ := util.ParseInputByLineAndRune(string(exampleResult))

	actual := runSimulationNTimes(parsedExampleData, 1, 2)

	if !reflect.DeepEqual(actual, parsedExampleResult) {
		t.Errorf("Test failed!")
		fmt.Println("Expected:")
		util.PrintGrid(parsedExampleResult)
		fmt.Println("\nGot:")
		util.PrintGrid(actual)
	}
}

func Test22(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	parsedExampleData, _ := util.ParseInputByLineAndRune(string(exampleData))

	exampleResult, _ := ioutil.ReadFile("./example2-2rot")
	parsedExampleResult, _ := util.ParseInputByLineAndRune(string(exampleResult))

	actual := runSimulationNTimes(parsedExampleData, 2, 2)

	if !reflect.DeepEqual(actual, parsedExampleResult) {
		t.Errorf("Test failed!")
		fmt.Println("Expected:")
		util.PrintGrid(parsedExampleResult)
		fmt.Println("\nGot:")
		util.PrintGrid(actual)
	}
}

func Test2CheckLineOfSight1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./exampleLOS-1")
	exampleResult := 8

	parsedExampleData, _ := util.ParseInputByLineAndRune(string(exampleData))
	actual := util.CheckLineOfSight(parsedExampleData, 4, 3, "#")

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2CheckLineOfSight2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./exampleLOS-2")
	exampleResult := 0

	parsedExampleData, _ := util.ParseInputByLineAndRune(string(exampleData))
	actual := util.CheckLineOfSight(parsedExampleData, 1, 1, "#")

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
