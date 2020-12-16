package main

import (
	"io/ioutil"
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

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example2")
	exampleResult := 1716

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	ticketConstraints, myTicket, otherTickets := util.ParseTicketInput(parsedExampleData)
	_, invalidTickets := part1(ticketConstraints, otherTickets)
	actual := part2(ticketConstraints, myTicket, otherTickets, invalidTickets)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
