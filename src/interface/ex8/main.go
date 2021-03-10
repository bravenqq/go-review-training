// Package main provides ...
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w interface{}
	w = os.Stdout
	switch v := w.(type) {
	case *bytes.Buffer:
		v.Write([]byte("buffer\n"))
	case *os.File:
		v.Write([]byte("file\n"))
	default:
		fmt.Printf("%T %v\n", v, v)
	}

	switch v := w.(type) {
	case io.Reader:
		p := make([]byte, 512)
		v.Read(p)
		fmt.Println(string(p))
	case io.Closer:
		fmt.Println("close")
	case io.Writer:
		v.Write([]byte("hello\n"))
	default:
		fmt.Printf("%T %v\n", v, v)

	}
}
