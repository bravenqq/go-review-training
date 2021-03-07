package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		fmt.Println("test")
	}()
	go f("")
}

func f(a string) {
	defer func() {
		fmt.Println(string(debug.Stack()))
	}()
	if a == "" {
		panic("a is nil")
	}
}
