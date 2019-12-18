// Package main provides ...
package main

import "fmt"

func main() {
	//计算机中存储的是字符对应码点编码之后的值
	//在编辑器中使用设置的utf-8解码序列值展示成'a','b','c'字符
	data := []byte{'a', 'b', 'd'}
	fmt.Printf("data len:%d cap:%d\n", len(data), cap(data))
	for _, c := range data {
		fmt.Printf("char:%c code:%d type:%T\n", c, c, c)
	}
	fmt.Println("------------------------------------------------")
	//s字符串中存储的是字符对应码点编码之后的值
	//在编辑器中使用设置的utf-8解码序列值展示成"abc"字符串
	s := "abc"
	for i := 0; i < len(s); i++ {
		fmt.Printf("char:%c code:%d type:%T\n", s[i], s[i], s[i])
	}
	fmt.Printf("s addr:%p\n", &s)
	fmt.Println("------------------------------------------------")
	//此处copy data中的值到新的内存中
	//s指向新的内存
	s = string(data)
	fmt.Printf("s addr:%p\n", &s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("char:%c code:%d type:%T\n", s[i], s[i], s[i])
	}
	fmt.Println("------------------------------------------------")
	//由于s = string(data)是开辟了新内存copy data值，所以修改data不会改变s的值
	data[0] = 'e'
	for i := 0; i < len(s); i++ {
		fmt.Printf("char:%c code:%d type:%T\n", s[i], s[i], s[i])
	}
}
