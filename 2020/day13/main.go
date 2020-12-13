package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
	"github.com/deanveloper/modmath/v1/bigmod"
)

func part1(earliestTime int, busTable []string) int {

	validBusTable := getOperationalBuses(busTable)
	ct := earliestTime

	for {

		for _, busID := range validBusTable {
			if util.RemIsZero(ct, busID) {
				return util.Abs(ct-earliestTime) * busID
			}
		}
		ct++

	}

}

func getOperationalBuses(bustable []string) []int {
	operationalBuses := []int{}
	for _, b := range bustable {
		if util.IsInt(b) {
			operationalBuses = append(operationalBuses, util.ToInt(b))
		}
	}
	return operationalBuses
}

func part2(table []string) int64 {
	constraints := findDepartureConstrains(table)
	var opBusses []bigmod.CrtEntry
	for b, t := range constraints {
		opBusses = append(opBusses, bigmod.CrtEntry{A: big.NewInt(int64(b - t)), N: big.NewInt(int64(b))})
	}
	return bigmod.SolveCrtMany(opBusses).Int64()
}

func findDepartureConstrains(busTable []string) map[int]int {

	constraints := make(map[int]int)
	t := 0

	for _, b := range busTable {
		if util.IsInt(b) {
			constraints[util.ToInt(b)] = t
		}
		t++
	}
	return constraints
}

func main() {

	rawInput := util.FetchInputForDay("2020", "13")
	earliestTime, busTable := util.ParseBusTableInput(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(earliestTime, busTable)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(busTable)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
