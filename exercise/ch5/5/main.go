package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// if len(os.Args) != 2 {
	// 	log.Fatal("请输入要统计的url")
	// }
	words, imags, err := CountWordsAndImages(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("words:", words)
	fmt.Println("images:", imags)
}

func CountWordsAndImages(input io.Reader) (words, images int, err error) {
	// resp, err := http.Get(url)
	// if err != nil {
	// 	err = fmt.Errorf("get url:%s err:%v", url, err)
	// 	return
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode != http.StatusOK {
	// 	err = fmt.Errorf("get url:%s status:%d", url, resp.StatusCode)
	// 	return
	// }
	n, err := html.Parse(input)
	if err != nil {
		err = fmt.Errorf("parse err:%v", err)
		return
	}
	words, images = countWordsAndImages(0, 0, n)
	return
}

func countWordsAndImages(words, images int, n *html.Node) (int, int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.ElementNode && n.Data == "script" || n.Data == "style" {
		return words, images
	}
	if n.Type == html.TextNode {
		words += CountWords(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		words, images = countWordsAndImages(words, images, c)
	}
	return words, images
}

func CountWords(str string) int {
	var words int
	scanner := bufio.NewScanner(bytes.NewBufferString(str))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return words
}
