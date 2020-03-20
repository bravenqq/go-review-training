package main

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCountRune(t *testing.T) {
	tests := []struct {
		input   string
		counts  map[string]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{input: "聂倩倩，你好", counts: map[string]int{"聂": 1, "倩": 2, "，": 1, "你": 1, "好": 1}, utflen: [5]int{3: 6}, invalid: 0},
	}

	for _, test := range tests {
		counts, utflen, invalid := countRune(strings.NewReader(test.input))
		if !mapEqual(counts, test.counts) {
			t.Errorf("countRune input:%s want counts:%+v get counts:%+v", test.input, test.counts, counts)
		}
		if !arrEqual(test.utflen, utflen) {
			t.Errorf("countRune input:%s want utflen:%+v get utflen:%+v", test.input, test.utflen, utflen)
		}
		if invalid != invalid {
			t.Errorf("countRune input:%s want invalid:%d get invalid:%d", test.input, test.invalid, invalid)
		}
	}
}

func mapEqual(map1 map[rune]int, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k, v := range map2 {
		key, _ := utf8.DecodeRuneInString(k)
		if map1[key] != v {
			return false
		}
	}
	return true
}

func arrEqual(arr1, arr2 [utf8.UTFMax + 1]int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
