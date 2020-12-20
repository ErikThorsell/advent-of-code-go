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
	ID   int
	Grid [][]string
}

func part1(input []string) int {

	tiles := getTiles(input)

	tileVariants := make(map[int][]Tile)
	for tileID := range tiles {
		tileVariants[tileID] = getAllTileOptions(tiles[tileID])
	}

	tileVariantsBorders := make(map[int]map[int][]string)
	for tileID, variants := range tileVariants {
		borders := make(map[int][]string)
		for variantID, variant := range variants {
			tileBorders := computeTileBorders(variant)
			borders[variantID] = tileBorders
		}
		tileVariantsBorders[tileID] = borders
	}

	// The "only" thing remaining is actually finding the tiling :)

	return -1
}

func getTiles(input []string) map[int]Tile {
	tiles := make(map[int]Tile)
	for _, idAndTile := range input {
		id := util.GetIntsAsInts(idAndTile)[0]
		grid, _ := util.ParseInputByLineAndRune(strings.Join(strings.Split(idAndTile, "\n")[1:], "\n"))
		tiles[id] = Tile{ID: id, Grid: grid}
	}
	return tiles
}

func getAllTileOptions(tile Tile) []Tile {
	flippedTiles := getTileFlips(tile)
	rotatedAndFlippedTiles := make([]Tile, 0)
	for _, mt := range flippedTiles {
		candidateTiles := getTileRotations(mt)
		for _, ct := range candidateTiles {
			if !tileInSlice(ct, rotatedAndFlippedTiles) {
				rotatedAndFlippedTiles = append(rotatedAndFlippedTiles, ct)
			}
		}
	}
	return rotatedAndFlippedTiles
}

func getTileFlips(tile Tile) []Tile {

	flips := make([]Tile, 4)

	// The original
	flips[0] = tile

	// Flip around horizontal axis
	grid := make([][]string, 0)
	for y := len(tile.Grid) - 1; y >= 0; y-- {
		grid = append(grid, tile.Grid[y])
	}
	flips[1] = Tile{Grid: grid}

	// Flip around vertical axis
	grid = make([][]string, 0)
	for y := 0; y < len(tile.Grid); y++ {
		grid = append(grid, make([]string, 0))
		for x := len(tile.Grid[y]) - 1; x >= 0; x-- {
			grid[y] = append(grid[y], tile.Grid[y][x])
		}
	}
	flips[2] = Tile{Grid: grid}

	// Flip around vertical then horizontal axis
	grid = make([][]string, 0)
	for y := 0; y < len(flips[1].Grid); y++ {
		grid = append(grid, make([]string, 0))
		for x := len(flips[1].Grid[y]) - 1; x >= 0; x-- {
			grid[y] = append(grid[y], flips[1].Grid[y][x])
		}
	}
	flips[3] = Tile{Grid: grid}

	return flips

}

func getTileRotations(tile Tile) []Tile {

	rotations := make([]Tile, 4)

	// Original
	rotations[0] = Tile{Grid: util.CopyGrid(tile.Grid)}

	// Take the previous grid and rotate it 90 degrees
	for i := 1; i <= 3; i++ {
		grid := util.CopyGrid(rotations[i-1].Grid)
		for x := range grid {
			for y := range grid[x] {
				grid[x][y] = rotations[i-1].Grid[len(grid[x])-y-1][x]
			}
		}
		rotations[i] = Tile{Grid: grid}
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

func tileInSlice(tile Tile, slice []Tile) bool {
	for _, t := range slice {
		if reflect.DeepEqual(t.Grid, tile.Grid) {
			return true
		}
	}
	return false
}

func computeTileBorders(tile Tile) []string {

	grid := tile.Grid

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
