package main

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	ID     int
	Salary int
}

func main() {
	var p Employee
	p.Name = "nqq"

	e := new(Employee)
	e.Name = "test"

}

type Node struct {
	next *Node
	Data Person
}
