package main

import (
	"fmt"
)

func main() {
	var x int = 124
	var y int = 12
	q := &x
	p := &x
	t := &y
	fmt.Printf("p: %v, *p %d\n", p, *p)
	*p = 100
	fmt.Printf("*p %d, x %d\n", *p, x)
	fmt.Println(p == q, p != t, p != nil && q != nil && t != nil)

	t = add(x, y)
	fmt.Printf("t %v\n", t)
	fmt.Println(add(x, y) == add(x, y))
	incr(p)
	fmt.Printf("p: %v, *p %d, x %d\n", p, *p, x)
}

func incr(p *int) {
	*p++
}

func add(x, y int) *int {
	res := x + y
	fmt.Printf("res %v\n", &res)
	return &res
}
