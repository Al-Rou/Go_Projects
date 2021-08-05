package Leap

import "testing"

func TestTableLeapYear(t *testing.T) {
	var tests = []struct {
		input    int
		expected bool
	}{
		{2020, true},
		{2021, false},
		{1900, false},
		{1600, true},
	}
	for _, check := range tests {
		output := IsLeap(check.input)
		if output != check.expected {
			t.Errorf("Test Failed: for \"%v\", \"%v\" is expected, but \"%v\" was received!", check.input, check.expected, output)
		}
	}
}
