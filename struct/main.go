package main

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	ID          int
	Salary      int
	Subordinate []Person
}

type Point struct {
	X float64
	Y float64
}

func main() {
	var p Employee
	p.Name = "nqq"

	e := new(Employee)
	e.Name = "test"
	f()
}

func f() Employee {
	var e Employee
	e.Subordinate = make([]Person, 10)
	return e
}

type Node struct {
	next *Node
	Data Person
}
