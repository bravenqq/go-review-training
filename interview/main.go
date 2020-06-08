// Package main provides ...
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

}

func Traverse(dir string) {
	for _, entry := range dirents(dir) {
		fmt.Printf("%+v\n", entry)
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
