// Package main provides ...
package main

import "fmt"

func main() {
	a := 1
	f := func() {
		a += 2
		fmt.Println("a=", a)
	}
	f()
	fmt.Println("a=", a)
	func(a int) {
		a += 2
		fmt.Println("a=", a)
	}(a)
	fmt.Println("a=", a)
}
