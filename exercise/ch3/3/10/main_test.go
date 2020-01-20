package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		input string
		ouput string
	}{
		{input: "+12345678.1234", ouput: "+12,345,678.1234"},
		{input: "123456789", ouput: "123,456,789"},
	}

	for _, test := range tests {
		res := Comma(test.input)
		if res != test.ouput {
			t.Errorf("Comma input %s get:%s want:%s", test.input, res, test.ouput)
		}
	}

}

func TestCommaFast(t *testing.T) {
	tests := []struct {
		input string
		ouput string
	}{
		{input: "12345678", ouput: "12,345,678"},
		{input: "123456789", ouput: "123,456,789"},
	}

	for _, test := range tests {
		res := CommaFast(test.input)
		if res != test.ouput {
			t.Errorf("Comma input %s get:%s want:%s", test.input, res, test.ouput)
		}
	}

}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma("12345678")
	}
}

func BenchmarkCommaFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommaFast("123456789134444563576567765")
	}
}
