// Package main provides ...
package main

import "fmt"

/*
* 求数组中是否存在重复的元素
* 方式一：双层循环，第一层循环取i指向的元素a，第二层循环用j来遍历素组，判断j执行位置值与a是否相同，且i与j的位置不同，
* 操作步骤：比较值
* 时间复杂度：最差情况数组中没有重复元素，i,j要遍历完内外两层循环，要走n^2步，时间复杂度为o(n^2)意味着低效
* 方式二：使用循环取i指向的值a，已值a为key保存到一个map,每次取下一个元素时判断map中是否存在已a为key的元素，有说明存在重复元素。
* 操作步骤：1、从map中取值 2、保存值到map中
* 时间复杂度：O(n)比第一种方式O(n^2)时间复杂度更高效，但是创建了一个map来保存数组中的值，消耗更多空间
*
 */

type Duplicater interface {
	Len() int
	ElementEqual(i, j int) bool
}

func (a Array) Len() int {
	return len(a)
}

func (a Array) ElementEqual(i, j int) bool {
	return a[i] == a[j]
}

type Array []int

func main() {
	arr := Array{12, 30, 15, 20, 10, 10}
	fmt.Println(HasDuplicateValue(arr))
	fmt.Println(HasDuplicateVal(arr))

}

func HasDuplicateVal(arr Array) bool {
	st := make(map[int]struct{})
	for i := 0; i < arr.Len(); i++ {
		if _, ok := st[arr[i]]; ok {
			return true
		}
		st[arr[i]] = struct{}{}
	}
	return false
}

func HasDuplicateValue(arr Duplicater) bool {
	for i := 0; i < arr.Len(); i++ {
		for j := 0; j < arr.Len(); j++ {
			if arr.ElementEqual(i, j) && i != j {
				return true
			}
		}
	}
	return false
}
