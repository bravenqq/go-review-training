// Package main provides ...
package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "welcom to the home page!")
	})
	n := negroni.Classic()
	n.Use()
	n.UseHandler(mux)
	n.Run(":3000")
}
