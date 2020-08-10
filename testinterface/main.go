package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var r io.Reader
	r = os.Stdout
	fmt.Printf("%T\n", r)
	// stdout := r.(*os.File)
	// fmt.Printf("%T\n", stdout)
	rw := r.(io.ReadWriter)
	fmt.Printf("%T\n", rw)
}

func Sprint(v interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := v.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		} else {
			return "false"
		}
	default:
		return "???"
	}
}
