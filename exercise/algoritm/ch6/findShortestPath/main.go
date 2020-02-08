/*
查找从A点到F点的最短路径是多少
*/

package main

import "fmt"

var paths = map[string][]string{
	"a": {"b", "d"},
	"b": {"c"},
	"d": {"g", "e"},
	"g": {"c"},
	"e": {"c"},
	"c": {"f"},
}

func main() {
	near := paths["a"]
	var path int
	var fund bool
	for len(near) != 0 {
		temp := near
		near = nil
		path++
		for _, p := range temp {
			if p == "f" {
				fund = true
				break
			}
			if pt, res := paths[p]; res {
				near = append(near, pt...)
			}
		}
	}
	if fund {
		fmt.Println("找到了最短路径，长度是：", path)
	}
}

func getShortestPath() string {
	// pathsAll := make(map[int][]string)
	// min := math.MaxInt8

	getPath()
}
