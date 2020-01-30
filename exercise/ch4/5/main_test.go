package main

import (
	"testing"
)

func TestDeduplicate(t *testing.T) {
	tests := []Test{Test{input: []string{"a", "bb", "dd", "cc", "cc", "bb"}, output: []string{"a", "bb", "dd", "cc", "bb"}}}
	for _, test := range tests {
		input := test.input
		res := Deduplicate(test.input)
		if !test.equal(res) {
			t.Errorf("Deduplicate input:%v get:%v want:%v\n", input, test.input, test.output)
		}
	}
}

func (t Test) equal(res []string) bool {
	if len(res) != len(t.output) {
		return false
	}
	for i := 0; i < len(res); i++ {
		if res[i] != t.output[i] {
			return false
		}
	}
	return true

}

type Test struct {
	input  []string
	output []string
}
