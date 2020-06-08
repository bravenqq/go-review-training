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
	//在栈中分配一个map保存value为*Student的数据
	students := make(map[string]*Student)
	//保存student到栈中的arraay中
	arr := []Student{
		{name: "nqq", age: 24},
		{name: "nieqianqian", age: 34},
	}
	//申明一个st变量，作用域为整个for循环
	for _, st := range arr {
		fmt.Printf("%T\n", st)
		//循环结束，st仍被使用，放到堆中
		students[st.name] = &st
	}

	//st地址保存arr中最后一个数据
	for k, v := range students {
		fmt.Println(k, ":", *v)
	}
}
