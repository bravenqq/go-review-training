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

//声明了一个名称为Printer的接口
type Printer interface {
	Print()
}

func main() {
	//io.Writer接口类型变量w
	var w io.Writer
	//
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	fmt.Printf("address:%p\n", &w)
	if w == nil {
		fmt.Println("w is nil")
	}
}
