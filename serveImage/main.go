// Package main provides ...
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Index)
	router.HandleFunc("/image", GetImage)
	log.Fatal(http.ListenAndServe(":8080", router))
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./fe6dae00-c0e2-11ea-9a0b-a45e60f2fb75")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w.Header().Add("Content-Type", "image/jpeg")
	io.Copy(w, f)

}
