package main

func main() {

}

func Count(l *List) int {
	if l == nil {
		return 0
	}
	return 1 + Count(l.Next)
}

func FindMax(max int, l *List) int {
	if l == nil {
		return max
	}

	if l.Data > max {
		return FindMax(l.Data, l.Next)
	} else {
		return FindMax(max, l.Next)
	}
}

type List struct {
	Data int
	Next *List
}
