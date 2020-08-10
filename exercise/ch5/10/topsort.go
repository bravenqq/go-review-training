// Package main provides ...
package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order := topSort(prereqs)
	for i := 0; i < len(order); i++ {
		fmt.Printf("%d:\t%s\n", i, order[i])
	}
}

func topSort(m map[string][]string) map[int]string {
	seen := make(map[string]bool)
	order := make(map[int]string)
	var visitAll func(items []string)
	var i int
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[i] = item
				i++
			}
		}
	}

	var items []string
	for key, _ := range m {
		items = append(items, key)
	}
	sort.Strings(items)
	visitAll(items)
	return order
}
