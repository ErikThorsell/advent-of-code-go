package util

// Coordinate is a struct with (x,y)
type Coordinate struct {
	X, Y int
}

// CoordinateWithDistance extends the Coordinate struct with a distance.
// D denotes how many steps it took to get to this coordinate.
type CoordinateWithDistance struct {
	X, Y, D int
}

// ManhattanWalk takes a list of instructions and walks the walk.
// The return value is a list of coordinates.
func ManhattanWalk(instructions []string) []Coordinate {

	var wc []Coordinate

	x := 0
	y := 0

	origo := Coordinate{X: x, Y: y}
	wc = append(wc, origo)

	for _, step := range instructions {

		direction := string(step[0])
		length := ToInt(step[1:])

		for i := 1; i <= length; i++ {

			switch direction {
			case "U":
				y++
			case "D":
				y--
			case "R":
				x++
			case "L":
				x--
			}

			wc = append(wc, Coordinate{x, y})

		}
	}

	return wc
}

// ManhattanWalkWithDistance extends the regular ManhattanWalk and also adds
// information about exactly how many steps it took to get to each coordinate.
func ManhattanWalkWithDistance(instructions []string) []CoordinateWithDistance {

	var wc []CoordinateWithDistance

	x := 0
	y := 0
	d := 0

	origo := CoordinateWithDistance{X: x, Y: y, D: d}
	wc = append(wc, origo)

	for _, step := range instructions {

		direction := string(step[0])
		length := ToInt(step[1:])

		for i := 1; i <= length; i++ {

			switch direction {
			case "U":
				y++
			case "D":
				y--
			case "R":
				x++
			case "L":
				x--
			}

			d++

			wc = append(wc, CoordinateWithDistance{x, y, d})

		}
	}

	return wc
}

// ManhattanDistance between two coordinates
func ManhattanDistance(c1 Coordinate, c2 Coordinate) int {
	return Abs(c1.X-c2.X) + Abs(c1.Y-c2.Y)
}

// ManhattanWalksIntersect returns the coordinates where two manhattan walks
// intersects each other.
func ManhattanWalksIntersect(walk1 []Coordinate, walk2 []Coordinate) []Coordinate {

	var intersectingCoordinates []Coordinate

	for i := 0; i < len(walk1); i++ {
		for j := 0; j < len(walk2); j++ {
			c1 := walk1[i]
			c2 := walk2[j]
			if c1.X == c2.X && c1.Y == c2.Y {
				if c1.X == 0 && c1.Y == 0 {
					continue
				}
				intersectingCoordinates = append(intersectingCoordinates, c1)
			}
		}
	}
	return intersectingCoordinates
}

// ManhattanWalksIntersectWithDistance cares only about how long it takes to
// reach each crossing.
// The return value is a slice with distances to crossings.
func ManhattanWalksIntersectWithDistance(walk1 []CoordinateWithDistance, walk2 []CoordinateWithDistance) []int {

	var intersectingCoordinatesDistance []int

	for i := 0; i < len(walk1); i++ {
		for j := 0; j < len(walk2); j++ {
			c1 := walk1[i]
			c2 := walk2[j]
			if c1.X == c2.X && c1.Y == c2.Y {
				if c1.X == 0 && c1.Y == 0 {
					continue
				}
				intersectingCoordinatesDistance = append(intersectingCoordinatesDistance, c1.D+c2.D)
			}
		}
	}
	return intersectingCoordinatesDistance
}
