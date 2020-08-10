package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	bf := bufio.NewScanner(os.Stdin)
	bf.Split(bufio.ScanWords)
	for bf.Scan() {
		if bf.Text() == "eof" {
			break
		}
		counts[bf.Text()]++
	}
	if err := bf.Err(); err != nil {
		fmt.Println("scanerr:", err)
	}
	for k, v := range counts {
		fmt.Printf("%s %d\n", k, v)
	}
}
