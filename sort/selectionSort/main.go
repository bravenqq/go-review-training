// Package main provides ...
package main

import "fmt"

/*
选择排序
1.选择排序的操作步骤：比较和交换
2.最差情况下假设数组逆序，
> 长度为n的数组比较次数:n-1+n-2+n-3+...+1 --->总共需要n*(n-1)/2=(n^2-n)/2,时间复杂度O(n^2)
> 长度为n的数组交换次数，最差情况逆序每次内部循环比较完都要交换，交换n次，时间复杂度为O(n)
> 所以选择排序的时间复杂度为O(n^2)+O(n)
> 由于冒泡排序每次比较都要交换n^2/2次，而选择排序只需要交换n次，所以选择排序几乎比冒牌排序快一倍，虽然时间复杂度都为O(n^2)
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
			//比较
			if arr.Compare(minIndex, j) {
				minIndex = j
			}
		}
		if i != minIndex {
			//交换
			arr.Swap(i, minIndex)
		}
	}
}
