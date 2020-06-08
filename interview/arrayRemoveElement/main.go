// Package main provides ...
package main

import "fmt"

/*给定一个int数组删除其中的所有出现的某个数例如10*/
func main() {
	arr := []int{2, 10, 5, 10, 0}
	fmt.Printf("%+v\n", arr)
}

func Remove(arr []int, num int, order bool) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			if order {
				arr = append(arr[:i], arr[i+1:]...)
			} else {
				j := len(arr) - 1
				for {
					if arr[j] != num {
						break
					}
					j--
				}
				if j != i {
					arr[j] = arr[i]
				}
				arr = arr[:j]
			}

		}
	}
	return arr
}
