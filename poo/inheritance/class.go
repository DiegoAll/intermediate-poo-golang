package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person   //Campo anónimo
	Employee //Campo anónimo
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Sujeto"
	ftEmployee.age = 10
	ftEmployee.id = 55555555
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee)
}
