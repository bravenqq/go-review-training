// Package main provides ...
package main

import "fmt"

/*
选择排序
*/

//Sorter 排序接口
type Sorter interface {
	Len() int
	Compare(i, j int) bool
	Swap(i, j int)
}

//Array ....
type Array []int

//Len ....
func (arr Array) Len() int {
	return len(arr)
}

//Compare .....
func (arr Array) Compare(i, j int) bool {
	return arr[j] < arr[i]
}

//Swap .....
func (arr Array) Swap(i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	array := Array{10, 20, 15, 30, 22, 43, 33}
	sort(array)
	fmt.Println(array)
}

func sort(arr Sorter) {
	for i := 0; i < arr.Len(); i++ {
		minIndex := i
		for j := i + 1; j < arr.Len(); j++ {
			if arr.Compare(minIndex, j) {
				minIndex = j
			}
		}
		if i != minIndex {
			arr.Swap(i, minIndex)
		}
	}
}
