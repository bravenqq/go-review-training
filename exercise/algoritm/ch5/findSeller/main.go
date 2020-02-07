/*
假设你经营着一个芒果农场，需要寻找芒果销售商，以便将芒果卖给他。在Facebook，你与芒果销售商有联系吗？为此，你可在朋友中查找关系最近的芒果商
*/

package main

import (
	"fmt"
	"strings"
)

var personNets = map[string][]string{
	"you":    {"alice", "bob", "claire"},
	"bob":    {"anuj", "peggy"},
	"claire": {"thom", "jonny"},
	"anuj":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

type List struct {
	data string
	next *List
}

func NewList() *List {
	return &List{"", nil}
}

func (l *List) AddAll(names []string) {
	last := l
	for last.next != nil {
		last = last.next
	}
	for _, name := range names {
		last.next = &List{data: name, next: nil}
		last = last.next
	}
}

func (l *List) Next() *List {
	return l.next
}

func main() {
	l := NewList()
	l.AddAll(personNets["you"])
	var seller string
	seen := make(map[string]bool)
	l = l.Next()
	for l != nil {
		if !seen[l.data] {
			if isSeller(l.data) {
				seller = l.data
				break
			}
			seen[l.data] = true
			if names, res := personNets[l.data]; res == true {
				l.AddAll(names)
			}
			l = l.Next()
		}
	}
	if len(seller) != 0 {
		fmt.Println(seller)
	}
}

func isSeller(name string) bool {
	return strings.Contains(name, "m")
}
