// Package main provides ...
package main

import (
	"os"
)

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func tempDirs() []string {
	return []string{"./test"}
}

func main() {

	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d               // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		i := 10
		rmdirs = append(rmdirs, func() {
			a := i
			_ = a
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}

}
