package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func isValidNumber(number int, index int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == number {
				return true
			}
		}
	}
	return false
}

func part1(input []int, preambleLength int) int {

	for i := 0; i < len(input)-preambleLength; i++ {

		startOfNumbers := preambleLength + i
		preamble := input[i:startOfNumbers]

		if !isValidNumber(input[startOfNumbers], i, preamble) {
			return input[startOfNumbers]
		}
	}

	return -1

}

func findContagiousSet(invNum int, numbers []int) []int {

	start := 0
	end := 2

	for {
		candidateSet := numbers[start:end]
		candidateSum := util.SumSlice(candidateSet)
		if candidateSum == invNum {
			return candidateSet
		} else if candidateSum < invNum {
			end++
		} else {
			start++
		}
	}
}

func part2(input []int, preambleLength int) int {
	invalidNumber := part1(input, preambleLength)
	set := findContagiousSet(invalidNumber, input)
	if len(set) > 0 {
		sort.Ints(set)
		return set[0] + set[len(set)-1]
	}
	return -1
}

func main() {

	rawInput := util.FetchInputForDay("2020", "9")
	parsedInput := util.ParseInputBySepToInts(rawInput, '\n')
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput, 25)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("Answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput, 25)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Answer retrieved in: ", e)

}
