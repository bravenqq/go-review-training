package intset

import (
	"math/rand"
	"testing"
)

func TestMapAdd(t *testing.T) {
	set := NewMapIntSet()
	tests := []struct {
		input  int
		output string
	}{
		{10, "{10}"},
	}
	for _, test := range tests {
		set.Add(test.input)
		out := set.String()
		if out != test.output {
			t.Errorf("MapIntSet Add input:%d want:%s get:%s\n", test.input, test.output, out)
		}
	}
}

func BenchmarkMap10(b *testing.B) {
	benchmarkMapAdd(b, 10)
}
func BenchmarkMap100(b *testing.B) {
	benchmarkMapAdd(b, 100)
}
func BenchmarkMap1000(b *testing.B) {
	benchmarkMapAdd(b, 1000)
}
func BenchmarkMap10000(b *testing.B) {
	benchmarkMapAdd(b, 10000)
}

func benchmarkMapAdd(b *testing.B, n int) {
	set := NewMapIntSet()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
	}
}
