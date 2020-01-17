// Package main provides ...
package main

import (
	"testing"
)

func TestDirectory(t *testing.T) {
	tests := []struct {
		input in
		want  int
	}{
		{input: in{data: []int{35, 38, 40, 55}, num: 40}, want: 2},
	}

	for _, test := range tests {
		index, err := DirectorySearch(test.input.data, test.input.num)
		if err != nil {
			t.Error(err)
		}
		if index != test.want {
			t.Errorf("input:%v find:%d res:%d want:%d", test.input.data, test.input.num, index, test.want)
		}
	}
}

func BenchmarkSortInsert100(b *testing.B) {
	data := make([]int, 0, 100)
	for i := 0; i < b.N; i++ {
		SortInsert(data, 100)
	}
}

func BenchmarkSortInsert1000(b *testing.B) {
	data := make([]int, 0, 1000)
	for i := 0; i < b.N; i++ {
		SortInsert(data, 1000)
	}
}

func BenchmarkSortInsert100000(b *testing.B) {
	data := make([]int, 0, 100000)
	for i := 0; i < b.N; i++ {
		SortInsert(data, 100000)
	}
}

func BenchmarkSortInsert360000(b *testing.B) {
	data := make([]int, 0, 360000)
	for i := 0; i < b.N; i++ {
		SortInsert(data, 360000)
	}
}

func BenchmarkSearch(b *testing.B) {
	data := make([]int, 100000)
	data = SortInsert(data, 100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(data, 10)
	}
}

func BenchmarkDirectorySearch(b *testing.B) {

	data := make([]int, 100000)
	data = SortInsert(data, 100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DirectorySearch(data, 10)
	}
}

type in struct {
	data []int
	num  int
}
