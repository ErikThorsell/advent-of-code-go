package main

import (
	"fmt"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input [][]string) int {
	return 0
}

func part2(input [][]string) int {
	return 0
}

func main() {

	rawInput := util.FetchInputForDay("2020", "0")
	parsedInput := util.ParseInputByLineAndSep(rawInput, ',')
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("Answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Answer retrieved in: ", e)

}
