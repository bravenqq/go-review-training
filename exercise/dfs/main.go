package main

import (
	"fmt"

	"golang.org/x/net/html"
)

type Stack []*html.Node

var stack Stack

func (s *Stack) Push(n *html.Node) {
	slice := []*html.Node(*s)
	slice = append(slice, n)
	stack := Stack(slice)
	s = &stack
}

func (s *Stack) Pop() (*html.Node, bool) {
	if len(*s) == 0 {
		return nil, false
	}

	slice := []*html.Node(*s)
	n := slice[len(slice)-1]
	slice = slice[:len(slice)-2]
	stack := Stack(slice)
	s = &stack
	return n, true
}

func main() {
	var n *html.Node
	var s Stack
	pushSub(s, n)
	for {
		node, res := s.Pop()
		if !res {
			break
		}
		fmt.Println(n.Data)
		next := node.NextSibling
		pushSub(s, next)
	}
}

func pushSub(s Stack, n *html.Node) {
	fmt.Println(n.Data)
	for {
		node := n.FirstChild
		if node == nil {
			break
		}
		s.Push(node)
	}
}
