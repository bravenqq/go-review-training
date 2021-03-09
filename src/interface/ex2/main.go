// Package main provides ...
package main

import (
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type RedeWriter interface {
	Reader
	Writer
}

func main() {
	var r Reader
	r = os.Stderr
	_ = r
}
