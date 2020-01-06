// Package strings provides ...
package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	f, err := os.Open("/Users/abbynie/tmp/article")
	defer f.Close()
	if err != nil {
		b.Error(err)
	}
	data, err := ioutil.ReadAll(f)
	slice := strings.Split(string(data), " ")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(slice, " ")
	}
}

func BenchmarkFastJoin(b *testing.B) {

	f, err := os.Open("/Users/abbynie/tmp/article")
	defer f.Close()
	if err != nil {
		b.Error(err)
	}
	data, err := ioutil.ReadAll(f)
	slice := strings.Split(string(data), " ")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastJoin(slice, " ")
	}
}
