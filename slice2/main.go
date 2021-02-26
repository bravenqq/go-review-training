// Package main provides ...
package main

func main() {
	var slice []int
	//slice 与nil做比较，底层数组没分配内存
	slice = make([]int, 5, 10)
}

func f() []int {
	return make([]int, 0, 10)
}
