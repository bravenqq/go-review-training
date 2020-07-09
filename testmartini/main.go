// Package main provides ...
package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World!"
	})
	m.Get("/report", func() (int, string) {
		return 418, "I'm a report"
	})
	m.Get("/code", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
		res.WriteHeader(200) // HTTP 200
	})
	m.Patch("/", func() string {
		return "patch"
	})
	m.Post("/", func() string {
		return "post"
	})
	m.Delete("/", func() string {
		return "delete"
	})
	m.Options("/", func() string {
		return "options"
	})
	m.NotFound(func() (int, string) {
		// handle 404
		return 404, "not found"
	})

	m.Get("/hello/**", func(params martini.Params) string {
		return "hello " + params["_1"]
	})
	authorize := func() bool {
		return false
	}
	m.Get("/secret", authorize, func() string {
		return "secret"
	})

	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")

		c.Next()

		log.Println("after a request")
	})

	db := &MyDatabase{}
	// the service will be available to all handlers as *MyDatabase
	// ...
	m.Map(db)
	m.Run()
}

type MyDatabase struct {
}
