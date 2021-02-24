// Package main provides ...
package main

func main() {
	var a = 1
	go func() {
		//使用的是main中a的地址来修改a变量
		a++
	}()
}
