package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(nodeMap map[string]util.SatelliteNode, messages []string) int {

	validPatterns := getValidPatterns(nodeMap)
	fmt.Println(validPatterns)

	return -1

}

func getValidPatterns(nodeMap map[string]util.SatelliteNode) map[string][]string {

	// Map an id to all patterns it supports
	patterns := initializePatterns(nodeMap)
	nodes := getNodesInMap(nodeMap)

	// If we were able to initilaze a node's pattern, discard the node
	for _, n := range getPatternsInMap(patterns) {
		nodes = util.RemoveStringByValue(n, nodes)
	}

	for {

		if len(nodes) == 0 {
			return patterns
		}

		for _, n := range nodes {
			node := nodeMap[n]

			// If all dep nodes have known patterns, this node can be resolved
			if util.StringsInSlice(node.Deps, getPatternsInMap(patterns)) {
				patterns[n] = resolveNode(node, patterns)
				nodes = util.RemoveStringByValue(n, nodes)
			}
		}

	}

}

func initializePatterns(nodeMap map[string]util.SatelliteNode) map[string][]string {

	patterns := make(map[string][]string)

	for _, nodeID := range getNodesInMap(nodeMap) {
		node := nodeMap[nodeID]
		if len(node.Rule) == 1 {
			patterns[nodeID] = append(patterns[nodeID], node.Rule)
		}
	}

	return patterns

}

func resolveNode(node util.SatelliteNode, patterns map[string][]string) []string {

}

func getInheritedPatterns(node util.SatelliteNode, patterns map[string][]string) []string {
	rules := parseRule(node.Rule)
	for _, rule := range rules {
		for _, node := range strings.Fields(rule) {
			return patterns[node]
		}
	}
	return []string{}
}

func parseRule(rule string) []string {
	rules := []string{}
	for _, r := range strings.Split(rule, "|") {
		rules = append(rules, r)
	}
	return rules
}

func getPatternsInMap(patterns map[string][]string) []string {
	keys := make([]string, 0, len(patterns))
	for k := range patterns {
		keys = append(keys, k)
	}
	return keys
}

func getNodesInMap(nodeMap map[string]util.SatelliteNode) []string {
	keys := make([]string, 0, len(nodeMap))
	for k := range nodeMap {
		keys = append(keys, k)
	}
	return keys
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
