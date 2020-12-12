package main

import (
	"fmt"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string) int {
	return 0
}

func part2(input []string) int {
	return 0
}

func main() {

	rawInput := util.FetchInputForDay("2020", "0")
	parsedInput := util.ParseInputByLine(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
