package main

import (
	"unicode"
	"unicode/utf8"
)

func main() {
}

func RemoveExtraSpace(data []byte) []byte {
	var rn []rune
	l := 0
	for l < len(data) {
		r, size := utf8.DecodeRune(data[l:])
		rn = append(rn, r)
		l += size
	}
	if len(rn) <= 1 {
		return data
	}
	i := 1
	for j := 1; j < len(rn); j++ {
		if rn[j] == rn[j-1] && unicode.IsSpace(rn[j]) {
			continue
		}
		rn[i] = rn[j]
		i++
	}
	rn = rn[:i]
	return []byte(string(rn))
	// for j := 1; j < len(data); j++ {
	// 	if data[j] == data[j-1] && data[j] == ' ' {
	// 		continue
	// 	}
	// 	data[i] = data[j]
	// 	i++
	// }
	// return data[:i]
}
