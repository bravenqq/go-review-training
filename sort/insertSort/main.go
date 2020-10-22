// Package main provides ...
package main

import "fmt"

func main() {
	arr := []int{10, 30, 40, 20, 25, 35, 70, 60}
	sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		//移除
		temp := arr[i]
		position := i
		for position > 0 && arr[position-1] > temp {
			//比较
			//平移
			arr[position] = arr[position-1]
			position--
		}
		arr[position] = temp

	}
}
