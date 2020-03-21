package intset

import (
	"math/rand"
	"testing"
)

func TestLen(t *testing.T) {
	tests := []struct {
		input  *IntSet
		output int
	}{
		{&IntSet{words: []uint{10, 11}}, 5},
	}
	for _, test := range tests {
		res := test.input.Len()
		if res != 5 {
			t.Errorf("Len input:%v get:%d want:%d", test.input, res, test.output)
		}
	}
}

func TestRemove(t *testing.T) {
	s := &IntSet{}
	s.Add(3)
	s.Add(4)
	s.Add(2000)
	tests := []struct {
		input1 *IntSet
		input2 int
		output string
	}{
		{s, 3, "{4 2000}"},
	}
	for _, test := range tests {
		test.input1.Remove(test.input2)
		if s.String() != test.output {
			t.Errorf("Remove input:%s get:%s want:%s", "{3 4 2000}", test.input1.String(), test.output)
		}
	}
}

func TestClear(t *testing.T) {

	s := &IntSet{}
	s.Add(3)
	s.Add(4)
	s.Add(2000)
	tests := []struct {
		input  *IntSet
		output string
	}{
		{s, "{}"},
	}
	for _, test := range tests {
		test.input.Clear()
		if s.String() != test.output {
			t.Errorf("Clear input:%s get:%s want:%s", "{3 4 2000}", test.input.String(), test.output)
		}
	}
}

func TestCopy(t *testing.T) {
	s1 := &IntSet{}
	s1.Add(3)
	s1.Add(4)
	s1.Add(2000)
	s2 := &IntSet{}
	s2.Add(3)
	s2.Add(4)
	s2.Add(2000)
	tests := []struct {
		input  *IntSet
		output *IntSet
	}{
		{s1, s2},
	}
	for _, test := range tests {
		test.input.Copy()
		if s1.String() != test.output.String() {
			t.Errorf("Copy input:%s get:%s want:%s", "{3 4 2000}", test.input.String(), test.output.String())
		}
	}
}

func TestAddAll(t *testing.T) {
	s := &IntSet{}
	tests := []struct {
		input1 *IntSet
		input2 []int
		output string
	}{
		{s, []int{1, 2, 3}, "{1 2 3}"},
	}
	for _, test := range tests {
		test.input1.AddAll(test.input2...)
		if test.input1.String() != test.output {
			t.Errorf("AddAll input:%s get:%s want:%s", "{1 2 3}", test.input1.String(), test.output)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	s1 := &IntSet{}
	s1.AddAll(1, 3, 80, 10)
	s2 := &IntSet{}
	s2.AddAll(3, 5, 6, 10, 20)
	s3 := &IntSet{}
	s3.AddAll(3, 10)
	tests := []struct {
		input1 *IntSet
		input2 *IntSet
		output *IntSet
	}{
		{s1, s2, s3},
	}
	for _, test := range tests {
		tmp := test.input1.Copy()
		test.input1.IntersectWith(test.input2)
		if test.input1.String() != test.output.String() {
			t.Errorf("IntersectWith input:%s get:%s want:%s\n", tmp.String(), test.input1.String(), test.output.String())
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	s1 := &IntSet{}
	s1.AddAll(1, 3, 80, 10)
	s2 := &IntSet{}
	s2.AddAll(3, 5, 6, 10, 20)
	s3 := &IntSet{}
	s3.AddAll(1, 80)
	tests := []struct {
		input1 *IntSet
		input2 *IntSet
		output *IntSet
	}{
		{s1, s2, s3},
	}
	for _, test := range tests {
		tmp := test.input1.Copy()
		test.input1.DifferenceWith(test.input2)
		if test.input1.String() != test.output.String() {
			t.Errorf("DifferenceWith input:%s get:%s want:%s\n", tmp.String(), test.input1.String(), test.output.String())
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	s1 := &IntSet{}
	s1.AddAll(1, 3, 80, 10)
	s2 := &IntSet{}
	s2.AddAll(3, 5, 6, 10, 20)
	s3 := &IntSet{}
	s3.AddAll(1, 80, 5, 6, 20)
	tests := []struct {
		input1 *IntSet
		input2 *IntSet
		output *IntSet
	}{
		{s1, s2, s3},
	}
	for _, test := range tests {
		tmp := test.input1.Copy()
		test.input1.SymmetricDifference(test.input2)
		if test.input1.String() != test.output.String() {
			t.Errorf("SymmetricDifference input:%s get:%s want:%s\n", tmp.String(), test.input1.String(), test.output.String())
		}
	}
}

func TestElems(t *testing.T) {
	s := &IntSet{}
	s.AddAll(2, 4, 10, 8, 90)
	tests := []struct {
		input  *IntSet
		output []int
	}{
		{s, []int{2, 4, 8, 10, 90}},
	}
	for _, test := range tests {
		res := test.input.Elems()
		if !equal(res, test.output) {
			t.Errorf("Elems input:%s get:%v want:%v", test.input.String(), res, test.output)
		}
	}
}

func equal(res, out []int) bool {
	if len(res) != len(out) {
		return false
	}
	for i, v := range res {
		if v != out[i] {
			return false
		}
	}
	return true
}

var max = 3200

func benchmarkAdd(b *testing.B, n int) {
	set := NewIntSet()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
	}
}

func Benchmark10(b *testing.B) {
	benchmarkAdd(b, 10)
}
func Benchmark100(b *testing.B) {
	benchmarkAdd(b, 100)
}
func Benchmark1000(b *testing.B) {
	benchmarkAdd(b, 1000)
}
func Benchmark10000(b *testing.B) {
	benchmarkAdd(b, 10000)
}
