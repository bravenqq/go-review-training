package main

import (
	"testing"
)

func TestQsort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{8, 20, 10, 5, 9}, []int{5, 8, 9, 10, 20}},
	}
	for _, test := range tests {
		res := Qsort(test.input)
		if !equal(res, test.output) {
			t.Errorf("Qsort input:%v want:%v get:%v\n", test.input, test.output, res)
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
