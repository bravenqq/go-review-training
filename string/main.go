package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
func main() {
	str := "h好，中国"
	bytes := []rune(str)
	for _, s := range bytes {
		fmt.Printf("%b ", s)
	}

}

func Out(s string) {
	fmt.Printf("s:%p", &s)
}
