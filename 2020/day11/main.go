package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func runSimulation(grid [][]string, part int) [][]string {

	oldSeatGrid := util.CopyGrid(grid)

	for {

		newSeatGrid := generateNewGrid(oldSeatGrid, part)

		if reflect.DeepEqual(newSeatGrid, oldSeatGrid) {
			return newSeatGrid
		}

		oldSeatGrid = util.CopyGrid(newSeatGrid)

	}
}

func runSimulationNTimes(oldSeatGrid [][]string, iterations int, part int) [][]string {

	newSeatGrid := util.CopyGrid(oldSeatGrid)
	fmt.Println("i:", 0)
	util.PrintGrid(newSeatGrid)
	fmt.Println()

	for i := 1; i < iterations+1; i++ {
		newSeatGrid = generateNewGrid(oldSeatGrid, part)

		fmt.Println("i:", i)
		util.PrintGrid(newSeatGrid)
		fmt.Println()

		oldSeatGrid = util.CopyGrid(newSeatGrid)

	}
	return newSeatGrid
}

func generateNewGrid(oldSeatGrid [][]string, part int) [][]string {

	newSeatGrid := util.CopyGrid(oldSeatGrid)

	for i := 0; i < len(oldSeatGrid); i++ {
		for j := 0; j < len(oldSeatGrid[i]); j++ {
			newSeatGrid[i][j] = generateNewCell(oldSeatGrid, i, j, part)
		}
	}

	return newSeatGrid

}

func generateNewCell(seatGrid [][]string, i int, j int, part int) string {

	seatStatus := seatGrid[i][j]
	occThresh := -1
	numOccSeats := -1

	if part == 1 {
		numOccSeats = util.CheckAdjacent(seatGrid, i, j, "#")
		occThresh = 4
	} else if part == 2 {
		numOccSeats = util.CheckLineOfSight(seatGrid, i, j, "#")
		occThresh = 5
	}

	if seatStatus == "L" && numOccSeats == 0 {
		return "#"
	}
	if seatStatus == "#" && numOccSeats >= occThresh {
		return "L"
	}

	return seatStatus
}

func part1(input [][]string) int {

	finalGrid := runSimulation(input, 1)
	return util.CountOccurences(finalGrid, "#")

}

func part2(input [][]string) int {

	finalGrid := runSimulation(input, 2)
	return util.CountOccurences(finalGrid, "#")

}

func main() {

	rawInput := util.FetchInputForDay("2020", "11")
	parsedInput, _ := util.ParseInputByLineAndRune(rawInput)
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
