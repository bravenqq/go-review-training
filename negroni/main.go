// Package main provides ...
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("test")
		res.Write([]byte("test\n"))
		res.WriteHeader(200)
	}))

	http.ListenAndServe(":8080", nil)
}
