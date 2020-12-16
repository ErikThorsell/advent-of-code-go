package main

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 71

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	ticketConstraints, _, otherTickets := util.ParseTicketInput(parsedExampleData)
	actual, _ := part1(ticketConstraints, otherTickets)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func TestRemoveString(t *testing.T) {

	s := []string{"one", "two", "three", "four", "five"}

	expected := []string{"one", "three", "four", "five"}
	actual := util.RemoveStringByIndex(1, s)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

	expected2 := []string{"one", "three", "five"}
	actual2 := util.RemoveStringByValue("four", expected)

	if !reflect.DeepEqual(expected2, actual2) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected2, actual2)
	}

}
