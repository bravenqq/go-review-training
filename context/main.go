// Package main provides ...
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//1. 使用context 取消goroutine的执行
func main() {
	tasks := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case t := <-tasks:
					fmt.Println(t)
				case <-ctx.Done():
					fmt.Println("done")
					return
				}
			}
		}(ctx)
	}

	go func() {
		for {
			tasks <- rand.Intn(100)
		}
	}()
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)

}
