package main

import (
	"flag"
	"fmt"
)

func main() {
	index := flag.Int("index", 3, "please input index value")
	flag.Parse()
	arr := []int{0, 1, 2}
	fmt.Println(arr[*index])
}
