// Package main provides ...
package main

import "testing"

var s *Stack

func init() {
	s = NewStack()
}

func BenchmarkPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Push("test")
	}
}

func BenchmarkPop(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s.Push("test")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
