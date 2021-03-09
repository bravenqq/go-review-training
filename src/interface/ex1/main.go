// Package main provides ...
//接口类型
package main

import (
	"fmt"
	"io"
	"os"
)

type Person struct {
}

type Printer interface {
	Print()
}

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	fmt.Printf("address:%p\n", &w)
	if w == nil {
		fmt.Println("w is nil")
	}
}
