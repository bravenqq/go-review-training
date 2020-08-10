/*“编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。下面给出了2个例子：

func ElementsByTagName(doc *html.Node, name...string) []*html.Node
images := ElementsByTagName(doc, "img")
headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")”
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	for _, url := range os.Args[1:] {
		err := outline(url)
		if err != nil {
			log.Fatal(err)
		}
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
	nodes := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, node := range nodes {
		fmt.Println(node.Data)
	}

	return nil
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	forEachNode(doc, func(n *html.Node) bool {
		for _, node := range name {
			if n.Type == html.ElementNode && n.Data == node {
				nodes = append(nodes, n)
				return true
			}
		}
		return false
	}, nil)
	return nodes
}

func forEachNode(n *html.Node, startElement, endElement func(n *html.Node) bool) {
	startElement(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, startElement, endElement)
	}
}
