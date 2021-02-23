// Package main provides ...
package main

import "fmt"

const (
	a int = 10*iota + 1
	b
	c = 10*iota + 1
	d
	e
)

func main() {
	fmt.Println(a, b, c, d, e)
}
