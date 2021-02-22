// https://segmentfault.com/a/1190000016412013
// https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
// Package main provides ...
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	// "runtime/pprof"
)

func main() {
	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	// pprof.WriteHeapProfile(os.Stdout)
	http.ListenAndServe("127.0.0.1:3333", nil)

}

func bigBytes() *[]byte {
	s := make([]byte, 1000000)
	return &s
}
