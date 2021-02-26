// Package main provides ...
package main

func main() {
	a := 5
	b := 6
	slice := make([]int, a, b)
	//slice does not escape
	add(slice, 10)
}

func add(slice []int, v int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = v
	}
}
