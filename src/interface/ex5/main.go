// Package main provides ...
package main

import (
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("name:%s age:%d", p.Name, p.Age)
}

func main() {
	//此处接口初始化，接口动态类型：nil,动态值：nil
	var w io.Writer
	fmt.Printf("interface 动态类型：%T,动态值：%v\n", w, w) //interface 动态类型：<nil>,动态值：<nil>
	//1. 接口的动态类型：*os.File,动态值：os.Stdout
	//2. 此处分配了一块内存保存os.File指针中的值,接口的动态值指向这块内存
	w = os.Stdout
	fmt.Printf("stdout address:%p\n", os.Stdout)          //stdout address:0xc00000e018
	fmt.Printf("stdout pointer address:%p\n", &os.Stdout) //stdout pointer address:0x1161958
	fmt.Printf("w 动态类型：%T,动态值：%v\n", w, w)                //w 动态类型：*os.File,动态值：&{0xc00004c0c0}
	f := w.(*os.File)
	fmt.Printf("f address:%p\n", f) //f address:0xc00000e018
	fmt.Println(w == os.Stdout)     //true, 由于w也指向了os.Stdout指向的地址

	w = nil
	fmt.Printf("interface 动态类型：%T,动态值：%v\n", w, w)
	fmt.Println(w == os.Stdout) //false,w的动态值是nil
}
