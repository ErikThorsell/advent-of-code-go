package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(ticketConstraints map[string][][]int, otherTickets [][]int) (int, [][]int) {
	validNumbers := findValidNumbers(ticketConstraints)
	invalidNumbers, invalidTickets := getInvalidTicketNumbers(otherTickets, validNumbers)
	return util.SumSlice(invalidNumbers), invalidTickets
}

func findValidNumbers(ticketConstraints map[string][][]int) []int {
	validNumbers := []int{}
	for _, v := range ticketConstraints {
		validNumbers = append(validNumbers, expandRange(v)...)
	}
	return validNumbers
}

func expandRange(rov [][]int) []int {
	rangeOfNumbers := []int{}
	for _, pair := range rov {
		for i := pair[0]; i <= pair[1]; i++ {
			rangeOfNumbers = append(rangeOfNumbers, i)
		}
	}
	return rangeOfNumbers
}

func getInvalidTicketNumbers(tickets [][]int, validNumbers []int) ([]int, [][]int) {
	invalidTickets := [][]int{}
	invalidNumbers := []int{}
	for _, t := range tickets {
		for _, n := range t {
			if !util.IntInSlice(n, validNumbers) {
				invalidNumbers = append(invalidNumbers, n)
				invalidTickets = append(invalidTickets, t)
			}
		}
	}
	return invalidNumbers, invalidTickets
}

func part2(ticketConstraints map[string][][]int, myTicket []int, otherTickets [][]int, invalidTickets [][]int) int {
	validTickets := getValidTickets(otherTickets, invalidTickets)
	constraintToRange := addIndiciesToConstraints(ticketConstraints, validTickets)
	constraintToIndex := getIndexForEachConstraint(constraintToRange)
	return computeDepartureProduct(constraintToIndex, myTicket)
}

func getValidTickets(candidateTickets [][]int, invalidTickets [][]int) [][]int {
	validTickets := [][]int{}
	valid := true
	for _, ct := range candidateTickets {
		for _, it := range invalidTickets {
			if reflect.DeepEqual(ct, it) {
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ct)
		}
		valid = true
	}
	return validTickets
}

func addIndiciesToConstraints(constraints map[string][][]int, tickets [][]int) map[string][]int {

	constraintToRange := make(map[string][]int)

	for c, v := range constraints {
		lowList := util.MakeRange(v[0][0], v[0][1])
		highList := util.MakeRange(v[1][0], v[1][1])
		vRange := append(lowList, highList...)

		availIds := util.MakeRange(0, len(tickets[0])-1)

		for _, ticket := range tickets {
			for valueIdx, ticketValue := range ticket {
				if !util.IntInSlice(ticketValue, vRange) {
					availIds = util.RemoveIntByValue(valueIdx, availIds)
				}
			}
		}
		constraintToRange[c] = append(constraintToRange[c], availIds...)
	}
	return constraintToRange
}

func getIndexForEachConstraint(constraintToRange map[string][]int) map[string]int {

	constraintToIndex := make(map[string]int)

	for {

		if len(constraintToRange) == 0 {
			return constraintToIndex
		}

		for c, r := range constraintToRange {

			if len(r) == 1 {
				constraintToIndex[c] = r[0]
				delete(constraintToRange, c)
				constraintToRange = removePossibleIndex(constraintToRange, r[0])
			}
		}
	}
}

func removePossibleIndex(constraintToRange map[string][]int, idx int) map[string][]int {
	newMap := make(map[string][]int)
	for c, r := range constraintToRange {
		newMap[c] = util.RemoveIntByValue(idx, r)
	}
	return newMap
}

func computeDepartureProduct(constraintIndex map[string]int, ticket []int) int {
	prod := 1
	for c, i := range constraintIndex {
		if strings.Contains(c, "departure") {
			prod *= ticket[i]
		}
	}
	return prod
}

func main() {

	rawInput := util.FetchInputForDay("2020", "16")
	parsedInput := util.ParseInputByLine(string(rawInput))
	ticketConstraints, myTicket, otherTickets := util.ParseTicketInput(parsedInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1, invalidTickets := part1(ticketConstraints, otherTickets)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(ticketConstraints, myTicket, otherTickets, invalidTickets)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
