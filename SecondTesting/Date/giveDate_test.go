package Date

import (
	"testing"
)

func TestTableGiveDate(t *testing.T) {
	var tests = []struct {
		inputYear   int
		inputNumber int
		expected    string
	}{
		{2021, 1, "1/January/2021"},
		{2020, 366, "31/December/2020"},
		{2020, 367, "The input number cannot be bigger than 366 or less than one!"},
		{2021, -23, "The input number cannot be bigger than 366 or less than one!"},
		{2021, 366, "The entered number cannot be 366 when the year is not a leap year!"},
		{2021, 32, "1/February/2021"},
		{2020, 60, "29/February/2020"},
	}
	for _, check := range tests {
		output := Date(check.inputYear, check.inputNumber)
		if output != check.expected {
			t.Errorf("Test Failed: for \"%v\" and \"%v\", \"%v\" is expected, but \"%v\" was received!",
				check.inputYear, check.inputNumber, check.expected, output)
		}
	}
}
