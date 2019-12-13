package main

import "fmt"

func main() {
	var a, b int8 = -10, 10
	//逻辑运算时操作数的类型必须相同
	z := a & b
	fmt.Printf("z %d\n", z)
	y := a | b
	fmt.Printf("y %d\n", y)
	//^用于一元表示取反
	k := ^a
	fmt.Printf("k %d\n", k)
	//^用于二元运算表示异或
	j := a ^ b
	fmt.Printf("j %d\n", j)
	// &^按位清空
	i := a &^ b
	fmt.Printf("i %d\n", i)
	n := a >> 3
	fmt.Printf("n %d\n", n)
	m := a << 4
	fmt.Printf("m %d\n", m)
}
