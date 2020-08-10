package main

import (
	"flag"
	"fmt"
)

func main() {
	var str1, str2 string
	flag.StringVar(&str1, "sub1", "clues", "please input substring")
	flag.StringVar(&str2, "sub2", "blue", "please input substring")
	flag.Parse()

	var res [][]int
	for i := 0; i < len(str1); i++ {
		arr := make([]int, len(str2))
		res = append(res, arr)
	}
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				res[i][j] = res[i-1][j-1] + 1
			}
		}
	}

	var sub []byte
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if res[i][j] != 0 {
				sub = append(sub, str1[i])
			}
		}
	}
	fmt.Println("最长公共子串：", string(sub))
}
