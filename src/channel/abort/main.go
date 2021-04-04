// Package main provides ...
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	tasks := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				select {
				case task := <-tasks:
					fmt.Println(task)
				case <-done:
					fmt.Println("exit")
					wg.Done()
					return
				}
			}
		}(i)
	}

	go func() {
		for {
			tasks <- rand.Intn(100)
		}
	}()
	time.Sleep(5 * time.Second)
	close(done)
	wg.Wait()
}
