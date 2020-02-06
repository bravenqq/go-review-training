package main

import (
	"bytes"
	"strings"
)

func main() {

}

func Expand(s string, f func(string) string) string {
	if !strings.Contains(s, "foo") {
		return s
	}
	data := []byte(s)
	rep := []byte("foo")
	i := bytes.Index(data, rep)
	data = append(data[:i], []byte(f("foo"))...)
	return string(data) + Expand(string([]byte(s)[i+len("foo"):]), f)
}
