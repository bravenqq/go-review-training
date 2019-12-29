package main

import "fmt"

func main() {
	//arr1初始化为对应成员的零值
	var arr1 [10]int
	fmt.Println("arr1:", arr1)

	arr2 := [...]int{1, 19, 20, 3}
	fmt.Printf("arr2 type:%T value:%v\n", arr2, arr2)

	//数组长度为100
	arr3 := [...]int{99: 1001}
	fmt.Printf("arr3 type:%T value:%v\n", arr3, arr3)

}
