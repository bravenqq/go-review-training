// Package main provides ...
package main

import "fmt"

/*
* 插入排序法
时间复杂度分析：最差情况是降序，每次都要操作所有步骤
操作：
1.移除:操作n-1步
2.比较:如果完全逆序需要比较1+2+3+...(n-1)=n^2/2次
3.平移:如果完全逆序每次比较需要平移:1+2+3+...(n-1)=n^2/2
4.插入:操作n-1
最差情况总共操作：n^2+2n-2---->O(n^2+n)时间复杂度--随着数据的海量增长n^2越来越抛离线性O(n)时间复杂度，保留最高阶层--->O(n^2)
*/

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
		//插入
		if position != i {
			arr[position] = temp
		}

	}
}
