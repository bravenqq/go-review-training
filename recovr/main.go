// Package main provides ...
package main

import (
	"fmt"
)

func main() {

	defer func() {
		v := recover()
		fmt.Println(v)
	}()
	defer func() {
		fmt.Println("hello,world")
	}()
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
		fmt.Println("test")
		panic("test panic")
	}()
	<-ch
}
