package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func isValid(password string) bool {

	digits := strings.Split(password, "")

	foundAdjacent := false

	if digits[0] == digits[1] ||
		digits[1] == digits[2] ||
		digits[2] == digits[3] ||
		digits[3] == digits[4] ||
		digits[4] == digits[5] {
		foundAdjacent = true
	}

	if !sort.StringsAreSorted(digits) {
		return false
	}

	return foundAdjacent

}

func isValidDeluxe(password string) bool {

	digits := strings.Split(password, "")

	if !sort.StringsAreSorted(digits) {
		return false
	}

	for _, d := range digits {
		if strings.Count(password, d) == 2 {
			return true
		}
	}
	return false
}

func part1(lowInt, highInt int) int {

	var validPasswords []int

	for i := lowInt; i <= highInt; i++ {
		if isValid(strconv.Itoa(i)) {
			validPasswords = append(validPasswords, i)
		}

	}

	return len(validPasswords)
}

func part2(lowInt, highInt int) int {

	var validPasswords []int

	for i := lowInt; i <= highInt; i++ {
		if isValidDeluxe(strconv.Itoa(i)) {
			validPasswords = append(validPasswords, i)
		}

	}

	return len(validPasswords)
}

func main() {

	rawInput := util.FetchInputForDay("2019", "4")
	lowInt, highInt := util.ParseNumberDashNumber(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(lowInt, highInt)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("Answer retrieved in: ", e)
	fmt.Println()

	// PART 2
	s = time.Now()
	ans2 := part2(lowInt, highInt)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Answer retrieved in: ", e)

}
