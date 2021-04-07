// Package main provides ...
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	sum := 0
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
			sum++
		}(i)
	}

	time.Sleep(15 * time.Second)
	fmt.Println(sum)
}
