package main

import (
	"flag"
	"fmt"
)

func main() {
	var str1, str2 string
	flag.StringVar(&str1, "sub1", "fish", "please input substring")
	flag.StringVar(&str2, "sub2", "fosh", "please input substring")
	flag.Parse()

	var res [][]int
	var sub []byte
	var count int
	for i := 0; i < len(str1); i++ {
		arr := make([]int, len(str2))
		res = append(res, arr)
	}
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				sub = append(sub, str1[i])
				count++
				if i == 0 || j == 0 {
					res[i][j] = 1
					continue
				}
				res[i][j] = res[i-1][j-1] + 1
			} else {
				if i == 0 && j == 0 {
					res[i][j] = 0
					continue
				}
				if i == 0 {
					res[i][j] = res[i][j-1]
					continue
				}
				if j == 0 {
					res[i][j] = res[i-1][j]
					continue
				}
				res[i][j] = res[i][j-1]
				if res[i][j] < res[i-1][j] {
					res[i][j] = res[i-1][j]
				}
			}
		}
	}

	fmt.Printf("最长公共子串:%s,长度:%d\n", string(sub), count)
}
