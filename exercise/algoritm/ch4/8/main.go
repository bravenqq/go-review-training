package main

func main() {

}

func Multiplicat(res, factor []int) []int {
	if len(factor) == 0 {
		return res
	}
	for i := 0; i < len(res); i++ {
:
		res[i] *= factor[0]
	}
	return Multiplicat(res, factor[1:])
}
