package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	buffer := bytes.NewBuffer(nil)

	f, err := os.Open("./test")
	if err != nil {
		log.Fatalf("open file test err:%v\n", err)
	}
	defer f.Close()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		_, err = buffer.ReadFrom(f)
		if err != nil && err != io.EOF {
			log.Fatalf("read from file test err:%v\n", err)
		}
		wg.Done()
	}()

	p := make([]byte, bytes.MinRead)

	go func() {
		for {
			_, err := buffer.Read(p)
			if err != nil && err != io.EOF {
				log.Fatal("read from buffer err:", err)
			}
			fmt.Print(string(p))
		}
	}()
	wg.Wait()
}
