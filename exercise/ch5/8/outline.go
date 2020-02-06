// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	id := flag.String("id", "page", "please input id")
	flag.Parse()
	fmt.Printf("id:%s of element:%s\n", *id, ElementByID(doc, *id).Data)
	//!+call
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}
	return nil
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
	}
	return false
}

// func endElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		depth--
// 		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
// 	}
// }

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}

//!-startend
