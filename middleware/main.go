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
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("request url:", req.URL.String())
		start := time.Now()
		h(w, req)
		log.Println("spend time:", time.Now().Sub(start))
	})
}

func LogHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("request url ", req.URL.RequestURI())
		start := time.Now()
		h.ServeHTTP(w, req)
		log.Println("spend time:", time.Now().Sub(start))
	})
}

type Martini struct {
	handelrs []http.HandlerFunc
	action   http.Handler
	index    int
}

func (m Martini) Use(h http.HandlerFunc) {
	m.handelrs = append(m.handelrs, h)
}

func (m Martini) Action(h http.HandlerFunc) {
	m.action = h
}

func (m Martini) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for m.index <= len(m.handelrs) {
		h := m.Get()
		h(w, req)
		m.index++
	}
}

func (m Martini) Get() http.HandlerFunc {
	if m.index < len(m.handelrs) {
		return m.handelrs[m.index]
	}
	return m.action.ServeHTTP
}

func main() {
	addr := flag.String("addr", ":8080", "please input addr")
	mux := http.NewServeMux()
	mux.HandleFunc("/index", Log(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
	}))
	mux.Handle("/hello", LogHandler(Person{"nqq"}))
	log.Println("start server addr ", *addr)

	log.Fatal(http.ListenAndServe(*addr, mux))
}

type Person struct {
	Name string
}

func (p Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + p.Name))
}
