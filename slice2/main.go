// Package main provides ...
package main

func main() {
	slice := f()
	_ = slice
}

func f() []int {
	//make([]int, 10, 20) escapes to heap
	slice := make([]int, 10, 20)
	return slice
}
