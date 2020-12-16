package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func getCookie(cookiePath string) string {

	data, err := ioutil.ReadFile(cookiePath)

	if err != nil {
		log.Fatal("Unable to read file: ", cookiePath)
	}

	return string(data)
}

// Check if the file exists
func fileExists(path string) bool {

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		fmt.Println("Unable to find file: ", path)
		return false
	}

	fmt.Println("File:", path, "already exists. Using it.")
	return true
}

// FetchInputForDay returns the data for the corresponding challenge
func FetchInputForDay(year string, day string) string {

	fmt.Println("🌟 Fetching today's input! 🌟")
	possibleFile := fmt.Sprintf("%v/data/%v", year, day)
	if fileExists(possibleFile) {
		data, _ := ioutil.ReadFile(possibleFile)
		return string(data)
	}

	fmt.Println(fmt.Sprintf("Fetching data for year: %v, and day: %v.", year, day))
	queryURL := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	req, err := http.NewRequest("GET", queryURL, nil)

	sessionCookie := getCookie("./cookie")
	cookie := http.Cookie{Name: "session", Value: sessionCookie}
	req.AddCookie(&cookie)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("HTTP Request failed: ", err)
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		log.Fatal("HTTP Error: ", resp.StatusCode)
	}

	// The 'ol []byte to string (for trimming) to [byte]-aroo
	data, _ := ioutil.ReadAll(resp.Body)
	data = []byte(strings.TrimSpace(string(data)))

	// Write file to avoid re-fetching
	err = ioutil.WriteFile(possibleFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input retrieved, happy coding!")
	fmt.Println(string(data))
	return string(data)
}

// ParseInputByLineAndSep takes a string and a sep as input.
// Returns a correctly parsed string of strings.
func ParseInputByLineAndSep(input string, sep rune) [][]string {

	listOfStrings := ParseInputByLine(input)

	var listOfItems [][]string
	for _, row := range listOfStrings {
		listOfItems = append(listOfItems, strings.Split(row, string(sep)))
	}

	return listOfItems
}

// ParseInputBySepToInts takes a string and a sep as input.
// Returns a correctly parsed string of ints.
func ParseInputBySepToInts(input string, sep rune) []int {

	listOfStrings := strings.Split(input, string([]rune{sep}))

	var listOfInts []int
	for _, x := range listOfStrings {
		v := ToInt(x)
		listOfInts = append(listOfInts, v)
	}
	return listOfInts
}

// ParseInputByLine takes the raw input, split the input on \n,
// and return a []string.
func ParseInputByLine(input string) []string {
	listOfStrings := strings.Split(input, "\n")
	return listOfStrings
}

// ParseInputByLineAndRune is used to parse a string first by row and then
// split that line into individual runes.
func ParseInputByLineAndRune(input string) ([][]string, int) {
	listOfStrings := strings.Split(input, "\n")
	numberOfRows := len(listOfStrings)

	var doubleList [][]string
	for _, row := range listOfStrings {
		doubleList = append(doubleList, strings.Split(row, ""))
	}

	return doubleList, numberOfRows
}

// ParseInputByBlankLine returns a list of strings where each item
// was originally separated by a blank line.
func ParseInputByBlankLine(input string) []string {
	return strings.Split(input, "\n\n")
}

// ParseNumberDashNumber takes: "number1-number2" and returns the numbers as ints
func ParseNumberDashNumber(input string) (int, int) {
	listOfStrings := strings.Split(input, "-")
	return ToInt(listOfStrings[0]), ToInt(listOfStrings[1])
}

// ParseBusTableInput parses a file that looks like this:
// 939
// 7,13,x,x,59,x,31,19
// https://adventofcode.com/2020/day/13
func ParseBusTableInput(input string) (int, []string) {
	twoStrings := ParseInputByLine(input)
	timestamp := ToInt(twoStrings[0])
	busTable := ParseInputByLineAndSep(twoStrings[1], ',')
	return timestamp, busTable[0]
}

// ParseTicketInput parses the input for Day 16, 2020
func ParseTicketInput(input []string) (map[string][][]int, []int, [][]int) {

	// Create a tickets map
	ticketConstraints := make(map[string][][]int)
	for _, row := range input {
		if row == "" {
			break
		}
		name := strings.TrimSpace(strings.Split(row, ":")[0])
		re := regexp.MustCompile("[0-9]+")
		ranges := re.FindAllString(row, -1)
		ticketConstraints[name] = [][]int{
			{ToInt(ranges[0]), ToInt(ranges[1])},
			{ToInt(ranges[2]), ToInt(ranges[3])},
		}
	}

	// Find my own ticket
	myTicket := []int{}
	for i, row := range input {
		if strings.Contains(row, "your ticket") {
			myRow := input[i+1]
			for _, v := range strings.Split(myRow, ",") {
				myTicket = append(myTicket, ToInt(v))
			}
		}
	}

	// Find nearby tickets
	collect := false
	otherTickets := [][]int{}
	for _, row := range input {
		if strings.Contains(row, "nearby tickets") {
			collect = true
			continue
		}
		if collect {
			ticket := []int{}
			for _, v := range strings.Split(row, ",") {
				ticket = append(ticket, ToInt(v))
			}
			otherTickets = append(otherTickets, ticket)
		}
	}

	return ticketConstraints, myTicket, otherTickets

}
