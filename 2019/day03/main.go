package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input [][]string) int {

	// Convert the instructions into wires (lists of Coordinate)
	var wires [][]util.Coordinate
	for _, wire := range input {
		wires = append(wires, util.ManhattanWalk(wire))
	}

	crossingWires := util.ManhattanWalksIntersect(wires[0], wires[1])

	var distances []int
	for _, c := range crossingWires {
		distances = append(distances, util.ManhattanDistance(util.Coordinate{0, 0}, c))
	}

	sort.Ints(distances)

	return distances[0]
}

func part2(input [][]string) int {
	// Convert the instructions into wires (lists of Coordinate)
	var wires [][]util.CoordinateWithDistance
	for _, wire := range input {
		wires = append(wires, util.ManhattanWalkWithDistance(wire))
	}

	crossingWires := util.ManhattanWalksIntersectWithDistance(wires[0], wires[1])

	sort.Ints(crossingWires)

	return crossingWires[0]

}

func main() {

	rawInput := util.FetchInputForDay("2019", "3")
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
