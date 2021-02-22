package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
func main() {
	// fmt.Println(comma("123456"))
	str := "你好，中国"
	for _, s := range str {
		fmt.Printf("%d ", s)
	}
	// fmt.Println()
	// for _, s := range str {
	// 	fmt.Printf("%s ", s)
	// /  }
	fmt.Println()
	bytes := []byte(str)
	fmt.Println(bytes)
	// runes := []rune(str)
	// fmt.Println(string(runes[:3]))
	// str := "hello,world"
	// c := str
	// fmt.Printf("str:%p str:%p\n", &str, &c)
	// Out(c)
	// c += "!"
	// Out(c)
	//
	// c += "!"
	// fmt.Printf("%p %p\n", &str, &c)
	//
	// subStr := str[:5]
	// fmt.Printf("subStr %s, str %s\n", subStr, str)
	//
	// str = "hi,world"
	// fmt.Printf("subStr %s, str %s\n", subStr, str)
	//
	// subStr = "hi,nqq"
	// fmt.Printf("subStr %s, str %s\n", subStr, str)
	//
	// str = string([]byte{2, 3})
	// fmt.Println(str)
	//改变字符串底层值，报错cannot assign to subStr[:5]
	//subStr[:5] = "hi"
}

func Out(s string) {
	fmt.Printf("s:%p", &s)
}
