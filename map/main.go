// Package main provides ...
package main

func main() {
	// var m map[int]string
	// m = make(map[int]string)
	// m[1] = "hello"
	m := f()
	_ = m
	// fmt.Println(m[1])//world
}

func f() map[int]string {
	m := make(map[int]string)
	m[1] = "world"
	return m
}
