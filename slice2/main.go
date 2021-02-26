// Package main provides ...
package main

func main() {
	a := 10
	b := 20
	//make([]int, a, b) escapes to heap
	slice := make([]int, a, b)
	_ = slice
}

func add(slice []int, v int) []int {
	l := cap(slice)
	for i := 0; i < l+5; i++ {
		slice = append(slice, v)
	}

	// fmt.Printf("slice:%p \n", &slice)
	return slice
}
