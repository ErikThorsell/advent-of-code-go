package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func getAllBags(input []string) []string {

	var bags []string
	for _, i := range input {
		bags = append(bags, strings.TrimSpace(strings.Split(i, "bags contain")[0]))
	}
	return bags
}

func parseBagRules(rules []string) map[string][]string {

	bagRules := make(map[string][]string)

	for _, r := range rules {

		bagAndCond := strings.Split(r, "bags contain")
		bag := strings.TrimSpace(bagAndCond[0])
		conds := parseConditions(bagAndCond[1])
		bagRules[bag] = conds
	}

	return bagRules

}

func parseBagRulesNoNum(rules []string) map[string][]string {

	bagRules := make(map[string][]string)

	for _, r := range rules {

		bagAndCond := strings.Split(r, "bags contain")
		bag := strings.TrimSpace(bagAndCond[0])
		conds := parseConditionsNoNum(bagAndCond[1])
		bagRules[bag] = conds
	}

	return bagRules

}

func parseConditions(rawConditions string) []string {

	conds := strings.Split(rawConditions, ",")

	var conditions []string
	for _, c := range conds {
		splitC := strings.Fields(c)
		trimC := splitC[0 : len(splitC)-1]
		joinC := strings.TrimSpace(strings.Join(trimC, " "))
		conditions = append(conditions, joinC)
	}
	return conditions
}

func parseConditionsNoNum(rawConditions string) []string {

	parsedConditions := parseConditions(rawConditions)

	var noNum []string
	for _, c := range parsedConditions {
		var pc string
		if c == "no other" {
			pc = c
		} else {
			pc = strings.TrimSpace(c[1:])
		}
		noNum = append(noNum, pc)
	}

	return noNum

}

func findBagHierarchies(startBag string, bagRules map[string][]string) []string {

	bagsToOpen := []string{startBag}
	openedBags := []string{}

	for {

		if len(bagsToOpen) == 0 {
			return openedBags[1:]
		}

		currentBag := bagsToOpen[0]
		bagsToOpen = bagsToOpen[1:]

		for _, v := range bagRules[currentBag] {
			if v != "no other" {
				bagsToOpen = append(bagsToOpen, v)
			}
		}

		openedBags = append(openedBags, currentBag)

	}

}

func getValidBags(hierarchies map[string][]string) []string {
	var validBags []string
	for k, v := range hierarchies {
		if util.StringInSlice("shiny gold", v) {
			validBags = append(validBags, k)
		}
	}
	return validBags
}

func remNum(verbBag string) string {
	return strings.Join(strings.Fields(verbBag)[1:], " ")
}

func countBagsInBag(startBag string, bagRules map[string][]string) int {
	countBags := 1
	for _, bag := range bagRules[startBag] {
		if bag != "no other" {
			numOfBags := util.ToInt(strings.Fields(bag)[0])
			countBags += numOfBags * countBagsInBag(remNum(bag), bagRules)
		}
	}
	return countBags
}

func part1(input []string) int {

	bags := getAllBags(input)
	bagRules := parseBagRulesNoNum(input)

	hierarchies := make(map[string][]string)
	for _, bag := range bags {
		hierarchies[bag] = findBagHierarchies(bag, bagRules)
	}

	validBags := getValidBags(hierarchies)
	return len(validBags)

}

func part2(input []string) int {

	bagRules := parseBagRules(input)
	return countBagsInBag("shiny gold", bagRules) - 1

}

func main() {

	rawInput := util.FetchInputForDay("2020", "7")
	parsedInput := util.ParseInputByLine(rawInput)
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
