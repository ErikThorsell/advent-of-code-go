package main

import (
	"fmt"
	"math"
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

type idCell struct {
	ID      int
	variant int
}

func part1(input []string) (int, []Tile, [][]idCell) {

	tiles := getTiles(input)
	tileSize := len(tiles[0].variants[0]) // assuming tiles are squares
	squareSize := int(math.Sqrt(float64(len(tiles))))

	tilingGrid := make([][]string, tileSize*squareSize)
	for row := 0; row < len(tilingGrid); row++ {
		tilingGrid[row] = make([]string, len(tilingGrid))
	}

	idGrid := make([][]idCell, squareSize)
	for row := 0; row < squareSize; row++ {
		idGrid[row] = make([]idCell, len(idGrid))
	}

	usedTiles := []Tile{}
	tilingGrid, idGrid, _ = findTiling(tiles, usedTiles, tilingGrid, idGrid, 0, 0)

	return sumCorners(idGrid), tiles, idGrid
}

func sumCorners(grid [][]idCell) int {
	xs := []int{grid[0][0].ID, grid[0][len(grid)-1].ID, grid[len(grid)-1][0].ID, grid[len(grid)-1][len(grid)-1].ID}
	ans := 1
	for _, s := range xs {
		ans *= s
	}
	return ans
}

func findTiling(tiles []Tile, usedTiles []Tile, grid [][]string, idGrid [][]idCell, ypos, xpos int) ([][]string, [][]idCell, []Tile) {

	for _, t := range tiles {

		if tileInSlice(t, usedTiles) {
			continue
		}

		for variantNumber := 0; variantNumber < len(t.variants); variantNumber++ {

			variant := t.variants[variantNumber]
			borders := t.borders[variantNumber] // {top, right, bottom, left}

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
				grid = addTileToGrid(variant, ypos, xpos, grid)
				idGrid = addIDToGrid(t.ID, variantNumber, ypos, xpos, idGrid)
				usedTiles = append(usedTiles, t)

				if xpos == len(grid[ypos])-10 {
					grid, idGrid, usedTiles = findTiling(tiles, usedTiles, grid, idGrid, ypos+10, 0)
				} else {
					grid, idGrid, usedTiles = findTiling(tiles, usedTiles, grid, idGrid, ypos, xpos+10)
				}

				if len(usedTiles) == len(tiles) {
					return grid, idGrid, usedTiles
				}

				usedTiles = removeTileFromSlice(t, usedTiles)
				grid = clearTile(grid, ypos, xpos)
				idGrid[ypos/10][xpos/10] = idCell{}
			}
		}
	}
	return grid, idGrid, usedTiles
}

func clearTile(grid [][]string, ypos, xpos int) [][]string {
	for y := ypos; y < ypos+10; y++ {
		for x := xpos; x < xpos+10; x++ {
			grid[y][x] = "x"
		}
	}
	return grid
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

func addIDToGrid(id, variant, ypos, xpos int, idGrid [][]idCell) [][]idCell {
	idGrid[ypos/10][xpos/10] = idCell{ID: id, variant: variant}
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
		grid, _ := util.ParseInputByLineAndRune(strings.Join(strings.Split(idAndTile, "\n")[1:], "\n"))
		variants := getAllGridVariants(grid)
		borders := make(map[int][]string)
		for i, v := range variants {
			borders[i] = util.GetGridBorders(v)
		}
		tiles = append(tiles, Tile{ID: tileID, variants: variants, borders: borders})
	}
	return tiles
}

func getAllGridVariants(grid [][]string) map[int][][]string {

	flippedGrids := util.GetGridFlips(grid)
	rotatedAndFlippedGrids := make([][][]string, 0)
	for _, mt := range flippedGrids {
		candidateGrids := util.GetGridRotations(mt)
		for _, ct := range candidateGrids {
			if !util.GridInSlice(ct, rotatedAndFlippedGrids) {
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

func tileInSlice(tile Tile, slice []Tile) bool {
	for _, t := range slice {
		if t.ID == tile.ID {
			return true
		}
	}
	return false
}

func part2(tiles []Tile, idGrid [][]idCell) int {
	gwb := gridWithoutBorders(tiles, idGrid)
	return findWaterRoughness(gwb)
}

func gridWithoutBorders(tiles []Tile, idGrid [][]idCell) [][]string {

	squareSize := int(math.Sqrt(float64(len(tiles))))
	tileSize := len(tiles[0].variants[0]) - 2

	grid := make([][]string, tileSize*squareSize)
	for g := 0; g < len(grid); g++ {
		grid[g] = make([]string, len(grid))
	}

	for y := 0; y < len(idGrid); y++ {
		for x := 0; x < len(idGrid[y]); x++ {
			variant := getTileVariant(tiles, idGrid[y][x].ID, idGrid[y][x].variant)
			variant = tileWithoutBorders(variant)
			grid = addTileToGrid(variant, y*len(variant), x*len(variant[0]), grid)
		}
	}

	return grid
}

func getTileVariant(tiles []Tile, id int, variant int) [][]string {
	for _, t := range tiles {
		if t.ID == id {
			return t.variants[variant]
		}
	}
	return [][]string{}
}

func tileWithoutBorders(tile [][]string) [][]string {

	twb := make([][]string, len(tile)-2)
	for i := 0; i < len(tile)-2; i++ {
		twb[i] = make([]string, len(tile[i])-2)
	}

	for y := 0; y < len(tile); y++ {
		for x := 0; x < len(tile[y]); x++ {
			if y != 0 && y != len(tile)-1 && x != 0 && x != len(tile[y])-1 {
				twb[y-1][x-1] = tile[y][x]
			}
		}
	}

	return twb
}

func findWaterRoughness(grid [][]string) int {

	rawSeaMonster := `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

	seaMonster, _ := util.ParseInputByLineAndRune(rawSeaMonster)
	hashesInGrid := util.CountOccurences(grid, "#")
	coveredHashes := countCoveredHashes(grid, seaMonster)

	return hashesInGrid - coveredHashes

}

func countCoveredHashes(grid [][]string, seaMonster [][]string) int {

	gridVariants := getAllGridVariants(grid)

	for _, gv := range gridVariants {
		if found, hashes := findSeaMonster(gv, seaMonster); found {
			return hashes
		}
	}
	fmt.Println("No Sea Monsters found!")
	return -1
}

func findSeaMonster(grid [][]string, seaMonster [][]string) (bool, int) {

	seaMonsterHashes := 0
	for y := 0; y+len(seaMonster) < len(grid); y++ {
		for x := 0; x+len(seaMonster[0]) < len(grid[y]); x++ {
			if found, hashes := scanArea(grid, y, x, seaMonster); found {
				seaMonsterHashes += hashes
			}
		}
	}
	if seaMonsterHashes > 0 {
		return true, seaMonsterHashes
	}
	return false, -1
}

func scanArea(grid [][]string, ypos, xpos int, seaMonster [][]string) (bool, int) {
	numHash := 0
	for y := 0; y < len(seaMonster); y++ {
		for x := 0; x < len(seaMonster[y]); x++ {
			if seaMonster[y][x] == "#" {
				if grid[ypos+y][xpos+x] != "#" {
					return false, -1
				}
				numHash++
			}
		}
	}
	return true, numHash
}

func main() {

	rawInput := util.FetchInputForDay("2020", "20")
	parsedInput := util.ParseInputByBlankLine(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1, tiles, idGrid := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(tiles, idGrid)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
