package main

import (
	"reflect"
	"testing"

	"github.com/ErikThorsell/advent-of-code-go/util"
)

func TestMain(t *testing.T) {

	inputs := [][]int{
		{1, 0, 0, 0, 99},
		{2, 3, 0, 3, 99},
		{2, 4, 4, 5, 99, 0},
		{1, 1, 1, 4, 99, 5, 6, 0, 99},
	}
	outputs := [][]int{
		{2, 0, 0, 0, 99},
		{2, 3, 0, 6, 99},
		{2, 4, 4, 5, 99, 9801},
		{30, 1, 1, 4, 2, 5, 6, 0, 99},
	}

	for i := 0; i < len(inputs); i++ {
		expected := outputs[i]
		actual := util.RunProgram(inputs[i])

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
		}
	}
}
