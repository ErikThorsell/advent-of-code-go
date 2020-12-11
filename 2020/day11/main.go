package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func stepModel1(oldSeatGrid [][]string) [][]string {

	newSeatGrid := util.CopyGrid(oldSeatGrid)

	for i := 0; i < len(oldSeatGrid); i++ {
		for j := 0; j < len(oldSeatGrid[i]); j++ {
			newSeatGrid[i][j] = applyRulesPart1(oldSeatGrid, i, j)
		}
	}

	return newSeatGrid

}

func applyRulesPart1(seatGrid [][]string, i int, j int) string {

	seatStatus := seatGrid[i][j]
	numAdjOccSeats := util.CheckAdjacent(seatGrid, i, j, "#")

	if seatStatus == "L" && numAdjOccSeats == 0 {
		return "#"
	}
	if seatStatus == "#" && numAdjOccSeats >= 4 {
		return "L"
	}
	return seatStatus
}

func part1(input [][]string) int {

	oldSeatGrid := util.CopyGrid(input)

	counter := 0
	for {
		counter++

		newSeatGrid := stepModel1(oldSeatGrid)

		if reflect.DeepEqual(newSeatGrid, oldSeatGrid) {
			return util.CountOccurences(newSeatGrid, "#")
		}

		oldSeatGrid = util.CopyGrid(newSeatGrid)

	}

}

func stepModel2(oldSeatGrid [][]string) [][]string {

	newSeatGrid := util.CopyGrid(oldSeatGrid)

	for i := 0; i < len(oldSeatGrid); i++ {
		for j := 0; j < len(oldSeatGrid[i]); j++ {
			newSeatGrid[i][j] = modelSeat2(oldSeatGrid, i, j)
		}
	}

	return newSeatGrid

}

func modelSeat2(seatGrid [][]string, i int, j int) string {

	seatStatus := seatGrid[i][j]
	numOccSeatsInSight := util.CheckLineOfSight(seatGrid, i, j, "#")

	if seatStatus == "L" && numOccSeatsInSight == 0 {
		return "#"
	}
	if seatStatus == "#" && numOccSeatsInSight >= 5 {
		return "L"
	}
	return seatStatus
}

func runSeatModel2Num(oldSeatGrid [][]string, iterations int) [][]string {

	newSeatGrid := util.CopyGrid(oldSeatGrid)
	fmt.Println("i:", 0)
	util.PrintGrid(newSeatGrid)
	fmt.Println()

	for i := 1; i < iterations+1; i++ {
		newSeatGrid = stepModel2(oldSeatGrid)

		fmt.Println("i:", i)
		util.PrintGrid(newSeatGrid)
		fmt.Println()

		oldSeatGrid = util.CopyGrid(newSeatGrid)

	}
	return newSeatGrid
}

func part2(input [][]string) int {

	oldSeatGrid := util.CopyGrid(input)

	counter := 0
	for {
		counter++

		newSeatGrid := stepModel2(oldSeatGrid)

		if reflect.DeepEqual(newSeatGrid, oldSeatGrid) {
			return util.CountOccurences(newSeatGrid, "#")
		}

		oldSeatGrid = util.CopyGrid(newSeatGrid)

	}

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
