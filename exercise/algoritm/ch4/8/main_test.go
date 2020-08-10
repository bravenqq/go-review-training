package main

import (
	"testing"
)

func TestMultiplicat(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{2, 4, 5}, []int{80, 160, 200}},
	}
	for _, test := range tests {
		res := Multiplicat(test.input, test.input)
		if equal(res, test.output) {
			t.Errorf("Multiplicat input:%v get:%v want:%v\n", test.input, res, test.output)
		}
	}
}
func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
