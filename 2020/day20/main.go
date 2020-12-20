package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

// Tile is a grid with an ID
type Tile struct {
	ID       int
	variants map[int][][]string
	borders  map[int][]string
}

func part1(input []string) int {

	tiles := getTiles(input)

	for _, t := range tiles {
		fmt.Println("TileID:", t.ID)
		for i, v := range t.variants {
			fmt.Println("Variant:", i)
			util.PrintGrid(v)
			fmt.Println()
		}
		for i, b := range t.borders {
			fmt.Println("Border:", i)
			fmt.Println(b)
			fmt.Println()
		}
	}

	// The "only" thing remaining is actually finding the tiling :)

	return -1
}

func getTiles(input []string) []Tile {
	tiles := []Tile{}
	for _, idAndTile := range input {
		tileID := util.GetIntsAsInts(idAndTile)[0]
		grid, _ := util.ParseInputByLineAndRune(strings.Join(strings.Split(idAndTile, "\n")[1:], "\n"))
		variants := getAllGridVariants(grid)
		borders := make(map[int][]string)
		for i, v := range variants {
			borders[i] = getGridBorders(v)
		}
		tiles = append(tiles, Tile{ID: tileID, variants: variants, borders: borders})
	}
	return tiles
}

func getAllGridVariants(grid [][]string) map[int][][]string {

	flippedGrids := getGridFlips(grid)
	rotatedAndFlippedGrids := make([][][]string, 0)
	for _, mt := range flippedGrids {
		candidateGrids := getGridRotations(mt)
		for _, ct := range candidateGrids {
			if !gridInSlice(ct, rotatedAndFlippedGrids) {
				rotatedAndFlippedGrids = append(rotatedAndFlippedGrids, ct)
			}
		}
	}
	gridVariants := make(map[int][][]string)
	for v, grid := range rotatedAndFlippedGrids {
		gridVariants[v] = grid
	}
	return gridVariants
}

func getGridFlips(grid [][]string) [][][]string {

	flips := make([][][]string, 4)

	// The original
	flips[0] = grid

	// Flip around horizontal axis
	vGrid := make([][]string, 0)
	for y := len(grid) - 1; y >= 0; y-- {
		vGrid = append(vGrid, grid[y])
	}
	flips[1] = vGrid

	// Flip around vertical axis
	vGrid = make([][]string, 0)
	for y := 0; y < len(grid); y++ {
		vGrid = append(vGrid, make([]string, 0))
		for x := len(grid[y]) - 1; x >= 0; x-- {
			vGrid[y] = append(vGrid[y], grid[y][x])
		}
	}
	flips[2] = vGrid

	// Flip around vertical then horizontal axis
	vGrid = make([][]string, 0)
	for y := 0; y < len(flips[1]); y++ {
		vGrid = append(vGrid, make([]string, 0))
		for x := len(flips[1]) - 1; x >= 0; x-- {
			vGrid[y] = append(vGrid[y], flips[1][y][x])
		}
	}
	flips[3] = vGrid

	return flips

}

func getGridRotations(grid [][]string) [][][]string {

	rotations := make([][][]string, 4)

	// Original
	rotations[0] = util.CopyGrid(grid)

	// Take the previous grid and rotate it 90 degrees
	for i := 1; i <= 3; i++ {
		grid := util.CopyGrid(rotations[i-1])
		for x := range grid {
			for y := range grid[x] {
				grid[x][y] = rotations[i-1][len(grid[x])-y-1][x]
			}
		}
		rotations[i] = grid
	}
	return rotations
}

func transposeGrid(grid [][]string) [][]string {
	tg := make([][]string, 0)
	for y := 0; y < len(grid); y++ {
		tg = append(grid, make([]string, 0))
		for x := len(grid[y]) - 1; x >= 0; x-- {
			tg[y][x] = grid[x][y]
		}
	}
	return tg
}

func gridInSlice(grid [][]string, slice [][][]string) bool {
	for _, t := range slice {
		if reflect.DeepEqual(t, grid) {
			return true
		}
	}
	return false
}

func getGridBorders(grid [][]string) []string {

	top := strings.Join(grid[0], "")

	right := ""
	for y := 0; y < len(grid); y++ {
		right += grid[y][0]
	}

	left := ""
	for y := 0; y < len(grid); y++ {
		left += grid[y][len(grid)-1]
	}

	bottom := strings.Join(grid[len(grid)-1], "")

	return []string{top, right, bottom, left}
}

func part2(input []string) int {
	return 0
}

func main() {

	rawInput := util.FetchInputForDay("2020", "20")
	parsedInput := util.ParseInputByBlankLine(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
