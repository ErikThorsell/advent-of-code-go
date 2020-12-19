package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func part1(nodeMap map[string]util.SatelliteNode, messages []string) int {
	validRegExp := getValidRegexp("0", nodeMap, "1")
	return countMatches(messages, validRegExp, "1")

}

func part2(nodeMap map[string]util.SatelliteNode, messages []string) int {

	id := "8"
	rule := "42 | 42 8"
	nodeMap[id] = util.SatelliteNode{ID: id, Rule: rule}

	id = "11"
	rule = "42 31 | 42 11 31"
	nodeMap[id] = util.SatelliteNode{ID: id, Rule: rule}

	validRegExp := getValidRegexp("0", nodeMap, "2")
	return countMatches(messages, validRegExp, "2")
}

func getValidRegexp(nodeID string, nodeMap map[string]util.SatelliteNode, part string) string {

	node := nodeMap[nodeID]

	// If the rule contains " it's a character ("a" or "b")
	if strings.Contains(node.Rule, "\"") {
		return strings.ReplaceAll(node.Rule, "\"", "")
	}

	pattern := "(" // every subExp is encapsulated in a capturing group

	subRules := strings.Split(node.Rule, "|")
	subRulesPatterns := []string{}

	for _, subRule := range subRules {
		concatRules := []string{}
		for _, subNode := range strings.Fields(subRule) { // for each subNode in the subRule

			if part == "2" {
				if subNode == "8" {
					sub8 := fmt.Sprintf("%s+", getValidRegexp("42", nodeMap, part))
					concatRules = append(concatRules, sub8)
					continue
				}
				if subNode == "11" {
					sub8 := getValidRegexp("42", nodeMap, part)
					sub11 := getValidRegexp("31", nodeMap, part)
					sub8And11 := fmt.Sprintf("(?<eleven>(%s%s|%[1]s(?&eleven)%[2]s))", sub8, sub11)
					concatRules = append(concatRules, sub8And11)
					continue
				}
			}

			subExp := getValidRegexp(subNode, nodeMap, part) // get the regexp for the subNode
			concatRules = append(concatRules, subExp)
		}

		subRulesPatterns = append(subRulesPatterns, strings.Join(concatRules, ""))

	}

	pattern += strings.Join(subRulesPatterns, "|")
	pattern += ")"
	return pattern

}

func countMatches(strings []string, regString string, part string) int {
	matchCount := 0
	if part == "1" {
		re := regexp.MustCompile("^" + regString + "$")
		for _, message := range strings {
			matches := re.MatchString(message)
			if matches {
				matchCount++
			}
		}
	} else {
		recurseTest := pcre.MustCompile("^"+regString+"$", 0)
		for _, v := range strings {
			if recurseTest.MatcherString(v, 0).Matches() {
				matchCount++
			}
		}
	}
	return matchCount
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
