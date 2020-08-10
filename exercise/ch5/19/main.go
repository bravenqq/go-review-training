/*
“使用panic和recover编写一个不包含return语句但能返回一个非零值的函数”
*/
package main

import "fmt"

func main() {
	fmt.Println(NoZeroReturn())
}

func NoZeroReturn() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 1
		}
	}()
	panic(1)
}
