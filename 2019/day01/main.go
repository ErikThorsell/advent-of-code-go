package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func getFuel(mass int) int {
	return (mass / 3) - 2
}

func getFuelForFuel(mass int) int {

	fuelForMass := getFuel(mass)

	fuelForFuel := 0
	currentFuel := fuelForMass

	for {

		currentFuel = getFuel(currentFuel)

		if currentFuel <= 0 {
			break
		}

		fuelForFuel += currentFuel

	}

	return fuelForMass + fuelForFuel

}

func part1(modules []string) int {

	totalFuel := 0

	for _, module := range modules {
		i, _ := strconv.Atoi(module)
		totalFuel += getFuel(i)
	}

	return totalFuel
}

func part2(modules []string) int {

	totalFuel := 0

	for _, module := range modules {
		i, _ := strconv.Atoi(module)
		totalFuel += getFuelForFuel(i)
	}

	return totalFuel
}

func main() {

	input := util.FetchInputForDay("2019", "1")
	data := strings.Fields(input)

	ans1 := part1(data)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(data)
	fmt.Println("Answer for second question: ", ans2)

}
