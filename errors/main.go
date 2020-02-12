package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "main.go", "please input file name")
	flag.Parse()
	_, err := os.Open(file)
	var openErr error
	if err != nil {
		openErr = fmt.Errorf("file:%s err:%w", file, err)
	}

	if errors.Is(openErr, os.ErrNotExist) {
		log.Println(openErr)
	}

	fmt.Println(errors.Unwrap(openErr))

}
