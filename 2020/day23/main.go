package main

import (
	"container/ring"
	"fmt"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []int) string {

	ring := createRing(input, 0)
	ring = playCrab(ring, 100)
	ring = putOneFirst(ring)
	fmt.Print("Sorted ring: ")
	printRing(ring)

	return stringRing(ring)[1:]
}

func playCrab(ring *ring.Ring, iterations int) *ring.Ring {

	ringSize := ring.Len()

	i := 0
	for i < iterations {

		//		fmt.Print("Ring: ")
		//		printRing(ring)

		currentCup := ring.Value.(int)
		destination := currentCup - 1
		if destination == 0 {
			destination = ringSize
		}
		subRing := ring.Unlink(3)

		for valueInRing(destination, subRing) {
			destination--
			destination = util.RealMod(destination, ringSize)
			if destination == 0 {
				destination = ringSize
			}
		}

		//		fmt.Println("Current:", currentCup)
		//		fmt.Println("Destination:", destination)
		//		fmt.Print("Sub Ring:")
		//		printRing(subRing)

		//		fmt.Println("Rotating to find destination:")
		for {
			//			fmt.Print(" > ")
			//			printRing(ring)
			v := ring.Value
			if v == destination {
				break
			}
			ring = ring.Next()
		}

		//		fmt.Print("Before merge: ")
		//		printRing(ring)

		ring = ring.Link(subRing)
		//		fmt.Print("After merge: ")
		//		printRing(ring)

		for ring.Value != currentCup {
			ring = ring.Next()
		}
		//		fmt.Print("After move back: ")
		//		printRing(ring)

		//		fmt.Println()

		// Rotate to get next current cup
		ring = ring.Next()

		i++

	}

	return ring
}

func putOneFirst(ring *ring.Ring) *ring.Ring {
	for i := 0; i < ring.Len(); i++ {
		if ring.Value != 1 {
			ring = ring.Next()
		}
	}
	return ring
}

func stringRing(ring *ring.Ring) string {
	s := ""
	for i := 0; i < ring.Len(); i++ {
		s += fmt.Sprint(ring.Value.(int))
		ring = ring.Next()
	}
	return s
}

func valueInRing(value int, ring *ring.Ring) bool {
	for i := 0; i < ring.Len(); i++ {
		if value == ring.Value {
			return true
		}
		ring = ring.Next()
	}
	return false
}

func createRing(slice []int, extra int) *ring.Ring {

	if extra > 0 {
		largest := util.GetLargestInSlice(slice)
		for i := largest + 1; i <= extra; i++ {
			slice = append(slice, i)
		}
	}

	ring := ring.New(len(slice))
	for _, s := range slice {
		ring.Value = s
		ring = ring.Next()
	}
	return ring
}

func printRing(ring *ring.Ring) {
	ring.Do(func(p interface{}) {
		fmt.Print(p.(int))
	})
	fmt.Println()
}

func part2(input []int) int {

	//ring := createRing(input, 1000000-len(input))
	//ring = playCrab(ring, 10000000)
	ring := createRing(input, 10-len(input))
	printRing(ring)
	ring = playCrab(ring, 10)
	ring = putOneFirst(ring)
	v1 := ring.Value.(int)
	ring = ring.Next()
	v2 := ring.Value.(int)

	return v1 * v2
}

func main() {

	rawInput := util.FetchInputForDay("2020", "23")
	parsedInput := util.GetStringsAsInts(strings.Split(rawInput, ""))
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
