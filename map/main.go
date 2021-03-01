// Package main provides ...
package main

import "fmt"

func main() {
	var m map[int]string
	if m == nil {
		fmt.Println("map is nil")
		fmt.Println("map count:", len(m))
	}
}
