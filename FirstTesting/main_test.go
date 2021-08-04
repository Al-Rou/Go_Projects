package main

import (
	"testing"
)

func TestTableFactorial(t *testing.T) {
	var tests = [] struct{
		in int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{5, 120},
		{6, 720},
	}
	for _, test := range tests{
		if output := Factorial(test.in); output != test.expected{
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.in, test.expected, output)
		}
	}
}