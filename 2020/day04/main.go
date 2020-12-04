package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func isValidPassport(passport []string) bool {

	foundKeys := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}

	for _, k := range passport {
		foundKeys[k] = true
	}

	valid := true
	for _, v := range foundKeys {
		valid = valid && v
	}

	return valid

}

func byrIsValid(byr string) bool {
	v, _ := strconv.Atoi(byr)
	if 1920 <= v && v <= 2002 {
		return true
	}
	return false
}

func iyrIsValid(iyr string) bool {
	v, _ := strconv.Atoi(iyr)
	if 2010 <= v && v <= 2020 {
		return true
	}
	return false
}

func eyrIsValid(eyr string) bool {
	v, _ := strconv.Atoi(eyr)
	if 2020 <= v && v <= 2030 {
		return true
	}
	return false
}

func hgtIsValid(hgt string) bool {

	h := 0
	u := hgt[len(hgt)-2:]

	if u == "cm" {
		h, _ = strconv.Atoi(hgt[0:3])
		if 150 <= h && h <= 193 {
			return true
		}
	}

	if u == "in" {
		h, _ = strconv.Atoi(hgt[0:2])
		if 59 <= h && h <= 76 {
			return true
		}
	}

	return false

}

func hclIsValid(hcl string) bool {

	if string(hcl[0]) != "#" {
		return false
	}
	if len(hcl[1:]) != 6 {
		return false
	}
	// check if hcl[1:] contains other characters than 0-9 and a-f ?
	return true
}

func eclIsValid(ecl string) bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, k := range validColors {
		if ecl == k {
			return true
		}
	}
	return false
}

func pidIsValid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	if _, err := strconv.Atoi(pid); err == nil {
		return true
	}
	return false
}

func isValidPassportStrict(passport map[string]string) bool {

	keys := extractKeysFromMap(passport)
	if !isValidPassport(keys) {
		return false
	}

	if !byrIsValid(passport["byr"]) {
		return false
	}

	if !iyrIsValid(passport["iyr"]) {
		return false
	}

	if !eyrIsValid(passport["eyr"]) {
		return false
	}

	if !hgtIsValid(passport["hgt"]) {
		return false
	}

	if !eclIsValid(passport["ecl"]) {
		return false
	}

	if !pidIsValid(passport["pid"]) {
		return false
	}

	if !hclIsValid(passport["hcl"]) {
		return false
	}

	return true

}

func convertPassportToMap(passport string) map[string]string {

	passportFields := strings.Fields(passport)
	passportAsMap := make(map[string]string)

	for _, x := range passportFields {
		keyValue := strings.Split(x, ":")
		passportAsMap[keyValue[0]] = keyValue[1]
	}

	return passportAsMap

}

func extractKeysFromMap(passport map[string]string) []string {
	var keys []string
	for k := range passport {
		keys = append(keys, k)
	}
	return keys
}

func part1(listOfPassports []string) int {

	validPassports := 0

	for _, p := range listOfPassports {
		passport := convertPassportToMap(p)
		keys := extractKeysFromMap(passport)
		if isValidPassport(keys) {
			validPassports++
		}
	}

	return validPassports

}

func part2(listOfPassports []string) int {

	validPassports := 0

	for _, p := range listOfPassports {
		passport := convertPassportToMap(p)
		if isValidPassportStrict(passport) {
			validPassports++
		}
	}
	return validPassports
}

func main() {

	rawInput := util.FetchInputForDay("2020", "4")
	parsedInput := util.ParseInputByBlankLine(rawInput)

	ans1 := part1(parsedInput)
	fmt.Println("Answer for first question: ", ans1)

	ans2 := part2(parsedInput)
	fmt.Println("Answer for second question: ", ans2)

}
