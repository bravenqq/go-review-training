// Package main provides ...
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Counter int

func (c *Counter) Write(p []byte) (int, error) {
	*c = Counter(len(p))
	return len(p), nil
}

func main() {
	var w io.Writer
	var c Counter
	w = &c
	w.Write([]byte("hello"))
	fmt.Println("counter:", c)
	if w == nil {
		fmt.Println("w is nil")
	}
	var b *bytes.Buffer
	b = &bytes.Buffer{}
	w = b
	if w == nil {
		fmt.Println("w is still nil")
	}
	b.Write([]byte("hello"))
	fmt.Printf("w %T\n", w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// rwc = new(bytes.Buffer)
	w = rwc
	fmt.Printf("w %T\n", w)

}
