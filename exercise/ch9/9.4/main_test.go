// Package main provides ...
package main

import (
	"testing"
)

func BenckmarkPipline(b *testing.B) {
	in, out := pipline(100000000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}
