// Package main provides ...
package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := "test"
	fmt.Println(a)
	http.ListenAndServe(":6060", nil)
}
