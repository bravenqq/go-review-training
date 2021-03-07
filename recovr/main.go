// Package main provides ...
package main

import "fmt"

func main() {
	defer func() {
		v := recover()
		fmt.Println(v)
	}()
	panic("test")
}
