// Package main provides ...
package main

import "fmt"

func main() {

	//计算机中存储的是字符对应码点编码之后的值
	data := []byte{'a', 'b', 'd'}
	fmt.Printf("data len:%d cap:%d\n", len(data), cap(data))
	for _, c := range data {
		fmt.Printf("char:%c code:%d type:%T\n", c, c, c)
	}

	data[2] = 'c'

	str := []rune{'聂', '倩', '倩'}
	fmt.Printf("str len:%d cap:%d\n", len(str), cap(str))
	for _, c := range str {
		fmt.Printf("char:%c code:%d type:%T\n", c, c, c)
	}

	s := "聂倩倩"
	//s的底层仍然是[]byte类型，所以len函数是按byte字节计算长度
	fmt.Printf("s len:%d\n", len(s))

	//将string转换成str之后按rune类型去长度
	str = []rune(s)
	fmt.Printf("str len:%d\n", len(str))

	s = "abc"
	data = []byte(s)
	data[0] = 'd'
	for _, c := range data {
		fmt.Printf("char:%c code:%d type:%T\n", c, c, c)
	}

	fmt.Println("s:", s)

}
