package main

import (
	"fmt"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []int) int {

	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {

			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}

		}
	}
	return -1
}

func part2(input []int) int {

	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			for k := j; k < len(input); k++ {

				if input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}

			}
		}
	}
	return -1
}

func main() {

	rawInput := util.FetchInputForDay("2020", "1")
	parsedInput := util.ParseInputToListOfInts(rawInput, '\n')

	ans1 := part1(parsedInput)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(parsedInput)
	fmt.Println("Answer for second question: ", ans2)

}
