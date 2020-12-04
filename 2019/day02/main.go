package main

import (
	"fmt"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(integers []int) int {

	program := make([]int, len(integers))
	copy(program, integers)

	program[1] = 12
	program[2] = 2

	program = util.RunProgram(program)

	return program[0]

}

func part2(integers []int) int {

	var noun int
	var verb int

	program := make([]int, len(integers))

	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {

			copy(program, integers)

			program[1] = noun
			program[2] = verb

			program = util.RunProgram(program)

			if program[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func main() {

	input := util.FetchInputForDay("2019", "2")
	intData := util.ParseInputBySepToInts(input, ',')

	ans1 := part1(intData)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(intData)
	fmt.Println("Answer for second question: ", ans2)

}
