// Package main provides ...
package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	//jiang buf变量赋值给out时，虽然buf是一个值为nil的指针，但是out的动态值!=nil,是buf指针的拷贝
	// ...do something...
	fmt.Printf("%T %v\n", out, out) //*bytes.Buffer <nil>
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
