package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{3, 2, 4, 6, 7}, 22},
		{[]int{}, 0},
		{[]int{2}, 2},
	}
	for _, test := range tests {
		res := Sum(test.input)
		if res != test.output {
			t.Errorf("sum input:%v get:%d want:%d\n", test.input, res, test.output)
		}
	}
}
