package main

func main() {

}

func Qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var left []int
	var right []int
	bench := arr[len(arr)/2]
	for i, v := range arr {
		if i == len(arr)/2 {
			continue
		}
		if v > bench {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}
	return append(append(Qsort(left), bench), Qsort(right)...)
}
