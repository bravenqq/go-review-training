// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(strings.Join(os.Args, " "))
}
