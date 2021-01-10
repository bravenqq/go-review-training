package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var slice []byte
	var a struct{}
	var n Node
	fmt.Println("slice size: ", unsafe.Sizeof(slice))
	fmt.Println("a size:", unsafe.Sizeof(a))
	fmt.Println("n size:", unsafe.Sizeof(n))

	// fillstruct

	var nqq = Person{
		Name: "fesf",
		Age:  0,
	}

	fmt.Println(nqq)
	// anonymouse struct | function

	var abbynie = struct {
		Name string
		age  int
	}{
		"niexiaoqian",
		16,
	}

	fmt.Println(abbynie)

}

type Node struct {
	next *Node
	Data Person
}

type Person struct {
	Name string
	Age  int
}
