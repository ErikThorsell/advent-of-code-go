package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string) int {

	res := 0
	for _, row := range input {
		elems := strings.Split(strings.ReplaceAll(row, " ", ""), "")
		for _, e := range eval(elems) {
			res += e
		}
	}

	return res
}

func eval(expr []string) []string {
	res := term(expr)
	for {
		if len(expr) == 0 || expr[0] != "*" {
			return res
		}
		res *= term(expr[1:])
	}
}

func term(expr []string) []string {
	res := factor(expr)
	for {
		if len(expr) == 0 || expr[0] != "+" {
			return res
		}
		res += factor(expr[1:])
	}
}

func factor(expr []string) string {
	v, expr := expr[0], expr[1:]
	if v == "(" {
		v = eval(expr)
		expr = expr[1:]
	}
	return v
}

func part2(input []string) int {
	return 0
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
