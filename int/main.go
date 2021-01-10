package main

import "fmt"

const PI = 3.14

func show(i interface{}) {
	fmt.Println(i)
}

// heap
// stack
func main() {
	//x在计算机中的存储形式补码，值为：11111111
	var x int8 = -1
	//右移动1位值为：11111111，补码值，转换为原码值为：10000001即10进制-1
	y := x >> 1
	//左移之后的存储值为11111100,原码值为：10000100即10进制-4
	z := x << 2
	fmt.Printf("%d\n", y)
	fmt.Println(z)
	var a uint = 64
	show(a)

	//i类型uint, 值一直非负数，导致死循环
	// for i := a; i >= 0; i-- {
	// 	fmt.Println(i)
	// }
}
