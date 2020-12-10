package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func getJoltages(input []int) []int {
	joltages := []int{0}
	for _, a := range input {
		joltages = append(joltages, a)
	}
	joltages = append(joltages, joltages[len(joltages)-1]+3)
	return joltages
}

func findDiffs(adapters []int) int {

	joltages := getJoltages(adapters)

	diffs := make(map[int]int)
	for i := 1; i < len(joltages); i++ {
		diffs[joltages[i]-joltages[i-1]]++
	}

	return diffs[1] * diffs[3]
}

func part1(input []int) int {
	sort.Ints(input)
	return findDiffs(input)
}

func part2(input []int) int {

	sort.Ints(input)

	joltages := getJoltages(input)

	cache := map[int]int{0: 1}
	for i := 1; i < len(joltages); i++ {
		for j := 0; j < i; j++ {
			diff := joltages[i] - joltages[j]
			if 1 <= diff && diff <= 3 { // not needed if joltages are unique, which they are
				cache[i] += cache[j]
			}
		}
	}

	return cache[len(joltages)-1]

}

func main() {

	rawInput := util.FetchInputForDay("2020", "10")
	parsedInput := util.ParseInputBySepToInts(rawInput, '\n')
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("Answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Answer retrieved in: ", e)

}
