package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	p := make(chan string)
	q := make(chan string)
	go func() {
		p <- "ping"
		for {
			count++
			p <- <-q
		}
	}()
	go func() {
		for {
			q <- <-p
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("count:", count)
}
