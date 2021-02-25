// Package main provides ...
package main

func main() {
	var slice []int
	//slice 与nil做比较，底层数组没分配内存
	a := 1048576
	slice = make([]int, a, a)
	_ = slice
}

func f() []int {
	return make([]int, 0, 10)
}
