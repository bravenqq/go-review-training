// Package main provides ...
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	ph := flag.String("path", "/Users/abbynie/Desktop", "please input path")
	flag.Parse()
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.Header.Get("user-agent"))
		log.Println(req.Header.Get("Remote Address"))
		http.Redirect(w, req, "http://192.168.10.57:8088/"+req.RequestURI, http.StatusMovedPermanently)
	}))
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		err := server.ListenAndServe()
		log.Println("listen port 8080")
		if err != nil {
			log.Fatal("listen port:8080 err:", err)
		}
	}()

	// http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./"))))
	log.Println("listen port 8088")
	err := http.ListenAndServe("0.0.0.0:8088", http.FileServer(http.Dir(*ph)))
	if err != nil {
		log.Fatal("listen port:8088 err:", err)
	}

}
