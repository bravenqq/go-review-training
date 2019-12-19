package main

import (
	"fmt"
)

func main() {
	//等价于
	//var a int
	//p = &a
	p := new(int)
	q := new(string)
	fmt.Printf("p %v, q %v\n", p, q)
	*q = "hello"
	*p = 1
	fmt.Printf("p %v, q %v\n", p, q)
	fmt.Printf("*p %d, *q %s\n", *p, *q)
}
