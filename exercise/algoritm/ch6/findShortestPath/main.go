/*
查找从A点到F点的最短路径是多少
*/

package main

import "fmt"
import "strings"

var paths = map[string][]string{
	"a": {"b", "d"},
	"b": {"c"},
	"d": {"g", "e"},
	"g": {"c"},
	"e": {"c"},
	"c": {"f"},
}

// type Node struct{
// 	data string
// 	next *Node
// }
//
// type LinkList struct{
// 	head *Node
// 	tail *Node
// }
//
// func NewLinkList() LinkList{
// 	return LinkList{head:nil,tail: nil}
// }
//
// func (l LinkList)Add(value string) LinkList {
// 	node := new(Node)
// 	node.data = value
// 	if l.tail == nil {
// 		l.tail = node
// 		l.head = node
// 		node.next = nil
// 		return l
// 	}
// 	l.tail.next = node
// 	l.tail = node
// 	l.tail = nil
// 	return l
// }

func main() {
	near := paths["a"]
	var path int
	var fund bool
	var ps []string
	for _, key := range near {
		ps = append(ps, "a"+key)
	}

	var findParent func(p string) string
	findParent = func(p string) string {
		for _, v := range ps {
			if v[len(v)-1:] == p {
				return v
			}
		}
		return ""
	}

	for len(near) != 0 {
		temp := near
		near = nil
		path++
		for _, p := range temp {
			if p == "f" {
				fund = true
				break
			}
			if pt, res := paths[p]; res {
				near = append(near, pt...)
				parent := findParent(p)
				for _, key := range pt {
					ps = append(ps, parent+key)
				}
			}
		}
	}

	if fund {
		fmt.Println("找到了最短路径，长度是：", path)
	}
	for _, v := range ps {
		if strings.Contains(v, "f") {
			fmt.Println("最短路径是：", v)
		}
	}
}
