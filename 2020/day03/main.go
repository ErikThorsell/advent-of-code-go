package main

import (
	"fmt"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func printTreeGrid(grid [][]string, d int, r int) {

	gridP := make([][]string, len(grid), len(grid[0]))
	copy(gridP, grid)

	if gridP[d][r] == "#" {
		gridP[d][r] = "X"
	} else {
		gridP[d][r] = "O"
	}

	fmt.Println(d, r)
	for _, row := range gridP {
		fmt.Println(row)
	}
	fmt.Println()
}

func run(grid [][]string, rows int, down int, right int) int {

	hitTrees := 0
	d := 0
	r := 0

	for {
		d += down
		r = (r + right) % len(grid[0])

		if grid[d][r] == "#" {
			hitTrees++
		}

		//printTreeGrid(grid, d, r)

		if d == rows-1 {
			return hitTrees
		}

	}

	return hitTrees

}

func part1(grid [][]string, rows int) int {
	return run(grid, rows, 1, 3)
}

func part2(grid [][]string, rows int) int {
	downVersions := []int{1, 1, 1, 1, 2}
	rightVersions := []int{1, 3, 5, 7, 1}

	treeProduct := 1
	for i := 0; i < len(downVersions); i++ {
		treesInRun := run(grid, rows, downVersions[i], rightVersions[i])
		treeProduct *= treesInRun
	}
	return treeProduct
}

func main() {

	rawInput := util.FetchInputForDay("2020", "3")
	parsedInput, numberOfRows := util.ParseInputByLineAndRune(rawInput)

	ans1 := part1(parsedInput, numberOfRows)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(parsedInput, numberOfRows)
	fmt.Println("Answer for second question: ", ans2)

}
