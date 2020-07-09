// Package main provides ...
package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(cap(a))
	fmt.Println(cap(b))
	fmt.Println(cap(c))
	fmt.Println(cap(s))

	var arr []int
	arr = append(arr, 1)
	fmt.Println("arr:", arr, " cap:", cap(arr))
	var m map[string]int
	// m["one"] = 1
	fmt.Println("m:", m)
}
