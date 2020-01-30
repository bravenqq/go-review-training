package main

func main() {

}

//向右旋转n位的计算方式位s[(i+n)%len(n)] = s[i]
//向左旋转n位相当于向右旋转len(s)-n
func LeftRotate(s []int, n int) []int {
	r := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[(i+len(s)-n)%len(s)] = s[i]
	}
	return r
}
