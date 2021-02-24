package main

import (
	"fmt"
	"testing"
)

func main() {
	r := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo(34)
		}
	})
	fmt.Println(r.String(), r.MemString())
}

func foo(x int) {
	add(func() int {
		return x + 1
	}, func() int {
		return x + 2
	}, T{x}.L)
}

type T struct {
	x int
}

func (t T) L() int {
	return t.x + 9
}

func add(args ...func() int) int {
	x := 0
	for _, a := range args {
		x += a()
	}
	return x
}
