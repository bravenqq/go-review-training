package main

import "fmt"

func main() {
	//rune是int32的别名
	//rune存储的是用编辑器设定的的编码方式编码的字符对应码点值
	//每个字符占用的内存与编码方式有关
	//与[]byte只能存储占用一字节的英文字符不同，rune可以存储任意的unicode集中字符
	// rune用编辑器设定的编码方式读取码点值转换成字符展示出来
	r := []rune{'a', 'b', 'c', 'd', '聂', '倩'}
	fmt.Println("r len:", len(r))
	for _, c := range r {
		fmt.Printf("rune:%c code:%d type:%T\n", c, c, c)
	}
	s := string(r)
	//len(s)是计算字符串占用的字节数
	//在utf-8编码中中文字符占用3字节所以len(s)是10
	fmt.Println("s len:", len(s))
	//s[i]也是通过一字节一字节访问字符串的
	for i := 0; i < len(s); i++ {
		fmt.Printf("s:%c code:%d type:%T\n", s[i], s[i], s[i])
	}
	//range s是通过rune类型为单位来遍历字符串的数据的
	// rune会根据编码方式读取正确的字符
	for _, c := range s {
		fmt.Printf("rune:%c code:%d type:%T\n", c, c, c)
	}
}
