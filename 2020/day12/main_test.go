package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func TestDirFinder(t *testing.T) {

	currentDirection := "W"
	turnDir := "R"
	turnAngle := 90

	calculatedNewDir := findDirection(currentDirection, turnDir, turnAngle)

	expected := "N"

	if calculatedNewDir != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, calculatedNewDir)
	}

}

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 25

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func TestRotWP(t *testing.T) {

	instr := []string{"R0", "L0", "R90", "L90", "R180", "L180", "R270", "L270", "R360", "L360"}
	wX := -3
	wY := 1

	ewpc := [][]int{{-3, 1}, {-3, 1}, {1, 3}, {-1, -3}, {3, -1}, {3, -1}, {-1, -3}, {1, 3}, {-3, 1}, {-3, 1}}

	for i := range instr {

		fmt.Print(i, ": ", instr[i], " -> ")
		_, _, nX, nY := followInstruction(0, 0, wX, wY, instr[i])
		fmt.Println(nX, nY)

		if nX != ewpc[i][0] && nY != ewpc[i][1] {
			t.Errorf("Test failed, expected: '(%v,%v)', got:  '(%v,%v)'", ewpc[i][0], ewpc[i][1], nX, nY)
		}

	}
}

func TestMoveWPAndShip(t *testing.T) {

	sX := 58
	sY := -6
	wX := 4
	wY := -10

	snX, snY, wnX, wnY := followInstruction(sX, sY, wX, wY, "L180")

	esX := 58
	esY := -6
	ewX := -10
	ewY := 4

	if wnX != ewX && wnY != ewY && snX != esX && snY != esY {
		t.Errorf("Test failed, expected: '(%v,%v)' and '(%v,%v)', got:  '(%v,%v)' and '(%v,%v)'", esX, esY, ewX, ewY, snX, snY, wnX, wnY)
	}

}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 286

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	actual := part2(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}
