// Package popcount provides ...
package popcount

import (
	"fmt"
	"testing"
)

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(32767)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(32767)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(32767)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(32767)
	}
}

func TestPopCount1(t *testing.T) {
	fmt.Println(PopCount1(32767))
	fmt.Println(PopCount3(32767))
	fmt.Println(PopCount4(32767))
}
