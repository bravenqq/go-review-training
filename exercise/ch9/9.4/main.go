// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	count := flag.Int("count", 100000, "please inout goroutien count")
	flag.Parse()
	start := time.Now()
	in, out := pipline(*count)
	in <- 1
	<-out
	close(out)
	fmt.Println("time:", time.Since(start))
}

func pipline(stages int) (chan int, chan int) {
	out := make(chan int)
	first := out
	for i := 0; i < stages; i++ {
		in := out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(in)
		}(in, out)
	}
	return first, out
}
