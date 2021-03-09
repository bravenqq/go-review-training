// Package main provides ...
package main

import "fmt"

func main() {
	var any interface{}

	// any = "test"
	// fmt.Printf("%T\n", any)
	// fmt.Printf("%v\n", any)
	fmt.Printf("%p\n", &any)
	fmt.Println(any == nil)
	fmt.Println("=================")
	any = "test"
	fmt.Printf("%T\n", any)
	fmt.Printf("%v\n", any)
	fmt.Printf("%p\n", &any)
	fmt.Println(any == nil)

}
