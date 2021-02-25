// Package main provides ...
package main

import "fmt"

func main() {
	var slice []int
	//slice 与nil做比较，底层数组没分配内存
	if slice == nil {
		fmt.Println("slice is nil")
	}

	a := 10
	slice = make([]int, a, 10)
	if slice == nil {
		fmt.Println("slice still nil")
	}
}
