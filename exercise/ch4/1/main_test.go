package main

import (
	"testing"
)

func TestDiffCount(t *testing.T) {
	tests := []struct {
		input1 [32]byte
		input2 [32]byte
		output uint8
	}{

		{input1: [32]byte{'a', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2'}, input2: [32]byte{'b', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2'}, output: 2},
	}
	for _, test := range tests {
		count := DiffCount(test.input1, test.input2)
		if test.output != count {
			t.Errorf("DiffCount input1:%x input2:%x get:%d want:%d", test.input1, test.input2, count, test.output)
		}
	}
}
