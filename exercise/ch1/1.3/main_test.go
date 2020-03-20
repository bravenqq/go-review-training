package main

import (
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"s", "b"})
	}
}

func BenchmarkFastJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastJoin([]string{"s", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q"})
	}
}
