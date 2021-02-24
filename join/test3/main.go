package main

import (
	"fmt"
	"testing"
)

func main() {
	r := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo(2343243)
		}
	})
	fmt.Println(r.String(), r.MemString())
}

func foo(x int) {
	add(F(func() int {
		return x + 1
	}), F(func() int {
		return x + 2
	}), T{x})
}

type F func() int

func (f F) L() int {
	return f()
}

type T struct {
	x int
}

func (t T) L() int {
	return t.x + 9
}

func (t T) With(w func() int) func() int {
	return func() int {
		return t.x + w()
	}
}

type Ler interface {
	L() int
}

func add(args ...Ler) int {
	x := 0
	for _, a := range args {
		x += a.L()
	}
	return x
}
