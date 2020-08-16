// Package main provides ...
package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("welcome to index"))
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

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./"))))
	log.Println("listen port 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("listen port:8090 err:", err)
	}

}
