package main

import (
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		input1 string
		input2 func(string) string
		output string
	}{

		{"worldfoo worldfoo world", replace, "worldnqq worldnqq world"},
	}
	for _, test := range tests {
		res := Expand(test.input1, test.input2)
		if res != test.output {
			t.Errorf("Expand input:%s get:%s want:%s", test.input1, res, test.output)
		}
	}
}

func replace(s string) string {
	return "nqq"
}
