// Package main provides ...
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//1. 使用context 取消goroutine的执行
//2. 使用context实现超时取消
func main() {
	tasks := make(chan int)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
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
	time.Sleep(10 * time.Second)
}
