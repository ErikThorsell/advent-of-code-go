package main

import (
	"fmt"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string) int {

	player1Deck := util.GetStringsAsInts(util.ParseInputByLine(input[0])[1:])
	player2Deck := util.GetStringsAsInts(util.ParseInputByLine(input[1])[1:])

	round := 0

	for {

		round++

		//		fmt.Println("P1:", player1Deck)
		//		fmt.Println("P2:", player2Deck)

		if len(player1Deck) == 0 || len(player2Deck) == 0 {
			break
		}

		c1 := player1Deck[0]
		player1Deck = player1Deck[1:]
		c2 := player2Deck[0]
		player2Deck = player2Deck[1:]

		if c1 > c2 {
			//			fmt.Println("Player 1 wins round", round)
			player1Deck = append(player1Deck, c1, c2)
		} else {
			//			fmt.Println("Player 2 wins round", round)
			player2Deck = append(player2Deck, c2, c1)
		}
		//		fmt.Println()

	}

	return calculateWinner(player1Deck, player2Deck)
}

func calculateWinner(player1Deck, player2Deck []int) int {
	winner := 0
	if len(player1Deck) != 0 {
		for i := 0; i < len(player1Deck)-1; i++ {
			card := player1Deck[len(player1Deck)-1-i]
			idx := i + 1
			winner += idx * card
		}
	} else if len(player2Deck) != 0 {
		for i := 0; i < len(player2Deck); i++ {
			idx := i + 1
			card := player2Deck[len(player2Deck)-idx]
			winner += idx * card
		}
	}
	return winner
}

func part2(input []string) int {

	player1Deck := util.GetStringsAsInts(util.ParseInputByLine(input[0])[1:])
	player2Deck := util.GetStringsAsInts(util.ParseInputByLine(input[1])[1:])

	player1Deck, player2Deck, _ = recursiveCombat(player1Deck, player2Deck)

	return calculateWinner(player1Deck, player2Deck)

}

func recursiveCombat(deck1, deck2 []int) ([]int, []int, int) {

	p1dh := make(map[string]bool)
	p2dh := make(map[string]bool)
	round := 0
	roundWinner := 0

	// This whole thing is inefficient.... ugh
	for !(len(deck1) == 0 || len(deck2) == 0) {

		c1, c2 := deck1[0], deck2[0]
		round++
		//		fmt.Println("Round:", round)

		//		fmt.Println("P1:", deck1)
		//		fmt.Println("P2:", deck2)
		sd1 := fmt.Sprintf("%v", deck1)
		sd2 := fmt.Sprintf("%v", deck2)

		if p1dh[sd1] || p2dh[sd2] {
			roundWinner = 1
		} else {
			p1dh[sd1] = true
			p2dh[sd2] = true

			if c1 > len(deck1)-1 || c2 > len(deck2)-1 {
				//				fmt.Println("A player does not have enough cards left")
				if c1 > c2 {
					roundWinner = 1
				} else if c1 < c2 {
					roundWinner = 2
				}
			} else {
				cd1 := make([]int, len(deck1[1:c1+1]))
				cd2 := make([]int, len(deck2[1:c2+1]))
				copy(cd1, deck1[1:c1+1])
				copy(cd2, deck2[1:c2+1])
				_, _, roundWinner = recursiveCombat(cd1, cd2)
			}
		}

		if roundWinner == 1 {
			//			fmt.Println("Player 1 wins round", round)
			deck1 = append(deck1, c1, c2)
		} else if roundWinner == 2 {
			//			fmt.Println("Player 2 wins round", round)
			deck2 = append(deck2, c2, c1)
		}

		deck1 = deck1[1:]
		deck2 = deck2[1:]

		//		fmt.Println()

	}

	return deck1, deck2, roundWinner

}

func main() {

	rawInput := util.FetchInputForDay("2020", "22")
	parsedInput := util.ParseInputByBlankLine(rawInput)
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
