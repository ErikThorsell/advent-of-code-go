package util

// Cube denotes a cube
type Cube struct {
	X int
	Y int
	Z int
}

// HyperCube is a 4D cube
type HyperCube struct {
	X int
	Y int
	Z int
	W int
}

// GetAdjacentCubes3D returns all adjacent cubes for the specified cube coordinate
func GetAdjacentCubes3D(cube Cube) []Cube {

	adjacentCubes := []Cube{}

	for _, dx := range MakeRange(-1, 1) {
		for _, dy := range MakeRange(-1, 1) {
			for _, dz := range MakeRange(-1, 1) {
				if dx == dy && dy == dz && dz == 0 {
					continue
				} else {
					adjacentCubes = append(adjacentCubes, Cube{cube.X + dx, cube.Y + dy, cube.Z + dz})
				}
			}
		}
	}
	return adjacentCubes
}

// GetAdjacentCubes4D returns all adjacent cubes for the specified cube coordinate
func GetAdjacentCubes4D(cube HyperCube) []HyperCube {

	adjacentCubes := []HyperCube{}

	for _, dx := range MakeRange(-1, 1) {
		for _, dy := range MakeRange(-1, 1) {
			for _, dz := range MakeRange(-1, 1) {
				for _, dw := range MakeRange(-1, 1) {
					if dx == dy && dy == dz && dz == dw && dw == 0 {
						continue
					} else {
						adjacentCubes = append(adjacentCubes, HyperCube{cube.X + dx, cube.Y + dy, cube.Z + dz, cube.W + dw})
					}
				}
			}
		}
	}
	return adjacentCubes
}

func cubeInSlice(cube Cube, slice []Cube) bool {
	for _, c := range slice {
		if c == cube {
			return true
		}
	}
	return false
}

func hyperCubeInSlice(cube HyperCube, slice []HyperCube) bool {
	for _, c := range slice {
		if c == cube {
			return true
		}
	}
	return false
}

// SimulateBlockOfCubes runs the block simulation one iteration
func SimulateBlockOfCubes(prevActiveCubes []Cube) []Cube {

	numberOfActiveCubes := make(map[Cube]int)
	for _, cube := range prevActiveCubes {
		for _, adjacentCube := range GetAdjacentCubes3D(cube) {
			numberOfActiveCubes[adjacentCube]++
		}
	}

	newActiveCubes := []Cube{}
	for _, cube := range prevActiveCubes {
		active := numberOfActiveCubes[cube]
		if active == 2 || active == 3 {
			newActiveCubes = append(newActiveCubes, cube)
		}
	}

	for cube := range numberOfActiveCubes {
		if numberOfActiveCubes[cube] == 3 {
			if !cubeInSlice(cube, newActiveCubes) {
				newActiveCubes = append(newActiveCubes, cube)
			}
		}
	}

	return newActiveCubes

}

// SimulateBlockOfHyperCubes runs the block simulation one iteration
func SimulateBlockOfHyperCubes(prevActiveCubes []HyperCube) []HyperCube {

	numberOfActiveCubes := make(map[HyperCube]int)
	for _, cube := range prevActiveCubes {
		for _, adjacentCube := range GetAdjacentCubes4D(cube) {
			numberOfActiveCubes[adjacentCube]++
		}
	}

	newActiveCubes := []HyperCube{}
	for _, cube := range prevActiveCubes {
		active := numberOfActiveCubes[cube]
		if active == 2 || active == 3 {
			newActiveCubes = append(newActiveCubes, cube)
		}
	}

	for cube := range numberOfActiveCubes {
		if numberOfActiveCubes[cube] == 3 {
			if !hyperCubeInSlice(cube, newActiveCubes) {
				newActiveCubes = append(newActiveCubes, cube)
			}
		}
	}

	return newActiveCubes

}
