package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr)
	test(arr)
	fmt.Println(arr)

	var p *int
	p = &arr[1]
	fmt.Println(*p)
	*p = 60
	fmt.Println(arr)

	var b [10]*int
	a := 10
	b[0] = &a

	s := "hello,world"
	q := &s
	*q = "nqq"
	fmt.Println(s)

	t := test
}

func test(arr [10]int) {
	arr[0] = 10
}
