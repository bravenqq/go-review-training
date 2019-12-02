package main

import (
	"fmt"
)

func main() {
	a := 1
	x := &a
	y := &a
	a = *x + *y
	x = add(*x, *y)
}

func add(x, y int) *int {
	res := x + y
	return &res
}
