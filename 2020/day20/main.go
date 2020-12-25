package main

import (
	"fmt"
	"math"
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

	idGrid := make([][]int, 3)
	for i := 0; i < 3; i++ {
		idGrid[i] = make([]int, 3)
	}

	grid := make([][]string, 10*int(math.Sqrt(float64(len(tiles)))))
	for g := 0; g < len(grid); g++ {
		grid[g] = make([]string, len(grid))
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x] = "x"
		}
	}

	fmt.Printf("Grid Size: (%v,%v)\n", len(grid), len(grid[0]))

	usedTiles := []Tile{}

	grid, idGrid = findTiling(tiles, usedTiles, grid, idGrid, 0, 0)

	fmt.Println("RETURNED")
	util.PrintGrid(grid)
	fmt.Println(idGrid)

	return -1
}

func findTiling(tiles []Tile, usedTiles []Tile, grid [][]string, idGrid [][]int, ypos, xpos int) ([][]string, [][]int) {

	fmt.Println("GRID")
	util.PrintGrid(grid)
	fmt.Println(idGrid)
	fmt.Println()

	if ypos == len(grid) {
		return grid, idGrid
	}
	if xpos == len(grid[ypos]) {
		grid, idGrid = findTiling(tiles, usedTiles, grid, idGrid, ypos+10, 0)
		return grid, idGrid
	}

	for _, t := range tiles {

		if tileInSlice(t, usedTiles) {
			continue
		}

		for i := 0; i < len(t.variants); i++ {

			//			fmt.Println("Trying to fit TileID:", t.ID, ", variant:", i)
			//			fmt.Printf("(y, x): (%v, %v)\n", ypos, xpos)

			variant := t.variants[i]
			borders := t.borders[i] // {top, right, bottom, left}

			//			util.PrintGrid(variant)
			//			fmt.Println("Top border:", borders[0])
			//			fmt.Println("Left border:", borders[3])
			//			fmt.Println()

			fits := true

			if ypos > 0 {
				bottomOfAboveTile := getBottomOfAboveTile(grid, ypos, xpos)
				fits = fits && bottomOfAboveTile == borders[0]
			}
			if xpos > 0 {
				rightSideOfLeftTile := getRightSideOfLeftTile(grid, ypos, xpos)
				fits = fits && rightSideOfLeftTile == borders[3]
			}
			if fits {
				//				fmt.Println("IT FITS!")
				grid = addTileToGrid(variant, ypos, xpos, grid)
				idGrid = addIDToGrid(t.ID, ypos, xpos, idGrid)
				usedTiles = append(usedTiles, t)
				grid, idGrid = findTiling(tiles, usedTiles, grid, idGrid, ypos, xpos+10)
				usedTiles = removeTileFromSlice(t, usedTiles)
				for y := ypos; y < ypos+10; y++ {
					for x := xpos; x < xpos+10; x++ {
						grid[y][x] = "x"
					}
				}
				idGrid[ypos/10][xpos/10] = -1
			}
		}
	}
}

func getBottomOfAboveTile(grid [][]string, ypos, xpos int) string {
	bottom := []string{}
	for x := xpos; x < xpos+10; x++ {
		bottom = append(bottom, grid[ypos-1][x])
	}
	return strings.Join(bottom, "")
}

func getRightSideOfLeftTile(grid [][]string, ypos, xpos int) string {
	right := []string{}
	for y := ypos; y < ypos+10; y++ {
		right = append(right, grid[y][xpos-1])
	}
	return strings.Join(right, "")
}

func addTileToGrid(tile [][]string, ypos, xpos int, grid [][]string) [][]string {
	for y := ypos; y < ypos+len(tile); y++ {
		for x := xpos; x < xpos+len(tile); x++ {
			grid[y][x] = tile[y-ypos][x-xpos]
		}
	}
	return grid
}

func addIDToGrid(id int, ypos, xpos int, idGrid [][]int) [][]int {
	idGrid[ypos/10][xpos/10] = id
	return idGrid
}

func removeTileFromSlice(tile Tile, slice []Tile) []Tile {
	for i, t := range slice {
		if t.ID == tile.ID {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func getTiles(input []string) []Tile {
	tiles := []Tile{}
	for _, idAndTile := range input {
		tileID := util.GetIntsAsInts(idAndTile)[0]
		grid, _ := util.ParseInputByLineAndRune(strings.Join(strings.Split(idAndTile, "\r\n")[1:], "\n"))
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
		right += grid[y][len(grid)-1]
	}

	bottom := strings.Join(grid[len(grid)-1], "")

	left := ""
	for y := 0; y < len(grid); y++ {
		left += grid[y][0]
	}

	return []string{top, right, bottom, left}
}

func tileInSlice(tile Tile, slice []Tile) bool {
	for _, t := range slice {
		if t.ID == tile.ID {
			return true
		}
	}
	return false
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
