package main

import "fmt"

func main() {
	str := "hello,world"
	c := str
	fmt.Printf("%p %p\n", &str, &c)

	c += "!"
	fmt.Printf("%p %p\n", &str, &c)

	subStr := str[:5]
	fmt.Printf("subStr %s, str %s\n", subStr, str)

	str = "hi,world"
	fmt.Printf("subStr %s, str %s\n", subStr, str)

	subStr = "hi,nqq"
	fmt.Printf("subStr %s, str %s\n", subStr, str)

	//改变字符串底层值，报错cannot assign to subStr[:5]
	//subStr[:5] = "hi"
}
