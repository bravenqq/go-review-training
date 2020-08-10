// Package main provides ...
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	defer wg.Wait()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("A:", i)
			defer wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("B:", i)
			defer wg.Done()
		}(i)
	}
}
