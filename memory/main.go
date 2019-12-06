package main

func main() {
	x := make([]int, 10)
	// escape to heap,申请的内存过大
	y := make([]int, 20000)
	l := 10
	// escap to heap，不确定大小
	d := make([]int, l)
	c := make(map[string]string)
	names := []string{"nqq", "hxw"}
	s := ""
	for _, name := range names {
		s += name + " "
	}
	f := Fibonacii
}

func Fibonacii() func() int {
	// escape to heap 函数结束时变量仍被别的函数使用
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func incr(x int) *int {
	x += 1

	//escape to heap,函数结束后反回了地址
	return &x
}
