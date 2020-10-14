// Package main provides ...
package main

/*bubble sort
* 需要做两个步骤
*  1.比较
*  2.交换
* 比较所操作步骤: n-1 + n-2 + n-3 + n-4 +...+1 ---> n(n-1)/2
* 交换的最差情况如果是从大到小的顺序需要交换操作的步骤：n-1 + n-2 + n-3 + n-4 + ... + 1 --->n(n-1)/2
* bubble sort 排序需要操作步骤：n(n-1)/2*2=n(n-1)=n^2-2n ~ O(n^2)
 */

import "fmt"

type Sorter interface {
	Len() int
	Compare(a, b int) bool
	Swap(a, b int)
}

type Array []int

func (a Array) Len() int {
	return len(a)
}

func (arr Array) Compare(i, j int) bool {
	return arr[i] > arr[j]
}

func (arr Array) Swap(i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	arr := Array{20, 50, 15, 60, 90, 80}
	fmt.Printf("before:%+v\n", arr)
	Sort(arr)
	fmt.Printf("after:%+v\n", arr)

}

func Sort(arr Array) {
	for i := 0; i < arr.Len(); i++ {
		for j := 0; j < arr.Len()-i-1; j++ {
			if arr.Compare(j, j+1) {
				arr.Swap(j, j+1)
			}
		}
	}
}
