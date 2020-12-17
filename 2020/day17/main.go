package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string, iterations int) int {

	activeCubes := []util.Cube{}
	for y, row := range input {
		for x, cube := range strings.Split(row, "") {
			if cube == "#" {
				activeCubes = append(activeCubes, util.Cube{X: x, Y: y, Z: 0})
			}
		}
	}

	for i := 0; i < iterations; i++ {
		activeCubes = util.SimulateBlockOfCubes(activeCubes)
	}

	return len(activeCubes)

}

func part2(input []string, iterations int) int {
	activeCubes := []util.HyperCube{}
	for y, row := range input {
		for x, cube := range strings.Split(row, "") {
			if cube == "#" {
				activeCubes = append(activeCubes, util.HyperCube{X: x, Y: y, Z: 0, W: 0})
			}
		}
	}

	for i := 0; i < iterations; i++ {
		activeCubes = util.SimulateBlockOfHyperCubes(activeCubes)
	}

	return len(activeCubes)
}

func main() {

	rawInput := util.FetchInputForDay("2020", "17")
	parsedInput := util.ParseInputByLine(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput, 6)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput, 6)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
