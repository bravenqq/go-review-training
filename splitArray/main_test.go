// Package main provides ...
package main

import (
	"testing"
)

func TestSplitArray(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		output [][]int
	}{
		{input1: []int{}, input2: 10, output: nil},
		{input1: []int{3, 5, 7}, input2: 0, output: nil},
		{input1: []int{3, 5, 10, 20, 4, 7}, input2: 8, output: [][]int{[]int{3}, []int{5}, []int{10}, []int{20}, []int{4}, []int{7}}},
		{input1: []int{3, 5, 10, 20, 4, 7}, input2: 2, output: [][]int{[]int{3, 5, 10}, []int{20, 4, 7}}},
		{input1: []int{3, 5, 10, 20, 4, 7}, input2: 4, output: [][]int{[]int{3, 5}, []int{10, 20}, []int{4}, []int{7}}},
	}

	for _, test := range tests {
		arr, _ := splitArray(test.input1, test.input2)
		if len(test.input1) == 0 || test.input2 == 0 {
			if test.output != nil {
				t.Errorf("intput1:%+v,input2:%v,output:%v want:%v", test.input1, test.input2, arr, test.output)
			}
		}

		if !equalArr(test.output, arr) {
			t.Errorf("intput1:%+v,input2:%v,output:%v want:%v", test.input1, test.input2, arr, test.output)
		}
	}
}

func equalArr(arr1 [][]int, arr2 [][]int) bool {

}
