package main

import (
	"testing"
)

func TestEqualLetter(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		output bool
	}{
		{input1: "abcde", input2: "edcba", output: true},
		{input1: "abcded", input2: "edcdba", output: true},
		{input1: "abcded", input2: "edcdfa", output: false},
	}
	for _, test := range tests {
		res := EqualLetter(test.input1, test.input2)
		if test.output != res {
			t.Errorf("EqualLetter input1:%s input2:%s get:%t want:%t\n", test.input1, test.input2, res, test.output)
		}
	}
}
