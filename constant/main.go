package main

import "fmt"

const (
	name string = "nieqianqian"
)

const (
	_ = 1 << (10*iota + 1)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB //1208925819614629174706176 值超出了计算机任何基础类型的范围，但是这里是无类型常量可以使用
)

func main() {
	var arr [10]int
	fmt.Println(YiB / ZiB)
	const length = len(arr)
}

func count(arr []int) int {
	return len(arr)
}
