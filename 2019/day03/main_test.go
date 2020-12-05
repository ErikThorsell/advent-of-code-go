package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleResult := []int{6, 159, 135, 2}

	for _, i := range []int{1, 2, 3, 4} {
		exampleFile := fmt.Sprintf("./example%v", i)
		fmt.Println("Testing with:", exampleFile)
		exampleData, _ := ioutil.ReadFile(exampleFile)
		parsedExampleData := util.ParseInputByLineAndSep(string(exampleData), ',')

		actual := part1(parsedExampleData)

		if exampleResult[i-1] != actual {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult[i-1], actual)
		}
	}
}

func Test2(t *testing.T) {

	exampleResult := []int{30, 610, 410}

	for _, i := range []int{1, 2, 3} {
		exampleFile := fmt.Sprintf("./example%v", i)
		fmt.Println("Testing with:", exampleFile)
		exampleData, _ := ioutil.ReadFile(exampleFile)
		parsedExampleData := util.ParseInputByLineAndSep(string(exampleData), ',')

		actual := part2(parsedExampleData)

		if exampleResult[i-1] != actual {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult[i-1], actual)
		}
	}
}
