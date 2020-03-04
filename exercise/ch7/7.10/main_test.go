package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := Palindrome("aabaa")
	if !IsPalindrome(Palindrome([]byte(s))) {
		t.Errorf("IsPalindrome is err")
	}
}

type Palindrome []byte

func (p Palindrome) Len() int {
	return len(p)
}

func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Palindrome) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
