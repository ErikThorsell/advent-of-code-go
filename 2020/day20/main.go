package main

import (
	"fmt"
	"math"
	"os"
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
	squareSize := int(math.Sqrt(float64(len(tiles))))

	idGrid := make([][]string, squareSize)
	for i := 0; i < squareSize; i++ {
		idGrid[i] = make([]string, squareSize)
	}

	variantGrid := make([][]string, squareSize)
	for i := 0; i < squareSize; i++ {
		variantGrid[i] = make([]string, squareSize)
	}

	grid := make([][]string, 10*squareSize)
	for g := 0; g < len(grid); g++ {
		grid[g] = make([]string, len(grid))
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x] = "x"
		}
	}

	usedTiles := []Tile{}
	grid, idGrid = findTiling(tiles, usedTiles, grid, idGrid, variantGrid, 0, 0)

	return -1
}

func p1ans(grid [][]string) int {
	xs := []string{grid[0][0], grid[0][len(grid)-1], grid[len(grid)-1][0], grid[len(grid)-1][len(grid)-1]}
	ans := 1
	for _, s := range xs {
		ans *= util.ToInt(strings.TrimSpace(s))
	}
	return ans
}

func findTiling(tiles []Tile, usedTiles []Tile, grid, idGrid, variantGrid [][]string, ypos, xpos int) ([][]string, [][]string) {

	if len(usedTiles) == len(tiles) {
		fmt.Println("Answer for Part 1:", p1ans(idGrid))
		gwb := gridWithoutBorders(tiles, idGrid, variantGrid)
		wr := findWaterRoughness(gwb)
		fmt.Println("Answer for Part 2:", wr)
		os.Exit(-1)
	}

	for _, t := range tiles {

		if tileInSlice(t, usedTiles) {
			continue
		}

		//		fmt.Printf("Trying to add Tile %v at (%v, %v)\n", t.ID, ypos, xpos)

		for i := 0; i < len(t.variants); i++ {

			variant := t.variants[i]
			borders := t.borders[i] // {top, right, bottom, left}

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
				//				fmt.Printf("Tile %v, Variant %v FITS!\n", t.ID, i)
				grid = addTileToGrid(variant, ypos, xpos, grid)
				idGrid = addIDToGrid(fmt.Sprint(t.ID)+" ", ypos, xpos, idGrid)
				variantGrid = addVariantToGrid(fmt.Sprint(i)+" ", ypos, xpos, variantGrid)
				//				util.PrintGrid(grid)
				//				util.PrintGrid(idGrid)
				usedTiles = append(usedTiles, t)

				if xpos == len(grid[ypos])-10 {
					//					fmt.Println("Trying to add next Tile on next row")
					//					fmt.Println()
					findTiling(tiles, usedTiles, grid, idGrid, variantGrid, ypos+10, 0)
				} else {
					//					fmt.Println("Trying to add next Tile on same row")
					//					fmt.Println()
					findTiling(tiles, usedTiles, grid, idGrid, variantGrid, ypos, xpos+10)
				}

				usedTiles = removeTileFromSlice(t, usedTiles)
				//				fmt.Printf("Unable to fit ANY tile.\nRemoving %v from the used tiles.\n\n", t.ID)
				grid = clearTile(grid, ypos, xpos)
				idGrid[ypos/10][xpos/10] = "-1"
				variantGrid[ypos/10][xpos/10] = "-1"
			}
		}
		//		fmt.Printf("No variant of Tile %v fit. NEXT!\n", t.ID)
	}
	return grid, idGrid
}

func gridWithoutBorders(tiles []Tile, idGrid, variantGrid [][]string) [][]string {

	squareSize := int(math.Sqrt(float64(len(tiles))))
	tileSize := len(tiles[0].variants[0]) - 2

	grid := make([][]string, tileSize*squareSize)
	for g := 0; g < len(grid); g++ {
		grid[g] = make([]string, len(grid))
	}

	for y := 0; y < len(idGrid); y++ {
		for x := 0; x < len(idGrid[y]); x++ {
			variant := getTileVariant(tiles, util.ToInt(strings.TrimSpace(idGrid[y][x])), util.ToInt(strings.TrimSpace(variantGrid[y][x])))
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
			//			fmt.Printf("Scanning for Sea Monster from (%v,%v)\n", y, x)
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
					//				fmt.Println("Not a sea monster")
					return false, -1
				}
				numHash++
			}
		}
	}
	return true, numHash
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

func addIDToGrid(id string, ypos, xpos int, idGrid [][]string) [][]string {
	idGrid[ypos/10][xpos/10] = id
	return idGrid
}

func addVariantToGrid(id string, ypos, xpos int, variantGrid [][]string) [][]string {
	variantGrid[ypos/10][xpos/10] = id
	return variantGrid
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

	//exampleData, _ := ioutil.ReadFile("2020/day20/example")
	//rawInput := string(exampleData)
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
