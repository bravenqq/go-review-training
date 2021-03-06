package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		fmt.Println("test")
	}()
	defer func() {
		fmt.Println(string(debug.Stack()))
	}()
	f("")
}

func f(a string) {
	if a == "" {
		panic("a is nil")
	}
}
