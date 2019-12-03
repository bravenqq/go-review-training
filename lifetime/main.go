package main

import (
	"fmt"
)

var a int

func main() {
	a = 10
	x := &a
	fmt.Printf("x %p\n", &x)
	incr(x)

	y := add(3, 4)
	x = y
	fmt.Printf("x %v,y %v\n", x, y)
	fmt.Printf("y %p\n", &y)

	arr := []int{10, 9, 8}
	for i := 0; i < len(arr); i++ {
		j := len(arr) - i
		arr[i] += j
	}
}

func incr(x *int) {
	fmt.Printf("x %p\n", &x)
	*x++
}

func add(a, b int) *int {
	y := a + b
	fmt.Printf("y %p\n", &y)
	return &y
}
