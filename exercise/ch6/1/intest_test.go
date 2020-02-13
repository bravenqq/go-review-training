package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	tests := []struct {
		input  *IntSet
		output int
	}{
		{&IntSet{words: []uint64{10, 11}}, 5},
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
			t.Errorf("Len input:%s get:%s want:%s", "{3 4 2000}", test.input1.String(), test.output)
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
			t.Errorf("Len input:%s get:%s want:%s", "{3 4 2000}", test.input.String(), test.output)
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
			t.Errorf("Len input:%s get:%s want:%s", "{3 4 2000}", test.input.String(), test.output.String())
		}
	}
}
