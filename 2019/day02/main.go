package main

import (
	"fmt"
	"strconv"
	"strings"

    "github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(integers []int) int {

	integers[1] = 12
	integers[2] = 2

	integers = util.RunProgram(integers)

	return integers[0]

}

func part2(modules []int) int {
	return 0
}

func main() {

	input := util.FetchInputForDay("2019", "2")
	rawData := strings.Split(input, ",")

	var intData []int
	for _, x := range rawData {
		v, _ := strconv.Atoi(x)
		intData = append(intData, v)
	}

	ans1 := part1(intData)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(intData)
	fmt.Println("Answer for second question: ", ans2)

}
