package main

import (
	"fmt"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []int) int {
	return 0
}

func part2(input []int) int {
	return 0
}

func main() {

	rawInput := util.FetchInputForDay("2020", "0")
	parsedInput := util.ParseInputToListOfInts(rawInput, '\n')

	ans1 := part1(parsedInput)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(parsedInput)
	fmt.Println("Answer for second question: ", ans2)

}
