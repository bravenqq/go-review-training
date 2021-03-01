// Package main provides ...
package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[1] = "hello"
	fmt.Println(m[1])
	fmt.Println(m[2])
	// if m == nil {
	// 	fmt.Println("map is nil")
	// 	fmt.Println("map count:", len(m))
	// }
}
