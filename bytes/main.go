package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	buffer := bytes.NewBuffer(nil)

	f, err := os.Open("./test")
	if err != nil {
		log.Fatalf("open file test err:%v\n", err)
	}
	defer f.Close()

	go func() {
		for {
			_, err := buffer.ReadFrom(f)
			if err == io.EOF {
				break
			}
		}
	}()

	for {
		_, err = buffer.WriteTo(os.Stdout)
		if err == io.EOF {
			break
		}

	}
}
