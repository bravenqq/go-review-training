// Package main provides ...
package main

import "net/http"

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

type Handle func(a, b int) int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	})
	var f func(a, b int) int
	var res bool
	if res {
		f = add
	} else {
		f = minus
	}

	_ = f
}

func Map(h Handle, a, b int) int {
	return h(a, b)
}
