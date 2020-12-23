package main

import (
	"container/ring"
	"fmt"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []int) string {

	ring, lut := createRing(input, len(input))
	ring = playCrab(ring, lut, 100)

	ring = lut[1]

	return stringRing(ring)[1:]
}

func part2(input []int) int {

	ring, lut := createRing(input, 1000000)
	ring = playCrab(ring, lut, 10000000)

	ring = lut[1]
	v1 := ring.Next()
	v2 := v1.Next()

	return v1.Value.(int) * v2.Value.(int)
}

func createRing(slice []int, length int) (*ring.Ring, map[int]*ring.Ring) {

	r := ring.New(length)
	largest := util.GetLargestInSlice(slice)
	lut := make(map[int]*ring.Ring)

	for _, s := range slice {
		r.Value = s
		lut[s] = r
		r = r.Next()
	}

	for i := largest + 1; i <= length; i++ {
		r.Value = i
		lut[i] = r
		r = r.Next()
	}

	return r, lut
}

func playCrab(ring *ring.Ring, lut map[int]*ring.Ring, iterations int) *ring.Ring {

	ringSize := ring.Len()

	i := 0
	for i < iterations {

		subRing := ring.Unlink(3)
		destination := ring.Value.(int) - 1

		if destination == 0 {
			destination = ringSize
		}

		for valueInRing(destination, subRing) {
			destination--
			if destination == 0 {
				destination = ringSize
			}
		}

		// At first I rotated the ring to find the correct place to put the subRing.
		// Then I rotated the ring back to ensure ring = ring.Next() on row 87 yield the
		// correct "current cup".
		// It took forever (after 15 minutes I was on iteration ~250 000) so I went to
		// the subreddit: https://www.reddit.com/r/adventofcode/comments/kimluc/2020_day_23_solutions/ggs00yp/
		// The trick below uses a map[int]*ring.Ring to keep track of which label correspond
		// to which ring orientation.
		// This yields an AMAZING speed up, thank you /u/A-UNDERSCORE-D and /u/status_maximizer for teaching
		// me about how garbage collect works in Golang. The script now terminates in ~3 seconds.
		dstRing := lut[destination]
		dstRing.Link(subRing)
		ring = ring.Next()

		i++

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

func printRing(ring *ring.Ring) {
	ring.Do(func(p interface{}) {
		fmt.Print(p.(int))
	})
	fmt.Println()
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
