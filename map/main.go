package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	parse_student()
}

func parse_student() {
	students := make(map[string]*Student)
	arr := []Student{
		{name: "nqq", age: 24},
		{name: "nieqianqian", age: 34},
	}
	for _, st := range arr {
		fmt.Printf("%T\n", st)
		students[st.name] = &st
	}
	for k, v := range students {
		fmt.Println(k, ":", *v)
	}
}
