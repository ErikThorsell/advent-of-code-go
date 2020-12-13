package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example1")
	exampleResult := 295

	ets, table := util.ParseBusTableInput(string(exampleData))
	actual := part1(ets, table)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	for i, f := range []int{1, 2, 3, 4, 5, 6} {
		exampleFile := fmt.Sprintf("./example%d", f)
		exampleData, _ := ioutil.ReadFile(exampleFile)

		exampleResults := []int64{1068781, 3417, 754018, 779210, 1261476, 1202161486}

		_, table := util.ParseBusTableInput(string(exampleData))
		actual := part2(table)

		if exampleResults[i] != actual {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResults[i], actual)
		}
		break
	}
}
