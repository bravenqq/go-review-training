// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	a()
	fmt.Println("test")
}

func a() {
	i := 0
	fmt.Printf("0.i:%p\n", &i)
	defer func() {
		fmt.Printf("1.i:%p\n", &i)
		fmt.Println(i)
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	i++
	if i == 1 {
		panic(i)
	}
	defer func(i int) {
		fmt.Printf("2.i:%p\n", &i)
		fmt.Println(i)
	}(i)
}
