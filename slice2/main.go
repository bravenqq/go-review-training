// Package main provides ...
package main

import "fmt"

func main() {
	var slice []int
	//slice 与nil做比较，底层数组没分配内存
	if slice == nil {
		fmt.Println("slice is nil")
	}
	slice = make([]int, 0, 10)
	if slice == nil {
		fmt.Println("slice still nil")
	}
}

func f() []int {
	return make([]int, 0, 10)
}
