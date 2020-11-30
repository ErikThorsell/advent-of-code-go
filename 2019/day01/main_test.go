package main

import "testing"

func TestMain(t *testing.T) {

	inputs := []int{1969, 100756}
	outputs := []int{966, 50346}

	for i := 0; i < len(inputs); i++ {
		expected := outputs[i]
		actual := getFuelForFuel(inputs[i])

		if expected != actual {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
		}
	}
}
