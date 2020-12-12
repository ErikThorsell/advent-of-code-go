package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string) int {
	xPos := 0
	yPos := 0
	dir := "E"

	for _, i := range input {
		xPos, yPos, dir = moveShip(xPos, yPos, dir, i)
	}

	return util.Abs(0-xPos) + util.Abs(0-yPos)

}

func moveShip(xPos int, yPos int, dir string, instr string) (int, int, string) {

	action := string(instr[0])
	value := util.ToInt(instr[1:])

	switch action {
	case "N":
		yPos += value
	case "S":
		yPos -= value
	case "E":
		xPos += value
	case "W":
		xPos -= value
	case "L":
		dir = findDirection(dir, action, value)
	case "R":
		dir = findDirection(dir, action, value)
	case "F":
		xPos, yPos = moveForward(dir, value, xPos, yPos)
	}

	return xPos, yPos, dir

}

func findDirection(currentDir string, turnDir string, turnAngle int) string {

	dirs := []string{"N", "E", "S", "W"}
	turn := turnAngle / 90
	turnIdx := 0

	if turnDir == "R" {
		turnIdx = turn
	} else if turnDir == "L" {
		turnIdx = -turn
	}

	if currentDir == "N" {
		return dirs[util.RealMod(turnIdx, len(dirs))]
	}
	if currentDir == "E" {
		return dirs[util.RealMod(turnIdx+1, len(dirs))]
	}
	if currentDir == "S" {
		return dirs[util.RealMod(turnIdx+2, len(dirs))]
	}
	if currentDir == "W" {
		return dirs[util.RealMod(turnIdx+3, len(dirs))]
	}

	log.Fatal("Oops.")
	return "INVALID"

}

func moveForward(currentDir string, value int, xPos, yPos int) (int, int) {

	switch currentDir {
	case "N":
		yPos += value
	case "S":
		yPos -= value
	case "E":
		xPos += value
	case "W":
		xPos -= value
	}

	return xPos, yPos

}

func part2(input []string) int {
	sx, sy := 0, 0
	wx, wy := 10, 1

	for _, i := range input {
		sx, sy, wx, wy = followInstruction(sx, sy, wx, wy, i)
	}

	return util.Abs(0-sx) + util.Abs(0-sy)
}

func followInstruction(sX, sY, wX, wY int, instr string) (int, int, int, int) {

	action := string(instr[0])
	value := util.ToInt(instr[1:])

	switch action {
	case "N":
		wY += value
	case "S":
		wY -= value
	case "E":
		wX += value
	case "W":
		wX -= value
	case "L":
		wX, wY = rotateWayPoint(wX, wY, action, value)
	case "R":
		wX, wY = rotateWayPoint(wX, wY, action, value)
	case "F":
		sX += wX * value
		sY += wY * value
	}

	return sX, sY, wX, wY

}

func rotateWayPoint(wX, wY int, rotDir string, rotAmount int) (int, int) {

	noTurns := util.Abs(rotAmount / 90)

	for i := 0; i < noTurns; i++ {

		if rotDir == "R" {
			wX, wY = wY, -wX
		} else if rotDir == "L" {
			wX, wY = -wY, wX
		}
	}

	return wX, wY

}

func main() {

	rawInput := util.FetchInputForDay("2020", "12")
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
