package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	p := make([]byte, 4096)
	var count int
	for {
		n, err := os.Stdin.Read(p[count+1:])
		if err != nil {
			log.Fatalln("input err:", err)
		}
		count += n
		if count == len(p) || strings.Contains(string(p[count-n:count]), "EOF") {
			break
		}
	}
	fmt.Print(string(p))
}
