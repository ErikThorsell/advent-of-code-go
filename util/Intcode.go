package util

// RunProgram which runs Intcode Programs
func RunProgram(integers []int) []int {

	for i := 0; i < len(integers); i += 4 {

		opCode := integers[i]

		if opCode == 99 {
			return integers
		}

		i1 := integers[i+1]
		v1 := integers[i1]

		i2 := integers[i+2]
		v2 := integers[i2]

		io := integers[i+3]

		if opCode == 1 {
			integers[io] = v1 + v2
		} else if opCode == 2 {
			integers[io] = v1 * v2
		}

	}

	return integers

}
