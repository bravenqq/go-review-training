/*
“编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。”
*/

// Package main provides ...
package main

import "sort"

func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

//IsPalindrome if s is a palindrome
func IsPalindrome(s sort.Interface) bool {
	//check if the i and s.Len()-1-i is equal
	max := s.Len() - 1
	for i := 0; i < max/2; i++ {
		if !equal(i, max-1, s) {
			return false
		}
	}
	return true
}
