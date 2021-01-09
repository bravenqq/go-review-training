// Package main provides ...
package main

import "fmt"

func main() {
	arr := []int{0, 5, 2, 1, 6, 3}
	Sort(arr)
	fmt.Printf("%+v\n", arr)
}

func Sort(arr []int) {
	if len(arr) == 1 {
		return
	}
	leftPoint := 0
	rightPoint := len(arr) - 2
	pivot := len(arr) - 1

	for {
		// 查找左边比轴大的值
		for leftPoint <= rightPoint {
			if arr[leftPoint] >= arr[pivot] {
				break
			}
			leftPoint++
		}

		// 查找右侧比轴小的值
		for rightPoint > leftPoint {
			if arr[rightPoint] <= arr[pivot] {
				break
			}
			rightPoint--
		}

		// 当leftPoint == rightPoint 停止
		if leftPoint >= rightPoint {
			break
		}

		// 交换轴
		swap(arr, leftPoint, rightPoint)
	}

	swap(arr, leftPoint, pivot)
	Sort(arr[:leftPoint])
	Sort(arr[leftPoint+1:])
}

func swap(arr []int, left, right int) {

	temp := arr[left]
	arr[left] = arr[right]
	arr[right] = temp
}
