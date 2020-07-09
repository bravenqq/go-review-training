// Package main provides ...
package main

import "net/http"

func main() {
	router := http.NewServeMux()
	router.Handle()

	http.ListenAndServe()
}
