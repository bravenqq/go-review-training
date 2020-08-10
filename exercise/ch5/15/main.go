/*
“编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本”
*/

package main

func main() {

}

func Max(in ...int) int {
	if len(in) == 0 {
		panic("请至少传一个参数")
	}

	max := in[0]
	for _, x := range in {
		if x > max {
			max = x
		}
	}
	return max
}

func Min(in ...int) int {

	if len(in) == 0 {
		panic("请至少传一个参数")
	}

	min := in[0]
	for _, x := range in {
		if x < min {
			min = x
		}
	}
	return min
}
