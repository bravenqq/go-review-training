// Package main provides ...
package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

//Log log requst url middleware
func Log(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Println("request url:", req.URL.String())
		start := time.Now()
		h(resp, req)
		log.Println("spend time:", time.Now().Sub(start))
	})
}

func main() {
	addr := flag.String("addr", ":8080", "please input addr")
	mux := http.NewServeMux()
	mux.HandleFunc("/index", Log(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))
	}))
	log.Println("start server addr ", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
