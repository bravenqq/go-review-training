// Package main provides ...
package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "http://127.0.0.1:8090/"+req.RequestURI, http.StatusMovedPermanently)
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
	log.Println("listen port 8090")
	err := http.ListenAndServe(":8090", http.FileServer(http.Dir("./")))
	if err != nil {
		log.Fatal("listen port:8090 err:", err)
	}

}
