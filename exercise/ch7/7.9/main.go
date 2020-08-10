// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//create coustomSort
var cs = NewCoustomSort(tracks)

var tmpl *template.Template

var tmplStr = `<html>
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width" />
        <title>Template</title>
     </head>
    <body>
        <table>
        <tr>
			<th><a href="http://localhost:8080/sort/0" target="_blank">Title</a></th>
            <th><a href="http://localhost:8080/sort/1" target="_blank">Artist</a></th>
            <th><a href="http://localhost:8080/sort/2" target="_blank">Album</a></th>
            <th><a href="http://localhost:8080/sort/3" target="_blank">Year</a></th>
            <th><a href="http://localhost:8080/sort/4" target="_blank">Length</a></th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Title}}</td>
            <td>{{.Artist}}</td>
            <td>{{.Album}}</td>
            <td>{{.Year}}</td>
            <td>{{.Length}}</td>
        </tr>
        {{end}}
        </table>
    </body>
</html>`

func main() {
	// var err error
	//init template from file
	tmpl = template.Must(template.New("template").Parse(tmplStr))

	//!+customcall
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/sort/:serial", sortTrack)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sort.Sort(cs)
	err := tmpl.ExecuteTemplate(w, "template", cs.t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func sortTrack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ser := ps.ByName("serial")
	serial, err := strconv.Atoi(ser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cs.Select(serial)
	sort.Sort(cs)
	err = tmpl.ExecuteTemplate(w, "template", cs.t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

//!+customcode
type customSort struct {
	t       []Track
	less    columnCmp
	cloumns []columnCmp
}

type columnCmp func(x, y Track) bool

func (x *customSort) LessArticle(a, b Track) bool {
	return a.Title < b.Title
}

func (x *customSort) LessArtist(a, b Track) bool {
	return a.Artist < b.Artist
}

func (x *customSort) LessAlbum(a, b Track) bool {
	return a.Album < b.Album
}

func (x *customSort) LessYear(a, b Track) bool {
	return a.Year < b.Year
}

func (x *customSort) LessLength(a, b Track) bool {
	return a.Length < b.Length
}

func NewCoustomSort(tracks []Track) *customSort {
	c := customSort{t: tracks}
	c.less = c.LessYear
	c.AddAllCondation()
	return &c
}

func (x *customSort) AddAllCondation() {
	x.cloumns = []columnCmp{x.LessArticle, x.LessArtist, x.LessAlbum, x.LessYear, x.LessLength}
}

func (x *customSort) Select(i int) {
	if i >= len(x.cloumns) {
		log.Fatalf("%d condition not exit", i)
	}
	x.less = x.cloumns[i]
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
