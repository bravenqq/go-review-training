package main

import "fmt"

func main() {
	a := '聂'
	fmt.Printf("%d %T\n", a, a)
	b := 'a'
	fmt.Printf("%d %T\n", b, b)
	c := "聂"
	fmt.Printf("c length: %d\n", len(c))
	fmt.Printf("c length: %d\n", len([]rune(c)))
}
