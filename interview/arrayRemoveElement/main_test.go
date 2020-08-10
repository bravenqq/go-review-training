package main

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		arr    []int
		num    int
		order  bool
		output []int
	}{
		{nil, 0, true, nil},
		{[]int{3, 4, 5, 6, 7, 10}, 0, true, []int{3, 4, 5, 6, 7, 10}},
		// {[]int{3, 4, 5, 6, 7, 10}, 0, false, []int{3, 4, 5, 6, 7, 10}},
		// {[]int{10, 10, 10}, 0, false, []int{}},
		{[]int{10, 10, 10}, 0, true, []int{}},
		{[]int{3, 4, 5, 6, 7, 10, 20, 10, 0}, 10, true, []int{3, 4, 5, 6, 7, 20, 0}},
		// {[]int{3, 4, 5, 6, 7, 10, 20, 10, 0}, 10, false, []int{3, 4, 5, 6, 0, 20}},
		// {[]int{3, 4, 5, 6, 7, 10, 20, 10, 0, 10, 10, 10}, 10, false, []int{3, 4, 5, 6, 0, 20}},
	}
	for _, test := range tests {
		fmt.Printf("%+v\n", test.arr)
		arr := Remove(test.arr, test.num, test.order)
		if !equal(arr, test.output) {
			t.Errorf("arr:%+v num:%d order:%t output:%+v got:%+v\n", test.arr, test.num, test.order, test.output, arr)
		}
	}
}

func equal(arr1, arr2 []int) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}
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
