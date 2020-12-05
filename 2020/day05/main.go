package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func getRow(bp string) int {

	runes := strings.Split(bp, "")
	min := 0
	max := 127

	for _, r := range runes {
		if max < min {
			log.Fatal("Max is less than min: ", max, min)
		}
		if r == "F" {
			max -= (max-min)/2 + 1
		} else if r == "B" {
			min += (max-min)/2 + 1
		}
	}

	if max == min {
		return max
	}

	log.Fatal("Inconclusive boarding pass!")
	return -1

}

func getCol(bp string) int {

	runes := strings.Split(bp, "")
	min := 0
	max := 7

	for _, r := range runes {
		if max < min {
			log.Fatal("Max is less than min: ", max, min)
		}
		if r == "L" {
			max -= (max-min)/2 + 1
		} else if r == "R" {
			min += (max-min)/2 + 1
		}
	}

	if max == min {
		return max
	}

	log.Fatal("Inconclusive boarding pass!")
	return -1
}

func getSeatID(bp string) int {
	return getRow(bp)*8 + getCol(bp)
}

func part1(boardingPasses []string) int {

	var seatIDs []int

	for _, row := range boardingPasses {
		seatIDs = append(seatIDs, getSeatID(row))
	}

	sort.Ints(seatIDs)

	return seatIDs[len(seatIDs)-1]

}

func generateSeats(min int, max int) []int {
	var seats []int
	for s := min; s <= max; s++ {
		seats = append(seats, s)
	}
	return seats
}

func findMySeat(allSeats []int, plane []int) int {
	for i := range allSeats {
		if allSeats[i] != plane[i] {
			return allSeats[i]
		}
	}
	return -1
}

func part2(boardingPasses []string) int {

	var seatIDs []int

	for _, row := range boardingPasses {
		seatIDs = append(seatIDs, getSeatID(row))
	}

	sort.Ints(seatIDs)

	return findMySeat(generateSeats(seatIDs[0], seatIDs[len(seatIDs)-1]), seatIDs)
}

func main() {

	rawInput := util.FetchInputForDay("2020", "5")
	parsedInput := util.ParseInputByLine(rawInput)

	ans1 := part1(parsedInput)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(parsedInput)
	fmt.Println("Answer for second question: ", ans2)

}
