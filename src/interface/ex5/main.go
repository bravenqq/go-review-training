// Package main provides ...
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("name:%s age:%d", p.Name, p.Age)
}

func main() {
	person := Person{Name: "nqq", Age: 22}
	hwx := Person{Name: "hwx", Age: 48}
	//p的类型描述符为Person类型
	//p的动态类型值
	var p1 fmt.Stringer = person
	var p2 fmt.Stringer = hwx
	fmt.Println(p1 == p2)
}
