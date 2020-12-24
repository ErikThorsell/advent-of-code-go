package main

import (
	"fmt"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

type hexTile struct {
	x int
	y int
	z int
}

func part1(input []string) (int, map[hexTile]int) {

	flipInstr := make([][]string, 0)
	for _, row := range input {
		flipInstr = append(flipInstr, parseInstruction(row))
	}

	tiles := make(map[hexTile]int)
	for _, instr := range flipInstr {
		x, y, z := move(instr)
		tiles[hexTile{x, y, z}]++
	}

	return countBlackTiles(tiles), tiles
}

func parseInstruction(instruction string) []string {
	parsedInstructions := []string{}
	for i := 0; i < len(instruction); {
		if string(instruction[i]) == "n" || string(instruction[i]) == "s" {
			parsedInstructions = append(parsedInstructions, instruction[i:i+2])
			i += 2
			continue
		} else {
			parsedInstructions = append(parsedInstructions, string(instruction[i]))
			i++
			continue
		}
	}
	return parsedInstructions
}

func move(instr []string) (int, int, int) {
	x, y, z := 0, 0, 0
	for _, i := range instr {
		switch i {
		case "ne":
			x++
			z--
		case "e":
			x++
			y--
		case "se":
			z++
			y--
		case "sw":
			x--
			z++
		case "w":
			x--
			y++
		case "nw":
			z--
			y++
		}
	}
	return x, y, z
}

func countBlackTiles(tiles map[hexTile]int) int {
	bt := 0
	for _, f := range tiles {
		if f%2 != 0 {
			bt++
		}
	}
	return bt
}

func part2(tiles map[hexTile]int) int {

	// Ensure all "original tiles" have neighbors
	tiles = padTiles(tiles)

	for i := 0; i < 100; i++ {

		toFlip := make(map[hexTile]bool)

		for tile, numberOfFlips := range tiles {

			adjBlackTiles, unseenNeighbors := exploreAdjacentTiles(tile, tiles)

			if numberOfFlips%2 == 0 {
				if len(adjBlackTiles) == 2 {
					toFlip[tile] = true
				}
			} else {
				if len(adjBlackTiles) == 0 || len(adjBlackTiles) > 2 {
					toFlip[tile] = true
				}
			}

			// ensure all tiles we looked at are taken into consideration next round
			for vk, vv := range unseenNeighbors {
				tiles[vk] = vv
			}

		}

		for t := range toFlip {
			tiles[t]++
		}

	}

	return countBlackTiles(tiles)

}

func padTiles(tiles map[hexTile]int) map[hexTile]int {
	for t := range tiles {
		adjacentTiles := getAdjacentTiles(t)
		for _, t := range adjacentTiles {
			_, tileExists := tiles[t]
			if tileExists {
				continue
			} else {
				tiles[t] = 0
			}
		}
	}
	return tiles
}

func getAdjacentTiles(tile hexTile) []hexTile {
	adjacentTiles := []hexTile{}
	adjacentTiles = append(adjacentTiles, hexTile{x: tile.x + 1, y: tile.y, z: tile.z - 1})
	adjacentTiles = append(adjacentTiles, hexTile{x: tile.x + 1, y: tile.y - 1, z: tile.z})
	adjacentTiles = append(adjacentTiles, hexTile{x: tile.x, y: tile.y - 1, z: tile.z + 1})
	adjacentTiles = append(adjacentTiles, hexTile{x: tile.x - 1, y: tile.y, z: tile.z + 1})
	adjacentTiles = append(adjacentTiles, hexTile{x: tile.x - 1, y: tile.y + 1, z: tile.z})
	adjacentTiles = append(adjacentTiles, hexTile{x: tile.x, y: tile.y + 1, z: tile.z - 1})
	return adjacentTiles
}

func exploreAdjacentTiles(tile hexTile, tiles map[hexTile]int) ([]hexTile, map[hexTile]int) {

	adjacentTiles := getAdjacentTiles(tile)
	blackAdjacentTiles := []hexTile{}
	tilesToAdd := make(map[hexTile]int)

	for _, t := range adjacentTiles {
		flips, tileExists := tiles[t]
		if tileExists {
			if flips%2 != 0 {
				blackAdjacentTiles = append(blackAdjacentTiles, t)
			}
		} else {
			tilesToAdd[t] = 0
		}
	}
	return blackAdjacentTiles, tilesToAdd
}

func main() {

	rawInput := util.FetchInputForDay("2020", "24")
	parsedInput := util.ParseInputByLine(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1, tiles := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(tiles)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
