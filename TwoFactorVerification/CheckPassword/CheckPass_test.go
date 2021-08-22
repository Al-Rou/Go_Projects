package CheckPassword

import (
	"testing"
)

func TestTableCheckPass(t *testing.T) {
	var tests = []struct {
		inputUsername string
		inputPassword string
		expected      bool
	}{
		{"alex", "1234", false},
		{"Al", "3456", false},
		{"Al", "1234", true},
	}
	for _, check := range tests {
		output := CheckPass(check.inputUsername, check.inputPassword)
		if output != check.expected {
			t.Errorf("Test Failed: for \"%v\" and \"%v\", \"%v\" is expected, but \"%v\" was received!",
				check.inputUsername, check.inputPassword, check.expected, output)
		}
	}
}
