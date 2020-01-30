package main

import "fmt"

func main() {
	fmt.Println(EqualLetter("abcder", "acderb"))
}

func EqualLetter(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	letter1 := make(map[byte]int)
	letter2 := make(map[byte]int)
	for i := 0; i < len(str1); i++ {
		ch := str1[i]
		letter1[ch]++
	}
	for i := 0; i < len(str2); i++ {
		ch := str2[i]
		letter2[ch]++
	}

	for key, value := range letter1 {
		if letter2[key] != value {
			return false
		}
	}
	return true
}
