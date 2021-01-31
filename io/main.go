// Package main provides ...
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("./test.md", os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for {
		n, err := f.WriteString("test ")
		if err != nil {
			log.Fatal("write err:", err)
		}
		fmt.Println("write: ", n)
		time.Sleep(time.Second)
	}

}
