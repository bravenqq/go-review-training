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
}

type Node struct {
	next *Node
	Data Person
}

type Person struct {
	Name string
	Age  int
}
