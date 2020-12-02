package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func validatePassword(low int, high int, letter string, password string) bool {

	numberOf := strings.Count(password, letter)

	if low <= numberOf && numberOf <= high {
		return true
	}

	return false

}

func validatePassword2(low int, high int, letter string, password string) bool {

	firstMatch := letter == string(password[low-1])
	secondMatch := letter == string(password[high-1])

	if firstMatch != secondMatch {
		return true
	}
	return false

}

func parsePasswordRow(row string) (int, int, string, string) {

	splitRow := strings.Fields(row)

	lowHigh := strings.Split(splitRow[0], "-")
	low, _ := strconv.Atoi(lowHigh[0])
	high, _ := strconv.Atoi(lowHigh[1])

	char := strings.Trim(splitRow[1], ":")

	password := splitRow[2]

	return low, high, char, password

}

func part1(input []string) int {

	validPasswords := 0

	for _, v := range input {

		low, high, char, password := parsePasswordRow(v)

		if validatePassword(low, high, char, password) {
			validPasswords++
		}

	}

	return validPasswords

}

func part2(input []string) int {

	validPasswords := 0

	for _, v := range input {

		low, high, char, password := parsePasswordRow(v)

		if validatePassword2(low, high, char, password) {
			validPasswords++
		}

	}

	return validPasswords

}

func main() {

	rawInput := util.FetchInputForDay("2020", "2")
	parsedInput := util.ParseInputToListOfStrings(rawInput)

	ans1 := part1(parsedInput)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(parsedInput)
	fmt.Println("Answer for second question: ", ans2)

}
