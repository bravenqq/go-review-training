package main

import (
	"fmt"
	"math"
)

func main() {
	a := -0.12444341
	fmt.Printf("%T\n", a)
	//float32能表示的最大值
	fmt.Printf("float32 max value:%g \n", math.MaxFloat32)
	//floa32能表示的最小正数
	fmt.Printf("float32 min value:%g \n", math.SmallestNonzeroFloat32)
	//float64能表示的最大值
	fmt.Printf("float64 max value:%g \n", math.MaxFloat64)
	//float64能表示的最小正数
	fmt.Printf("float64 min value:%g \n", math.SmallestNonzeroFloat64)

	fmt.Println("-Inf ", math.Inf(-1))
	fmt.Println("Inf ", math.Inf(1))

}
