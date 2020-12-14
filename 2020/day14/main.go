package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func part1(input []string) int {

	var mask string
	memory := make(map[int]int)

	for _, row := range input {
		if strings.Contains(row, "mask") {
			mask = strings.Fields(row)[2]
		} else {
			re := regexp.MustCompile("[0-9]+")
			mv := re.FindAllString(row, -1)
			address := util.ToInt(mv[0])
			value := util.ToInt(mv[1])
			memory[address] = applyMask(value, mask)
		}
	}

	return sumMemory(memory)
}

func applyMask(value int, mask string) int {
	for i, m := range mask {
		pos := uint(len(mask) - i - 1)
		if m == 'X' {
			continue
		} else if m == '0' {
			value = util.ClearBit(value, uint(pos))
		} else if m == '1' {
			value = util.SetBit(value, uint(pos))
		}
	}
	return value
}

func sumMemory(memory map[int]int) int {
	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

func part2(input []string) int64 {

	var mask string
	memory := make(map[int64]int64)

	for _, row := range input {
		if strings.Contains(row, "mask") {
			mask = strings.Fields(row)[2]
		} else {
			re := regexp.MustCompile("[0-9]+")
			mv := re.FindAllString(row, -1)
			startAddress, _ := strconv.ParseInt((mv[0]), 10, 64)
			value, _ := strconv.ParseInt((mv[1]), 10, 64)
			memory = runMAD(value, mask, startAddress, memory)
		}
	}

	return sumMemory64(memory)

}

func runMAD(value int64, mask string, startAddress int64, memory map[int64]int64) map[int64]int64 {
	stringAddress := getBinaryRep(startAddress, 36)
	madmask := applyMADMask(stringAddress, mask)
	madmasks := expandMADMask(madmask)
	for _, m := range madmasks {
		i, _ := strconv.ParseInt(m, 2, 64)
		memory[i] = value
	}
	return memory

}

func getBinaryRep(value int64, size int) string {
	bin := strconv.FormatInt(value, 2)
	binSlice := strings.Split(bin, "")
	for i := 0; i < size-len(bin); i++ {
		binSlice = append([]string{"0"}, binSlice...)
	}
	return strings.Join(binSlice, "")
}

func applyMADMask(address string, mask string) string {
	addressSlice := strings.Split(address, "")
	for i, m := range mask {
		if m == '0' {
			continue
		} else if m == '1' {
			addressSlice[i] = "1"
		} else {
			addressSlice[i] = string(m)
		}
	}
	return strings.Join(addressSlice, "")
}

func expandMADMask(madmask string) []string {

	var expandedMADMask []string

	tmpMasks := []string{madmask}

	for {
		if len(tmpMasks) == 0 {
			return expandedMADMask
		}

		cMask := tmpMasks[0]
		tmpMasks = tmpMasks[1:]

		zmask := strings.Replace(cMask, "X", "0", 1)
		omask := strings.Replace(cMask, "X", "1", 1)

		if strings.Contains(zmask, "X") {
			tmpMasks = append(tmpMasks, zmask)
		} else {
			expandedMADMask = append(expandedMADMask, zmask)
		}
		if strings.Contains(omask, "X") {
			tmpMasks = append(tmpMasks, omask)
		} else {
			expandedMADMask = append(expandedMADMask, omask)
		}

	}
}

func sumMemory64(memory map[int64]int64) int64 {
	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}

func main() {

	rawInput := util.FetchInputForDay("2020", "14")
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
