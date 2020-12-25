package util

import (
	"fmt"
	"reflect"
	"strings"
)

// PrintGrid prints a grid
func PrintGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		//		if i >= 10 && RealMod(i, 10) == 0 {
		//			fmt.Print("\n")
		//		}
		for j := 0; j < len(grid[i]); j++ {
			//			if j >= 10 && RealMod(j, 10) == 0 {
			//				fmt.Print(" ")
			//			}
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

// CopyGrid returns a deep copy of the grid
func CopyGrid(originalGrid [][]string) [][]string {
	newGrid := make([][]string, len(originalGrid))
	for i, rowInGrid := range originalGrid {
		newGrid[i] = make([]string, len(rowInGrid))
	}

	for i := 0; i < len(originalGrid); i++ {
		for j := 0; j < len(originalGrid[i]); j++ {
			newGrid[i][j] = originalGrid[i][j]
		}
	}

	return newGrid
}

// CountOccurences counts the number of interesting chars in the provided grid
func CountOccurences(grid [][]string, charOfInterest string) int {

	interestingCharacters := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == charOfInterest {
				interestingCharacters++
			}
		}
	}
	return interestingCharacters
}

// CheckAdjacent returns the number of adjacent tiles to (i,j) with interesting characters
func CheckAdjacent(grid [][]string, i int, j int, charOfInterest string) int {
	directionsToLook := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	numCharsOfInterest := 0

	for _, dir := range directionsToLook {

		y := i + dir[0]
		x := j + dir[1]

		if 0 <= y && y < len(grid) && 0 <= x && x < len(grid[i]) {
			if grid[y][x] == charOfInterest {
				numCharsOfInterest++
			}
		}
	}

	return numCharsOfInterest

}

// CheckLineOfSight returns the number of interesting characters that can be seen from (i,j)
func CheckLineOfSight(grid [][]string, i, j int, charOfInterest string) int {

	directionsToLook := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	numIntCharInSight := 0

	for _, dir := range directionsToLook {
		watchDistance := 1

		for {

			y := i + (watchDistance * dir[0])
			x := j + (watchDistance * dir[1])

			if 0 <= y && y < len(grid) && 0 <= x && x < len(grid[i]) {
				if grid[y][x] == "." {
					watchDistance++
				} else if grid[y][x] == "L" {
					break
				} else if grid[y][x] == charOfInterest {
					numIntCharInSight++
					break
				}
			} else {
				break
			}
		}
	}

	return numIntCharInSight
}

// GetGridBorders returns the borders of a grid
func GetGridBorders(grid [][]string) []string {

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

// GetGridFlips returns ass possible flips of a grid
// The flips are:
//   1. The original grid
//   2. Flip around vertical axis
//   3. Flip around horizontal axis
//   4. Flip around both axis
func GetGridFlips(grid [][]string) [][][]string {

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

// GetGridRotations returns as 90 degree rotations of grid
func GetGridRotations(grid [][]string) [][][]string {

	rotations := make([][][]string, 4)

	// Original
	rotations[0] = CopyGrid(grid)

	// Take the previous grid and rotate it 90 degrees
	for i := 1; i <= 3; i++ {
		grid := CopyGrid(rotations[i-1])
		for x := range grid {
			for y := range grid[x] {
				grid[x][y] = rotations[i-1][len(grid[x])-y-1][x]
			}
		}
		rotations[i] = grid
	}
	return rotations
}

// GetGridTranspose returns the grid, transposed
func GetGridTranspose(grid [][]string) [][]string {
	tg := make([][]string, 0)
	for y := 0; y < len(grid); y++ {
		tg = append(grid, make([]string, 0))
		for x := len(grid[y]) - 1; x >= 0; x-- {
			tg[y][x] = grid[x][y]
		}
	}
	return tg
}

// GridInSlice checks whether grid is in the slice
func GridInSlice(grid [][]string, slice [][][]string) bool {
	for _, t := range slice {
		if reflect.DeepEqual(t, grid) {
			return true
		}
	}
	return false
}
