package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if CalSum(2, 3) != 5 {
		t.Error("Expected 2+3=5")
	}
}

func TestTblDataToCalculate(t *testing.T) {
	var tests = []struct {
		inputx   int
		inputy   int
		expected int
	}{
		{2, 3, 5},
		{-4, 4, 0},
		{5, 9, 14},
		{22222, 5, 22227},
	}

	for _, test := range tests {
		if output := CalSum(test.inputx, test.inputy); output != test.expected {
			t.Error("Test Failed {} Expected, received: {}", test.expected, output)
		}
	}

}
