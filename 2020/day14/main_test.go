package main

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func TestMasking(t *testing.T) {

	value := 11
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	result := 73

	actual := applyMask(value, mask)

	if result != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", result, actual)
	}
}

func Test1(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example")
	exampleResult := 165

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	actual := part1(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func Test2(t *testing.T) {

	exampleData, _ := ioutil.ReadFile("./example2")
	exampleResult := int64(208)

	parsedExampleData := util.ParseInputByLine(string(exampleData))
	actual := part2(parsedExampleData)

	if exampleResult != actual {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", exampleResult, actual)
	}
}

func TestBinaryRep(t *testing.T) {

	in := int64(42)
	er := "000000000000000000000000000000101010"

	actual := getBinaryRep(in, 36)

	if er != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", er, actual)
	}

}

func TestMadMask(t *testing.T) {

	addr := "000000000000000000000000000000101010"
	mask := "000000000000000000000000000000X1001X"
	errr := "000000000000000000000000000000X1101X"

	actual := applyMADMask(addr, mask)

	if errr != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", errr, actual)
	}

}

func TestMadMaskExpander(t *testing.T) {

	mask := "00000000000000000000000000000001X0XX"
	expanded := []string{
		"000000000000000000000000000000010000",
		"000000000000000000000000000000010001",
		"000000000000000000000000000000010010",
		"000000000000000000000000000000010011",
		"000000000000000000000000000000011000",
		"000000000000000000000000000000011001",
		"000000000000000000000000000000011010",
		"000000000000000000000000000000011011",
	}

	actual := expandMADMask(mask)

	if !reflect.DeepEqual(expanded, actual) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expanded, actual)
	}

}
