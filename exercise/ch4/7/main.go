// Package main provides ...
package main

import (
	"unicode/utf8"
)

func main() {

}

func Mbreverse(b []byte) []byte {
	res := make([]byte, 0, len(b))
	for i := len(b); i > 0; {
		_, size := utf8.DecodeLastRune(b[:i])
		res = append(res, b[i-size:i]...)
		i = i - size
	}
	return res
}
