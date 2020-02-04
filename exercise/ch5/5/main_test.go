package main

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"hello 聂倩倩", 2},
	}

	for _, test := range tests {
		res := CountWords(test.input)
		if res != test.output {
			t.Errorf("CountWords input:%s get:%d want:%d", test.input, res, test.output)
		}
	}
}
