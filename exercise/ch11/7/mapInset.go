// Package intset provides ...
package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	words map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{words: make(map[int]bool)}
}

func (set *MapIntSet) Len() int {
	return len(set.words)
}

func (set *MapIntSet) String() string {
	values := make([]int, len(set.words))
	i := 0
	for k, v := range set.words {
		if v == true {
			values[i] = k
			i++
		}
	}
	sort.Ints(values)
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range values {
		fmt.Fprintf(&buf, "%d", word)
		if i != len(values)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (set *MapIntSet) Add(x int) {
	set.words[x] = true
}

func (set *MapIntSet) Has(x int) bool {
	return set.words[x]
}

func (set *MapIntSet) Remove(x int) {
	delete(set.words, x)
}

func (set *MapIntSet) Clear() {
	for k := range set.words {
		delete(set.words, k)
	}
}

func (set *MapIntSet) Copy() *MapIntSet {
	t := &MapIntSet{make(map[int]bool)}
	for k, v := range set.words {
		t.words[k] = v
	}
	return t
}

func (set *MapIntSet) AddAll(nums ...int) {
	for _, num := range nums {
		set.words[num] = true
	}
}

func (set *MapIntSet) UnionWith(t *MapIntSet) {
	for k, v := range t.words {
		if v == true {
			set.words[k] = v
		}
	}
}

func (set *MapIntSet) IntersectWith(t *MapIntSet) *MapIntSet {
	m := &MapIntSet{make(map[int]bool)}
	for k, v := range set.words {
		if t.words[k] == true && v == true {
			m.words[k] = v
		}
	}
	return m
}

func (set *MapIntSet) DifferenceWith(t *MapIntSet) *MapIntSet {
	m := &MapIntSet{make(map[int]bool)}
	for k, v := range set.words {
		if t.words[k] == false && v == true {
			m.words[k] = v
		}
	}
	return m
}

func (set *MapIntSet) SymmetricDifference(t *MapIntSet) *MapIntSet {
	s := set.IntersectWith(t)
	set.UnionWith(t)
	m := &MapIntSet{make(map[int]bool)}
	for k, v := range set.words {
		if s.words[k] == false {
			m.words[k] = v
		}
	}
	return m
}
