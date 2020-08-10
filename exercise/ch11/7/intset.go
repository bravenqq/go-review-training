// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

const bit = 32 << (^uint(0) >> 63)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

func NewIntSet() *IntSet {
	return &IntSet{}
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bit, uint(x%bit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bit, uint(x%bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bit*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

//!+len
//Len return the number of element
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		for j := 0; j < bit; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

//!-len

//!+remove
//Remove remove x from the set
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/bit, x%bit
		s.words[word] = s.words[word] & ^(1 << uint(bit))
	}
}

//!-remove

//!+clear
//Clear remove all elements from the set
func (s *IntSet) Clear() {
	for i, word := range s.words {
		s.words[i] = word & 0
	}
}

//!-clear

//!+copy
//Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var is IntSet
	is.words = make([]uint, 0, len(s.words))
	for _, word := range s.words {
		is.words = append(is.words, word)
	}
	return &is
}

//!-copy

//!+AddAll
//AddAll add all nums
func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

//!-AddAll

//IntersectWith 交集：元素在A集合B集合均出现
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, word := range s.words {
		if i < len(t.words) {
			s.words[i] = word & t.words[i]
		} else {
			s.words[i] = s.words[i] & 0
		}
	}
}

//DifferenceWith “差集：元素出现在A集合，未出现在B集合”
func (s *IntSet) DifferenceWith(t *IntSet) {
	//先求交集
	t.IntersectWith(s)
	//再用交集异或
	for i, word := range t.words {
		s.words[i] = s.words[i] ^ word
	}
}

//SymmetricDifference “并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A”
func (s *IntSet) SymmetricDifference(t *IntSet) {
	tmp := s.Copy()
	//先求并集
	s.UnionWith(t)
	//再求交集
	t.IntersectWith(tmp)
	//再求差集
	s.DifferenceWith(t)
}

//Elems 返回集合中所有元素
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < bit; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, i*bit+j)
			}
		}
	}
	return elems
}
