package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func uniqueRunes(s string) []rune {

	m := make(map[rune]bool)
	for _, r := range s {
		if !unicode.IsSpace(r) {
			m[r] = true
		}
	}

	var uniqueRunes []rune
	for k, v := range m {
		if v {
			uniqueRunes = append(uniqueRunes, k)
		}
	}

	return uniqueRunes

}

func part1(groups []string) int {

	uniqueAnswers := 0
	for _, g := range groups {
		uniqueAnswers += len(uniqueRunes(g))
	}

	return uniqueAnswers
}

func unisonRunes(s string) []rune {

	pplInGroup := strings.Count(s, "\n") + 1

	m := make(map[rune]int)
	for _, r := range s {
		if !unicode.IsSpace(r) {
			m[r]++
		}
	}

	var unisonRunes []rune
	for k, v := range m {
		if v == pplInGroup {
			unisonRunes = append(unisonRunes, k)
		}
	}

	return unisonRunes

}

func part2(groups []string) int {

	unisonAnswers := 0
	for _, g := range groups {
		unisonAnswers += len(unisonRunes(g))
	}

	return unisonAnswers
}

func main() {

	rawInput := util.FetchInputForDay("2020", "6")
	parsedInput := util.ParseInputByBlankLine(rawInput)
	fmt.Println("Done parsing input. Time to start solving!")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(parsedInput)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("Answer retrieved in: ", e)
	fmt.Println()

	// PART 2
	s = time.Now()
	ans2 := part2(parsedInput)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Answer retrieved in: ", e)

}
