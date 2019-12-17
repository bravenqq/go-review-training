// Package main provides ...
package main

import "fmt"

func main() {

	var a int8 = 10
	var b int16 = 255

	//b转为int8类型发生了溢出
	c := a + int8(b)
	fmt.Printf("c %d\n", c)

	d := a / 3
	fmt.Printf("d %d\n", d)

	//2.0为无类型浮点数，如果2.0转换为整形不会发绳截断则会被隐式转换为int8类型
	f := a / 2.0
	fmt.Printf("f %d\n", f)

	e := 5.5 / 3
	fmt.Printf("e %f\n", e)

}
