package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"gopl.io/ch5/links"
)

var wg *sync.WaitGroup
var maxDepth = 3
var token chan struct{}
var lock sync.Mutex
var seen map[string]bool

func main() {
	core := flag.Int("core", 2, "please inout core num")
	flag.Parse()
	runtime.GOMAXPROCS(*core)
	start := time.Now()
	wg = &sync.WaitGroup{}
	token = make(chan struct{}, 20)
	seen = make(map[string]bool)
	for _, link := range os.Args[1:] {
		wg.Add(1)
		go crawl(0, link, wg)
	}
	wg.Wait()
	fmt.Println("time:", time.Since(start))
}

func crawl(depth int, url string, wg *sync.WaitGroup) {
	fmt.Printf("depth:%d,url:%s\n", depth, url)
	defer wg.Done()
	if depth >= maxDepth {
		return
	}
	token <- struct{}{}
	list, err := links.Extract(url)
	<-token
	if err != nil {
		log.Println(err)
	}
	for _, link := range list {
		lock.Lock()
		if seen[link] {
			lock.Unlock()
			continue
		}
		seen[link] = true
		lock.Unlock()
		wg.Add(1)
		go crawl(depth+1, link, wg)
	}
}
