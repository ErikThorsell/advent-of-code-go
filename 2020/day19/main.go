package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(nodeMap map[string]util.SatelliteNode, messages []string) int {

	validRegExp := getValidRegexp("0", nodeMap)

	re := regexp.MustCompile("^" + validRegExp + "$")
	matchCount := 0
	for _, message := range messages {
		matches := re.MatchString(message)
		if matches {
			matchCount++
		}
	}

	return matchCount

}

func getValidRegexp(nodeID string, nodeMap map[string]util.SatelliteNode) string {

	node := nodeMap[nodeID]

	if node.IsLiteral {
		return node.Rule
	}

	pattern := strings.Builder{}
	pattern.WriteString("(?:")

	split := strings.Split(node.Rule, "|")
	subRulesPatterns := []string{}

	for _, subRule := range split {
		split := strings.Fields(subRule)
		concatRules := []string{}

		for _, newNode := range split {
			toAdd := getValidRegexp(newNode, nodeMap)
			concatRules = append(concatRules, toAdd)
		}

		subRulesPatterns = append(subRulesPatterns, strings.Join(concatRules, ""))

	}

	pattern.WriteString(strings.Join(subRulesPatterns, "|"))
	pattern.WriteRune(')')
	return pattern.String()

}

func part2(nodeMap map[string]util.SatelliteNode, messages []string) int {
	return 0
}

func main() {

	rawInput := util.FetchInputForDay("2020", "19")
	rules, messages := util.ParseSatelliteInput(rawInput)
	fmt.Println("Done parsing input.")
	fmt.Println()

	// PART 1
	s := time.Now()
	ans1 := part1(rules, messages)
	t := time.Now()
	e := t.Sub(s)
	fmt.Println("Answer for first question: ", ans1)
	fmt.Println("First answer retrieved in: ", e)
	fmt.Println()

	s = time.Now()
	ans2 := part2(rules, messages)
	t = time.Now()
	e = t.Sub(s)
	fmt.Println("Answer for second question: ", ans2)
	fmt.Println("Second answer retrieved in: ", e)

}
