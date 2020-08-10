// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", Comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
// func Comma(s string) string {
// 	n := len(s)
// 	if n <= 3 {
// 		return s
// 	}
// 	return Comma(s[:n-3]) + "," + s[n-3:]
// }

func Comma(s string) string {
	var symbol byte
	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		symbol = s[0]
		s = s[1:]
	}

	slice := strings.Split(s, ".")
	s = CommaFast(slice[0])
	if len(slice) > 1 {
		s = s + "." + slice[1]
	}
	if symbol == '+' || symbol == '-' {
		s = string(symbol) + s
	}
	return s
}

func CommaFast(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	count := n / 3
	if n%3 == 0 {
		count--
	}

	buffer := bytes.NewBuffer(make([]byte, 0, count+n))

	var i int
	if n%3 == 0 {
		buffer.WriteString(s[:3])
	} else {
		buffer.WriteString(s[:n%3])
	}
	i = buffer.Len()
	for {
		buffer.WriteByte(',')
		if i+3 >= n {
			buffer.WriteString(s[i:])
			break
		}
		buffer.WriteString(s[i : i+3])
		i += 3
	}
	return buffer.String()
}

//!-
