// Package main provides ...
package main

import (
	"testing"
)

func TestMbreverse(t *testing.T) {

	tests := []Test{Test{input: []byte("聂倩倩"), output: []byte("倩倩聂")}}
	for _, test := range tests {
		res := Mbreverse(test.input)
		if !test.equal(res) {
			t.Errorf("Mbreverse get:%s want:%s\n", res, string(test.output))
		}
	}
}
func (t Test) equal(res []byte) bool {
	if len(res) != len(t.output) {
		return false
	}
	for i := 0; i < len(res); i++ {
		if res[i] != t.output[i] {
			return false
		}
	}
	return true

}

type Test struct {
	input  []byte
	output []byte
}
