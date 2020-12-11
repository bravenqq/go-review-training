// Package main provides ...
package main

import "testing"

var q *Queue

func init() {
	q = NewQueue()
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q.Join("test")
	}
}

func BenchmarkLeave(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		q.Join("test")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.Leave()
	}
}
