// Package main provides ...
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	path := flag.String("path", "/Users/abbynie/Desktop/bin", "please input path")
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(*path)))
	if err != nil {
		log.Fatal(err)
	}
}
