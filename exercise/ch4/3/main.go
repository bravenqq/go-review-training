package main

var n int

func main() {

}

func reverse(s *[32]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j+1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
