package main

import (
	"testing"
)

func TestRemoveExtraSpace(t *testing.T) {

	tests := []Test{Test{input: []byte("聂倩倩   def  ddc "), output: []byte("聂倩倩 def ddc ")}}
	for _, test := range tests {
		res := RemoveExtraSpace(test.input)
		if !test.equal(res) {
			t.Errorf("Deduplicate get:%s want:%s\n", string(test.input), string(test.output))
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
