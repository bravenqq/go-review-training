package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

const PI = 3.14

func show(i interface{}) {
	fmt.Println(i)
}

// heap
// stack
func main() {
	// //x在计算机中的存储形式补码，值为：11111111
	// var x int8 = -1
	// //右移动1位值为：11111111，补码值，转换为原码值为：10000001即10进制-1
	// y := x >> 1
	// //左移之后的存储值为11111100,原码值为：10000100即10进制-4
	// z := x << 2
	// fmt.Printf("%d\n", y)
	// fmt.Println(z)
	// var a uint = 64
	// show(a)

	//i类型uint, 值一直非负数，导致死循环
	// for i := a; i >= 0; i-- {
	// 	fmt.Println(i)
	// }
	var a, b int32
	a = math.MaxInt32
	b = 100
	sum, err := AddSafety32(a, b)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sum)
	a = math.MinInt32
	b = -100
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sum)
}

var ErrOverFlow = errors.New("数据溢出")

func AddSafety32(a, b int32) (int32, error) {
	if a > 0 {
		if a > math.MaxInt32-b {
			return 0, ErrOverFlow
		}
	}

	if a < 0 {
		if a < math.MinInt32-b {
			return 0, ErrOverFlow
		}
	}
	return a + b, nil
}
