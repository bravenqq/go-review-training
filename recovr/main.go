// Package main provides ...
package main

import (
	"fmt"
)

func main() {

	defer func() {
		v := recover()
		fmt.Println(v)
	}()
	defer func() {
		fmt.Println("hello,world")
	}()
	panic("test panic")
}
