// Package main provides ...
package main

/*
取出数组中间隔元素
*/

func main() {
	arr1 := []int{20, 8, 7, 30, 50}
	arr2 := make([]int, len(arr1)/2+1)
	var j int
	//方法一：遍历arr1数组,判断i为偶数插入到arr2中
	//操作：比较+读取+插入数组
	//比较：n次
	//读取arr1中元素：n/2
	//插入元素到arr2：n/2次 共2*n 时间复杂度O(n)
	for i := 0; i < len(arr1); i++ {
		if i%2 == 0 {
			arr2[j] = arr1[i]
			j++
		}
	}

	j = 0
	//方法二：直接读取数组中间隔元素
	//操作:
	//1.读取arr1中元素：n/2
	//2.插入元素到arr2：n/2次 共n次时间复杂度O(n),但是第一种方式是2n,比第一种方式快1倍
	for i := 0; i < len(arr1); i += 2 {
		arr2[j] = arr1[i]
		j++
	}

}
