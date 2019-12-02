package main

import (
	"fmt"
)

func main() {
	p := new(int)
	q := new(string)
	fmt.Printf("p %v, q %v\n", p, q)
	*q = "hello"
	*p = 1
	fmt.Printf("p %v, q %v\n", p, q)
	fmt.Printf("*p %d, *q %s\n", *p, *q)
}
