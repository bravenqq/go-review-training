package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"golang.org/x/net/html"
)

var wg *sync.WaitGroup
var maxDepth = 3
var token chan struct{}
var lock sync.Mutex
var seen map[string]bool

var done chan struct{}
var cancel chan struct{}

func main() {
	wg = &sync.WaitGroup{}
	token = make(chan struct{}, 20)
	seen = make(map[string]bool)
	for _, link := range os.Args[1:] {
		wg.Add(1)
		go crawl(0, link, wg)
	}

	done = make(chan struct{})
	cancel = make(chan struct{})
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
		cancel <- struct{}{}
		close(done)
	}()
	wg.Wait()

	// panic(1)
}

func cacelled() bool {
	for {
		select {
		case <-done:
			return true
		default:
			return false

		}
	}
}

func crawl(depth int, url string, wg *sync.WaitGroup) {
	if cacelled() {
		fmt.Println("cancel")
		return
	}
	fmt.Printf("depth:%d,url:%s\n", depth, url)
	defer wg.Done()
	if depth >= maxDepth {
		return
	}
	token <- struct{}{}
	list, err := Extract(url)
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

func Extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
