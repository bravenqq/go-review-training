package main

import (
	"fmt"
	"io"
)

func main() {
	var a int
	fmt.Println(a)
	var res bool
	fmt.Println(res)
	var str string
	fmt.Printf("str value:%s lengths:%d\n", str, len(str))
	var arr [10]int
	fmt.Printf("arr value:%v lengths:%d\n", arr, len(arr))
	var person Person
	fmt.Println(person)
	var in io.Reader
	fmt.Println(in)
	// var data []byte
	// in.Read(data)
	var slice []int
	fmt.Printf("slice value %v lengths:%d\n", slice, len(slice))
	var p *Person
	// fmt.Printf("p value %v name:%s age:%d\n", p, p.name, p.age)
	fmt.Printf("p value %v \n", p)
	var counts map[string]int
	fmt.Printf("counts value %v lengths:%d\n", counts, len(counts))
	var add func(int, int) int
	fmt.Println(add)
}

type Person struct {
	name string
	age  int
}
