package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//Hemos creado un Employee nuevo basado del struct
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	//2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	e3 := new(Employee)
	//fmt.Printf("%v\n", e3)
	//fmt.Printf("%v\n", *e3)

	e3.id = 20
	e3.name = "Jose"
	//fmt.Printf("%v\n", *e3)

	e4 := NewEmployee(4, "The Ramon", true)
	fmt.Printf("%v\n", *e4)
}
