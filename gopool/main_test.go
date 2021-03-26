// Package main provides ...
package main

import "testing"

var runTimes = 1000000
var sum = Sum{a: 10, b: 10}
var capacity = 1000000

func BenchmarkOneGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum.Add()
	}
}

func BenchmarkGoroutinesSetTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go sum.Add()
	}
}

func BenchmarkPutSetTimes(b *testing.B) {
	p, _ := NewPool(capacity)
	// go func(){
	// 	p.Put(w)
	// }()

	go func() {
		for i := 0; i < b.N; i++ {
			p.Put(sum)
		}
	}()
	// for i := 0; i < runTimes; i++ {
	p.Run()
	// }
}
