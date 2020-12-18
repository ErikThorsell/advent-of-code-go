package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string) int {

	res := 0

	precedenceMap := make(map[string]int)
	precedenceMap["+"] = 0
	precedenceMap["*"] = 0

	for _, row := range input {
		elems := strings.Split(strings.ReplaceAll(row, " ", ""), "")
		r := util.ShuntingYard(elems, precedenceMap)
		res += r
	}

	return res

}

func part2(input []string) int {

	res := 0

	precedenceMap := make(map[string]int)
	precedenceMap["+"] = 1
	precedenceMap["*"] = 0

	for _, row := range input {
		elems := strings.Split(strings.ReplaceAll(row, " ", ""), "")
		r := util.ShuntingYard(elems, precedenceMap)
		res += r
	}

	return res

}

func main() {

	rawInput := util.FetchInputForDay("2020", "18")
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
	fmt.Println()

	s = time.Now()
	ans2 := part2(parsedInput)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
