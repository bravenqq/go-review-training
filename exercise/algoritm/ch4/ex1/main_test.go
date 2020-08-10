package main

import (
	"math/rand"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		input *List
		ouput int
	}{
		{create(10), 10},
	}
	for _, test := range tests {
		res := Count(test.input)
		if res != test.ouput {
			t.Errorf("Count List get:%d want:%d\n", res, test.ouput)
		}
	}
}

func TestFindMax(t *testing.T) {
	l := New()
	l.Add(2)
	l.Add(5)
	l.Add(6)
	l.Add(4)
	tests := []struct {
		input *List
		ouput int
	}{
		{l, 6},
	}
	for _, test := range tests {
		res := FindMax(test.input.Data, test.input.Next)
		if res != test.ouput {
			t.Errorf("FindMax List get:%d want:%d\n", res, test.ouput)
		}
	}
}

func New() *List {
	return &List{Data: 0, Next: nil}
}

func (l *List) Add(data int) {
	next := List{data, nil}
	temp := l
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &next
}

func create(n int) *List {
	l := &List{Data: 0, Next: nil}
	head := l
	for i := 1; i < n; i++ {
		ls := List{Data: rand.Int() % 100, Next: nil}
		l.Next = &ls
		l = &ls
	}
	return head
}
