// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Finddatas1 prints the datas in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "finddatas1: %v\n", err)
		os.Exit(1)
	}
	for _, data := range visit(nil, doc) {
		fmt.Println(data)
	}
}

//!-main

//!+visit
// visit appends to datas each link found in n and returns the result.
func visit(datas []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "script" || n.Data == "style" {
		return datas
	}
	if n.Type == html.TextNode && n.DataAtom != atom.Script && n.DataAtom != atom.Style {
		datas = append(datas, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		datas = visit(datas, c)
	}
	return datas
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
