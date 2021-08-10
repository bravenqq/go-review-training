// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	// 关键人物出场
	// m := sync.Map{}
	// m.Store(1, 1)
	m := make(map[int]int)
	go do(m)
	go do(m)

	time.Sleep(1 * time.Second)
	_, ok := m[1]
	fmt.Println(ok)
}

func do(m map[int]int) {
	i := 0
	for i < 10000 {
		// m.Store(1, 1)
		m[1] = 1
		i++
	}
}
