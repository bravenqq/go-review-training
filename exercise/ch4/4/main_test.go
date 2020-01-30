// Package main provides ...
package main

import (
	"testing"
)

func TestLeftRotate(t *testing.T) {

	tests := []Test{Test{input1: []int{0, 1, 2, 3, 4, 5}, input2: 2, output: []int{2, 3, 4, 5, 0, 1}}}
	for _, test := range tests {
		res := LeftRotate(test.input1, test.input2)
		if !test.equal(res) {
			t.Errorf("LeftRotate input1:%v left rotate:%d get:%v want:%v\n", test.input1, test.input2, res, test.output)
		}
	}
}

func (t Test) equal(res []int) bool {
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
	input1 []int
	input2 int
	output []int
}
