package main

import (
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test0(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := []int{357, 567, 119, 820}

	parsedExampleData := util.ParseInputByLine(string(exampleData))

	for i, r := range parsedExampleData {
		actual := getSeatID(r)
		if exampleResult[i] != actual {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
		}
	}

}
