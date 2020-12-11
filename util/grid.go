package util

import (
	"fmt"
)

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

// PrintGrid prints a grid
func PrintGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}
