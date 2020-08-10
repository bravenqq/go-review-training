package main

import (
	"strings"
)

func main() {

}

func Join(args []string) string {
	s, sep := "", ""
	for _, str := range args {
		s = s + sep + str
		sep = " "
	}
	return s
}

func FastJoin(args []string) string {
	return strings.Join(args, " ")
}
