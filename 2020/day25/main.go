package main

import (
	"fmt"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string) int {

	publicKeys := make(map[string]int)
	publicKeys["card"] = util.ToInt(input[0])
	publicKeys["door"] = util.ToInt(input[1])

	subjectNumber := 7
	loopSizes := make(map[string]int)
	for kName, kValue := range publicKeys {

		value := 1
		loopSize := 0

		for {

			if value == kValue {
				loopSizes[kName] = loopSize
				break
			}

			value *= subjectNumber
			value = value % 20201227

			loopSize++

		}
	}

	privateKeys := make(map[string]int)
	privateKeys["card"] = encrypt(publicKeys["door"], loopSizes["card"])
	privateKeys["door"] = encrypt(publicKeys["card"], loopSizes["door"])

	if privateKeys["card"] != privateKeys["door"] {
		fmt.Println("Different Private keys! Something is wrong.")
		return -1
	}

	return privateKeys["card"]
}

func encrypt(subjectNumber int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value = value % 20201227
	}
	return value
}

func part2(input []string) int {
	return 0
}

func main() {

	rawInput := util.FetchInputForDay("2020", "25")
	parsedInput := util.ParseInputByLine(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	//	fmt.Println()

	//	s = time.Now()
	//	ans2 := part2(parsedInput)
	//	t = time.Now()
	//	e = t.Sub(s)
	//	fmt.Println("Answer for second question: ", ans2)
	//	fmt.Println("Second answer retrieved in: ", e)

}
