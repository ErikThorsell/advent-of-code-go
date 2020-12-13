package util

import (
	"strings"
)

// HHInstr correspond to a Hand Held Instruction
type HHInstr struct {
	op  string
	arg int
}

// GetNumInstr returns the number of instructions that are interesting
func GetNumInstr(instructions []HHInstr) int {
	numNJ := 0
	for _, inst := range instructions {
		if inst.op == "nop" || inst.op == "jmp" {
			numNJ++
		}
	}
	return numNJ
}

// ParseInstructions returns a list of Handheld instructions
func ParseInstructions(input []string) []HHInstr {
	var instructions []HHInstr
	for _, inst := range input {
		split := strings.Fields(inst)
		op := split[0]
		arg := ToInt(split[1])
		instructions = append(instructions, HHInstr{op, arg})
	}
	return instructions
}

// AlternateProgram swappes the variant:th instruction
func AlternateProgram(instructions []HHInstr, variant int) []HHInstr {

	cp := make([]HHInstr, len(instructions))
	copy(cp, instructions)

	ci := 0
	for i := range cp {

		if cp[i].op == "nop" || cp[i].op == "jmp" {

			if ci >= variant {
				switch cp[i].op {
				case "nop":
					cp[i].op = "jmp"
				case "jmp":
					cp[i].op = "nop"
				}
				break
			}

			ci++

		}

	}
	return cp
}

// RunHandHeld executes a Hand Held program
func RunHandHeld(instructions []HHInstr) (bool, int) {

	acc := 0

	idx := 0
	vi := []int{}

	instCount := 0
	for {

		if IntInSlice(idx, vi) {
			return false, acc
		}
		if idx == len(instructions) {
			return true, acc
		}

		vi = append(vi, idx) // keep track of visited indicies

		inst := instructions[idx]
		switch inst.op {
		case "nop":
			idx++
		case "acc":
			idx++
			acc += inst.arg
		case "jmp":
			idx += inst.arg
		}

		instCount++
		//fmt.Println(inst.op, inst.arg, "|", instCount)
		//fmt.Println()

	}

}
