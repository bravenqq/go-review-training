package main

import "fmt"

func main() {
	ages := map[string]int{"nqq": 25, "nzq": 25}
	delete(ages, "nqq")
	_, ok := ages["nqq"]
	fmt.Println(ok)
}
