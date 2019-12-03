package main

// import (
// 	"os"
// )

func main() {
	names := []string{"abby", "nqq", "nieqianqian"}
	s, sep := "", ""
	for _, name := range names {
		s += sep + name
		sep = " "
	}

}
