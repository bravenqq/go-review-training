package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

var (
	NotFound string = "没有找到数字：%d\n"
)

func main() {
	num := flag.Int("number", 0, "请输入你要查找的数据")
	length := flag.Int("cap", 100, "请输入测试的数据量")
	flag.Parse()

	data := make([]int, 0, *length)
	data = SortInsert(data, *length)
	fmt.Printf("data:%v\n", data)
	index, err := DirectorySearch(data, *num)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%d的位置是:%d\n", *num, index)
}

func SortInsert(data []int, length int) []int {
	for i := 0; i < length; i++ {
		x := rand.Int() % length
		data = Insert(data, x)
	}
	return data
}

func Insert(data []int, x int) []int {
	data = append(data, x)
	if len(data) == 1 {
		return data
	}
	var pos int
	for i, y := range data {
		if y >= x {
			pos = i
			break
		}
	}

	if pos == len(data)-1 {
		return data
	}

	for i := len(data) - 1; i > pos; i-- {
		data[i] = data[i-1]
	}
	data[pos] = x
	return data
}

func Search(data []int, num int) (int, error) {
	for i, v := range data {
		if v == num {
			return i, nil
		}
	}
	return 0, fmt.Errorf(NotFound, num)
}

func DirectorySearch(data []int, num int) (int, error) {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == num {
			return mid, nil
		}

		if data[mid] <= num {
			low = mid + 1
		}

		if data[mid] > num {
			high = mid - 1
		}
	}
	return 0, fmt.Errorf(NotFound, num)
}
