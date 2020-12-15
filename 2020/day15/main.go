package main

import (
	"fmt"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []int, end int) int {

	mem := make(map[int][]int)
	var lastNum int

	counter := 0
	for {

		if end < 100 {
			fmt.Println("T:", counter, "LN:", lastNum, "Mem:", mem)
		}

		if counter == end {
			return lastNum
		}

		if counter < len(input) {
			mem[input[counter]] = append(mem[input[counter]], counter)
			lastNum = input[counter]

		} else {

			if len(mem[lastNum]) == 1 { // first time spoken
				lastNum = 0
				mem[lastNum] = append(mem[lastNum], counter)

			} else {
				lastNum = mem[lastNum][len(mem[lastNum])-1] - mem[lastNum][len(mem[lastNum])-2]
				mem[lastNum] = append(mem[lastNum], counter)
			}

		}

		counter++

	}
}

func part2(input []int, end int) int {

	mem := make(map[int][]int)
	var lastNum int
	var counter int

	for counter = 0; counter < len(input); counter++ {
		mem[input[counter]] = []int{-1, counter}
		lastNum = input[counter]
	}

	for {

		if counter == end {
			return lastNum
		}

		if mem[lastNum][0] == -1 {
			lastNum = 0
			mem[lastNum] = []int{mem[lastNum][1], counter}
		} else {
			prev := mem[lastNum][1]
			preprev := mem[lastNum][0]
			lastNum = prev - preprev
			_, ok := mem[lastNum] // first time?
			if !ok {
				mem[lastNum] = []int{-1, counter} // Yes
			} else {
				mem[lastNum] = []int{mem[lastNum][1], counter} // No
			}
		}
		counter++
	}
}

func main() {

	rawInput := util.FetchInputForDay("2020", "15")
	parsedInput := util.ParseInputBySepToInts(rawInput, ',')
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput, 2020)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput, 30000000)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
