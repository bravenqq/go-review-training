package main

import (
	"fmt"
)

// START OMIT
func main() {
	var x int
	x = 1
	p := &x
	*p = 156
	var count [10]int
	for i, _ := range count {
		count[i] += 100
	}

	var i, j, k int
	i, j, k = 3, 5, 10
	fmt.Println(i, j, k)
	fmt.Println(fib(100))
}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

// END OMIT
